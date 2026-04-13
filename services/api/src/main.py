from fastapi import FastAPI
from .routers import health, decisions, reports

app = FastAPI(title="ChangeLock API", version="0.1.0")
app.include_router(health.router)
app.include_router(decisions.router, prefix="/v1")
app.include_router(reports.router, prefix="/v1")
