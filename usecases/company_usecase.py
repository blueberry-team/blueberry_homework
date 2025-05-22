import uuid
from datetime import datetime
from fastapi import HTTPException

from constants.error_response import ERROR_RESPONSES
from dtos.request.company_req_dto import CreateCompanyReqDTO, ChangeCompanyReqDTO
from dtos.response.company_res_dto import CompanyResDTO, CompanyListResDTO
from entities.company_entity import CompanyEntity
from repositories.company_repository import CompanyRepository
from repositories.name_repository import NameRepository
from repositories.auth_repository import AuthRepository


class CompanyUseCase:
    @staticmethod
    def create_company(company_req: CreateCompanyReqDTO) -> CompanyResDTO:
        """회사 생성 (boss만 가능)"""
        # 1. 유저 조회 및 권한 체크
        user = AuthRepository.find_by_id(company_req.user_id)
        if not user or user.role != "boss":
            raise HTTPException(
                status_code=403, detail="회사 생성은 boss만 가능합니다."
            )
        # 2. 회사 생성
        now = datetime.now()
        company = CompanyEntity(
            id=str(uuid.uuid4()),
            user_id=company_req.user_id,
            company_name=company_req.company_name,
            company_address=company_req.company_address or "",
            total_staff=company_req.total_staff or 0,
            created_at=now,
            updated_at=now,
        )
        created = CompanyRepository.add_company(company)
        return CompanyResDTO(**created.dict())

    @staticmethod
    def get_companies() -> CompanyListResDTO:
        """회사 전체 목록 조회"""
        companies = CompanyRepository.get_companies()
        return CompanyListResDTO(
            companies=[CompanyResDTO(**c.dict()) for c in companies]
        )

    @staticmethod
    def get_company_by_id(id: str) -> CompanyResDTO:
        """회사 단일 조회"""
        company = CompanyRepository.get_company_by_id(id)
        if not company:
            raise HTTPException(status_code=404, detail="존재하지 않는 회사입니다.")
        return CompanyResDTO(**company.dict())

    @staticmethod
    def change_company(id: str, change_req: ChangeCompanyReqDTO) -> CompanyResDTO:
        """회사 정보 수정"""
        company = CompanyRepository.get_company_by_id(id)
        if not company:
            raise HTTPException(status_code=404, detail="존재하지 않는 회사입니다.")
        # 변경 필드만 반영
        if change_req.company_name:
            company.company_name = change_req.company_name
        if change_req.company_address:
            company.company_address = change_req.company_address
        if change_req.total_staff is not None:
            company.total_staff = change_req.total_staff
        company.updated_at = datetime.now()
        updated = CompanyRepository.update_company(company)
        return CompanyResDTO(**updated.dict())

    @staticmethod
    def delete_company(id: str) -> CompanyResDTO:
        """회사 삭제"""
        company = CompanyRepository.get_company_by_id(id)
        if not company:
            raise HTTPException(status_code=404, detail="존재하지 않는 회사입니다.")
        deleted = CompanyRepository.delete_company(id)
        return CompanyResDTO(**company.dict())

    @staticmethod
    def get_company_by_name(name: str):
        """사용자 이름으로, 해당 사용자의 회사를 찾습니다."""
        try:
            company = CompanyRepository.find_by_name(name)
            if not company:
                raise HTTPException(
                    status_code=404, detail=ERROR_RESPONSES["COMPANY_NOT_FOUND"]
                )
            return company
        except HTTPException:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
