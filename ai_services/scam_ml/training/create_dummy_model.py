from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.linear_model import LogisticRegression
import joblib
from pathlib import Path

# Минимальный тренировочный набор
texts = [
    "urgent verify your account",
    "click here to reset password",
    "you won a prize",
    "hello how are you",
    "let's meet tomorrow",
    "this is a normal message"
]

labels = [1, 1, 1, 0, 0, 0]  # 1 = scam, 0 = normal

vectorizer = TfidfVectorizer()
X = vectorizer.fit_transform(texts)

model = LogisticRegression()
model.fit(X, labels)

pipeline = {
    "vectorizer": vectorizer,
    "model": model
}

BASE_DIR = Path(__file__).resolve().parent.parent
MODEL_PATH = BASE_DIR / "models" / "scam_model.pkl"

MODEL_PATH.parent.mkdir(exist_ok=True)
joblib.dump(pipeline, MODEL_PATH)

print(f"Dummy model saved to {MODEL_PATH}")
