import requests
import os
import redis
import json
import requests
import os
import redis
import json
import threading
import uuid
import datetime

from langchain_community.llms import HuggingFacePipeline
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline

model_name = "rinna/japanese-gpt2-medium"
tokenizer = AutoTokenizer.from_pretrained(model_name, use_fast=False)
model = AutoModelForCausalLM.from_pretrained(model_name)
pipe = pipeline(
    "text-generation",
    model=model,
    tokenizer=tokenizer,
    max_new_tokens=48,
    temperature=1.1,
    top_p=0.85,
    repetition_penalty=1.7
)
llm = HuggingFacePipeline(pipeline=pipe)

EMBEDDING_API_URL = os.environ.get("EMBEDDING_API_URL", "http://embedding:8001/embed")
WEAVIATE_URL = os.environ.get("WEAVIATE_URL", "http://weaviate:8080")
REDIS_HOST = os.environ.get("BROKER_HOST", "broker")
REDIS_PORT = int(os.environ.get("BROKER_PORT", "6379"))
REDIS_DB = int(os.environ.get("BROKER_DB", "0"))
REDIS_PASSWORD = os.environ.get("BROKER_PASSWORD", "")

redis_client = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, db=REDIS_DB, password=REDIS_PASSWORD, decode_responses=True)

def handle_rag_event(event_json):
    try:
        event = json.loads(event_json)
        payload = event.get("payload", {})
        # toがリストの場合は最初の要素を使う
        to_value = event.get("to")
        if isinstance(to_value, list):
            participant_id = to_value[0] if to_value else None
        else:
            participant_id = to_value
        question = payload.get("content") or payload.get("question")
        chat_id = event.get("chat_id") or payload.get("chat_id")
        question_id = payload.get("id") or payload.get("id")
        opponent_id = payload.get("participant_id") or payload.get("participant_id")

        event_id = str(uuid.uuid4())
        answer_id = str(uuid.uuid4())

        embed_resp = requests.post(
            EMBEDDING_API_URL,
            json={"texts": [question]}
        )
        embed_resp.raise_for_status()
        embedding = embed_resp.json()["embeddings"][0]

        w_resp = requests.post(
            f"{WEAVIATE_URL}/v1/graphql",
            json={
                "query": f"{{Get{{QAPair(nearVector:{{vector:{embedding}}},limit:3){{question answer}}}}}}"
            }
        )
        w_resp.raise_for_status()
        qa_pairs = []
        try:
            qa_pairs = w_resp.json()["data"]["Get"]["QAPair"]
        except Exception:
            qa_pairs = []

        context = "\n".join([
            f"{qa.get('question','')}\n{qa.get('answer','')}" for qa in qa_pairs
        ])
        system_prompt = (
            "あなたは以下の参考情報（context）だけを根拠に質問に答えてください。\n"
            "contextに直接書かれていない内容は推測せず、「わかりません」や「contextに情報がありません」と答えてください。\n"
            "context以外の知識や一般論は使わないでください。\n\n"
        )
        prompt = f"{system_prompt}context:\n{context}\n\n質問: {question}\n答え:"
        from transformers import TextIteratorStreamer
        import time
        streamer = TextIteratorStreamer(tokenizer, skip_prompt=True, skip_special_tokens=True)

        gen_thread = threading.Thread(target=model.generate, kwargs={
            "input_ids": tokenizer(prompt, return_tensors="pt").input_ids,
            "max_new_tokens": 48,
            "temperature": 1.1,
            "top_p": 0.85,
            "repetition_penalty": 1.7,
            "streamer": streamer
        })
        gen_thread.start()
        answer = ""
        for token in streamer:
            answer += token

            stream_event = {
                "id": event_id,
                "chat_id": chat_id,
                "type": "stream",
                "from": participant_id,
                "to": [opponent_id] if opponent_id is not None else [],
                "timestamp": int(time.time()),
                "payload": {
                    "id": answer_id,
                    "chat_id": chat_id,
                    "question_id": question_id,
                    "participant_id": participant_id,
                    "content": answer,
                    "created_at": datetime.datetime.utcnow().isoformat() + 'Z',
                    "attachments": None
                }
            }
            redis_client.publish("chat_stream", json.dumps(stream_event))
        gen_thread.join()

        chat_event = {
            "id": event_id,
            "chat_id": chat_id,
            "type": "ai_answer",
            "from": participant_id,
            "to": [opponent_id] if opponent_id is not None else [],
            "timestamp": int(time.time()),
            "payload": {
                "id": answer_id,
                "chat_id": chat_id,
                "question_id": question_id,
                "participant_id": participant_id,
                "content": answer,
                "created_at": datetime.datetime.utcnow().isoformat() + 'Z',
                "attachments": None
            }
        }
        redis_client.publish("chat_events", json.dumps(chat_event))
    except Exception as e:
        print(f"[RAG] Error handling event: {e}", flush=True)

def subscribe_rag_requests():
    pubsub = redis_client.pubsub()
    pubsub.subscribe("rag_requests")
    print("[RAG] Subscribed to rag_requests")
    for message in pubsub.listen():
        if message["type"] == "message":
            threading.Thread(target=handle_rag_event, args=(message["data"],)).start()

if __name__ == "__main__":
    subscribe_rag_requests()
