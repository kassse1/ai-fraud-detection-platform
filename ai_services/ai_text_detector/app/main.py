from ai_services.common.base_service import create_base_app
from ai_services.common.schemas import TextRequest
from pydantic import BaseModel
from ai_services.ai_text_detector.app.metrics import detect_ai_generated

app = create_base_app("AI Text Detector Service")


class AIDetectResponse(BaseModel):
    is_ai_generated: bool
    score: float
    explanation: str


@app.post("/analyze", response_model=AIDetectResponse)
def analyze(req: TextRequest):
    is_ai, score, explanation = detect_ai_generated(req.text)
    return AIDetectResponse(
        is_ai_generated=is_ai,
        score=score,
        explanation=explanation,
    )
