from fastapi import APIRouter, Depends
from dtos.request.company_req_dto import CompanyReqDTO
from handlers.company_handler import CompanyHandler

company_router = APIRouter(prefix="/companies", tags=["companies"])


# 의존성 주입을 위한 함수
def get_company_handler():
    return CompanyHandler()


@company_router.post("/")
def create_company(
    company_req: CompanyReqDTO, handler: CompanyHandler = Depends(get_company_handler)
):
    return handler.create_company(company_req)


@company_router.get("/")
def get_companies(handler: CompanyHandler = Depends(get_company_handler)):
    return handler.get_companies()


@company_router.get("/{name}")
def get_company_by_name(
    name: str, handler: CompanyHandler = Depends(get_company_handler)
):
    return handler.get_company_by_name(name)
