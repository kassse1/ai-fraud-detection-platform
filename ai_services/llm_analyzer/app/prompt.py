SYSTEM_PROMPT = """
You are an AI system designed to detect social engineering,
fraudulent intent, urgency manipulation and psychological pressure
in user messages.

Analyze the message and estimate the likelihood of fraud.
"""

INDICATORS = [
    "urgent",
    "act now",
    "verify immediately",
    "limited time",
    "account will be closed",
    "click the link",
]
