from sentence_transformers import SentenceTransformer
from fastapi import FastAPI, HTTPException, Request
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse
from pydantic import BaseModel
from typing import List
import json

app = FastAPI()
model = SentenceTransformer('sonoisa/sentence-bert-base-ja-mean-tokens')

class EmbeddingRequest(BaseModel):
    texts: List[str]

class EmbeddingResponse(BaseModel):
    embeddings: List[List[float]]

@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request: Request, exc: RequestValidationError):
    print(f"[EMBEDDING] Validation error: {exc}")
    print(f"[EMBEDDING] Request body: {await request.body()}")
    return JSONResponse(
        status_code=422,
        content={"detail": "Validation error", "errors": exc.errors()}
    )

@app.get("/health")
def health_check():
    return {"status": "ok"}

@app.post("/embed", response_model=EmbeddingResponse)
async def embed(request: Request, req: EmbeddingRequest):
    print(f"[EMBEDDING] Received request: {req}")
    print(f"[EMBEDDING] texts: {req.texts}")
    try:
        embs = model.encode(req.texts, convert_to_numpy=True)
        return EmbeddingResponse(embeddings=embs.tolist())
    except Exception as e:
        print(f"[EMBEDDING] error: {e}")
        raise HTTPException(status_code=500, detail=str(e))
