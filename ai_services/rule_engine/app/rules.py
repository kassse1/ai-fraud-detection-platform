import re
from typing import Tuple, List

SCAM_KEYWORDS: List[str] = [
    "urgent",
    "verify",
    "account",
    "suspended",
    "click",
    "wallet",
    "crypto",
    "transfer",
    "prize",
]


def analyze_text(text: str) -> Tuple[float, str]:
    text = text.lower()
    score = 0.0
    reasons = []

    for kw in SCAM_KEYWORDS:
        if kw in text:
            score += 0.1
            reasons.append(f"keyword='{kw}'")

    if re.search(r"http[s]?://", text):
        score += 0.2
        reasons.append("contains_url")

    return min(score, 1.0), ", ".join(reasons)
