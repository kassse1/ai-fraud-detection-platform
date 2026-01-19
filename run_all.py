import subprocess
import sys
import time

services = [
    ("Rule Engine", "ai_services.rule_engine.app.main:app", 8001),
    ("Scam ML", "ai_services.scam_ml.app.main:app", 8002),
    ("LLM Analyzer", "ai_services.llm_analyzer.app.main:app", 8003),
    ("AI Text Detector", "ai_services.ai_text_detector.app.main:app", 8004),
]

processes = []

for name, app, port in services:
    print(f"🚀 Starting {name} on port {port}")
    p = subprocess.Popen(
        [
            sys.executable,
            "-m",
            "uvicorn",
            app,
            "--port",
            str(port),
        ]
    )
    processes.append(p)
    time.sleep(1)

print("\n✅ All AI services are running")

try:
    for p in processes:
        p.wait()
except KeyboardInterrupt:
    print("\n🛑 Stopping all services...")
    for p in processes:
        p.terminate()
