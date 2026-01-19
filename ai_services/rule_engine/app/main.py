from ai_services.common.base_service import create_base_app
from ai_services.common.schemas import TextRequest, ScoreResponse
from ai_services.rule_engine.app.rules import analyze_text

app = create_base_app("Rule Engine Service")


@app.post("/analyze", response_model=ScoreResponse)
def analyze(req: TextRequest):
    score, explanation = analyze_text(req.text)
    return ScoreResponse(score=score, explanation=explanation)
