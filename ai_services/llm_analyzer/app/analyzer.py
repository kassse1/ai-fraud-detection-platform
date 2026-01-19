from ai_services.llm_analyzer.app.prompt import INDICATORS


def analyze_text_semantics(text: str) -> tuple[float, str]:
    text = text.lower()

    score = 0.0
    reasons = []

    for phrase in INDICATORS:
        if phrase in text:
            score += 0.15
            reasons.append(f"semantic_indicator='{phrase}'")

    score = min(score, 1.0)

    explanation = (
        ", ".join(reasons)
        if reasons
        else "no strong semantic manipulation detected"
    )

    return score, explanation
