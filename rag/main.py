from fastapi import FastAPI, Request, Header, HTTPException
from pydantic import BaseModel
from typing import Optional, List
import requests
import os


from langchain_community.llms import HuggingFacePipeline
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline



app = FastAPI()


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




class QueryRequest(BaseModel):
    question: str
    chat_id: Optional[str] = None
    question_id: Optional[str] = None
    participant_id: Optional[str] = None



class QueryResponse(BaseModel):
    answer: str



@app.get("/")
async def root():
    return {"message": "Hello, FastAPI!"}




@app.post("/rag/query", response_model=QueryResponse)
async def rag_query(request: QueryRequest, fastapi_request: Request):
    auth_header = fastapi_request.headers.get("authorization")
    print(f"[DEBUG] Authorization header: {auth_header}")
    if not auth_header:
        raise HTTPException(status_code=401, detail="Authorization token required")

    embed_resp = requests.post(
        EMBEDDING_API_URL,
        json={"texts": [request.question]}
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
    prompt = f"{system_prompt}context:\n{context}\n\n質問: {request.question}\n答え:"
    answer = llm(prompt)

    api_url = os.environ.get("API_URL", "http://api:8000")
    if request.chat_id and request.question_id and request.participant_id:
        # フロントからのAuthorizationヘッダーをそのまま転送
        headers = {}
        auth_header = fastapi_request.headers.get("authorization")
        if auth_header:
            headers["Authorization"] = auth_header
        try:
            resp = requests.post(
                f"{api_url}/chats/{request.chat_id}/answers",
                json={
                    "question_id": request.question_id,
                    "content": answer,
                    "participant_id": request.participant_id
                },
                headers=headers if headers else None
            )
            resp.raise_for_status()
        except Exception as e:
            print(f"[WARN] Failed to POST answer to API: {e}")

    return QueryResponse(answer=answer)
