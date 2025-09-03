from sentence_transformers import SentenceTransformer
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List

app = FastAPI()
model = SentenceTransformer('sonoisa/sentence-bert-base-ja-mean-tokens')

class EmbeddingRequest(BaseModel):
    texts: List[str]

class EmbeddingResponse(BaseModel):
    embeddings: List[List[float]]

@app.post("/embed", response_model=EmbeddingResponse)
def embed(req: EmbeddingRequest):
    try:
        embs = model.encode(req.texts, convert_to_numpy=True)
        return EmbeddingResponse(embeddings=embs.tolist())
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
