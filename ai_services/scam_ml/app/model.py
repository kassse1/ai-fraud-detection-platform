from pathlib import Path
import joblib

BASE_DIR = Path(__file__).resolve().parent.parent
MODEL_PATH = BASE_DIR / "models" / "scam_model.pkl"


class ScamModel:
    def __init__(self):
        if not MODEL_PATH.exists():
            raise FileNotFoundError(f"Model file not found: {MODEL_PATH}")

        pipeline = joblib.load(MODEL_PATH)
        self.vectorizer = pipeline["vectorizer"]
        self.model = pipeline["model"]

    def predict(self, text: str) -> float:
        X = self.vectorizer.transform([text])
        return float(self.model.predict_proba(X)[0][1])
