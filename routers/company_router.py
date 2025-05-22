from fastapi import APIRouter, Depends
from dtos.request.company_req_dto import (
    CreateCompanyReqDTO,
    ChangeCompanyReqDTO,
    DeleteCompanyReqDTO,
)
from dtos.response.company_res_dto import CompanyResDTO, CompanyListResDTO
from handlers.company_handler import CompanyHandler
from typing import List

company_router = APIRouter(prefix="/companies", tags=["companies"])


# 의존성 주입을 위한 함수
def get_company_handler():
    return CompanyHandler()


@company_router.post("/", response_model=CompanyResDTO)
def create_company(
    company_req: CreateCompanyReqDTO,
    handler: CompanyHandler = Depends(get_company_handler),
):
    return handler.create_company(company_req)


@company_router.get("/", response_model=CompanyListResDTO)
def get_companies(handler: CompanyHandler = Depends(get_company_handler)):
    return handler.get_companies()


@company_router.get("/{id}", response_model=CompanyResDTO)
def get_company_by_id(id: str, handler: CompanyHandler = Depends(get_company_handler)):
    return handler.get_company_by_id(id)


@company_router.put("/{id}", response_model=CompanyResDTO)
def change_company(
    id: str,
    change_req: ChangeCompanyReqDTO,
    handler: CompanyHandler = Depends(get_company_handler),
):
    return handler.change_company(id, change_req)


@company_router.delete("/{id}", response_model=CompanyResDTO)
def delete_company(id: str, handler: CompanyHandler = Depends(get_company_handler)):
    return handler.delete_company(id)
