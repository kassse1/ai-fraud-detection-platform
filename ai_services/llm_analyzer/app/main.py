from ai_services.common.base_service import create_base_app
from ai_services.common.schemas import TextRequest, ScoreResponse
from ai_services.llm_analyzer.app.analyzer import analyze_text_semantics

app = create_base_app("LLM Analyzer Service")


@app.post("/analyze", response_model=ScoreResponse)
def analyze(req: TextRequest):
    score, explanation = analyze_text_semantics(req.text)
    return ScoreResponse(score=score, explanation=explanation)
