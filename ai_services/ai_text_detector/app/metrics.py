import re

AI_MARKERS = [
    r"\bAs an AI language model\b",
    r"\bI cannot provide\b",
    r"\bI don't have access\b",
    r"\bIn conclusion\b",
    r"\bOverall\b",
]


def detect_ai_generated(text: str) -> tuple[bool, float, str]:
    text_lower = text.lower()
    score = 0.0
    reasons = []

    # Длинный, формальный текст
    if len(text.split()) > 120:
        score += 0.2
        reasons.append("long_form_structure")

    # Высокая плотность запятых (типично для LLM)
    if text.count(",") / max(len(text), 1) > 0.03:
        score += 0.2
        reasons.append("comma_density")

    # Типичные фразы ИИ
    for pattern in AI_MARKERS:
        if re.search(pattern.lower(), text_lower):
            score += 0.3
            reasons.append("ai_phrase_match")

    score = min(score, 1.0)
    is_ai = score >= 0.5

    explanation = (
        ", ".join(reasons)
        if reasons
        else "no strong ai patterns detected"
    )

    return is_ai, score, explanation