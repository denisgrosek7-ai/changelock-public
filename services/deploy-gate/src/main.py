from fastapi import FastAPI
app = FastAPI(title="ChangeLock Deploy Gate", version="0.1.0")

@app.post("/admission/review")
def admission_review(payload: dict):
    return {"allowed": False, "status": {"message": "implement admission decision"}}
