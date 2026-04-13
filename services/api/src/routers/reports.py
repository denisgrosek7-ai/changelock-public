from fastapi import APIRouter
router = APIRouter()

@router.get("/reports/{report_id}")
def get_report(report_id: str):
    return {"report_id": report_id, "status": "pending"}
