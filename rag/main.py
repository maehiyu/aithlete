
from fastapi import FastAPI, Request
from pydantic import BaseModel
from typing import Optional
from langchain.llms import OpenAI
import os


app = FastAPI()

# OpenAI APIキーは環境変数から取得
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
llm = OpenAI(openai_api_key=OPENAI_API_KEY)


class QueryRequest(BaseModel):
    question: str
    user_id: Optional[str] = None


class QueryResponse(BaseModel):
    answer: str


@app.get("/")
async def root():
    return {"message": "Hello, FastAPI!"}


@app.post("/rag/query", response_model=QueryResponse)
async def rag_query(request: QueryRequest):
    # LangChainでquestionからanswerを生成
    answer = llm(request.question)
    return QueryResponse(answer=answer)
