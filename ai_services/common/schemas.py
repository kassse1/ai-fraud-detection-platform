from pydantic import BaseModel


class TextRequest(BaseModel):
    text: str


class ScoreResponse(BaseModel):
    score: float
    explanation: str | None = None
