from ai_services.scam_ml.app.model import ScamModel

model = ScamModel()


def analyze_text(text: str) -> float:
    return model.predict(text)
