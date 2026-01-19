from ai_services.common.base_service import create_base_app
from ai_services.common.schemas import TextRequest, ScoreResponse
from ai_services.scam_ml.app.inference import analyze_text

app = create_base_app("Scam ML Service")


@app.post("/analyze", response_model=ScoreResponse)
def analyze(req: TextRequest):
    score = analyze_text(req.text)
    return ScoreResponse(score=score)
