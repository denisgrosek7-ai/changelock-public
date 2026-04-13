from fastapi import FastAPI
app = FastAPI(title="ChangeLock Attestation Verifier", version="0.1.0")

@app.post("/verify/github-attestation")
def verify_github_attestation(payload: dict):
    return {"verified": False, "reasons": ["implement offline or API verification"]}

@app.post("/verify/cosign")
def verify_cosign(payload: dict):
    return {"verified": False, "reasons": ["implement cosign verify wrapper"]}
