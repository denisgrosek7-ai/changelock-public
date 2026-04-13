from fastapi import FastAPI
app = FastAPI(title="ChangeLock Policy Engine", version="0.1.0")

@app.post("/evaluate/change")
def evaluate_change(payload: dict):
    return {"decision": "ALLOW", "reasons": []}

@app.post("/evaluate/artifact")
def evaluate_artifact(payload: dict):
    return {"decision": "ALLOW", "reasons": []}
