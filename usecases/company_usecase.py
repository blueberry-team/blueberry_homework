import uuid
from datetime import datetime
from fastapi import HTTPException

from constants.error_response import ERROR_RESPONSES
from dtos.request.company_req_dto import CompanyReqDTO
from entities.company_entity import CompanyEntity
from repositories.company_repository import CompanyRepository
from repositories.name_repository import NameRepository


class CompanyUseCase:
    @staticmethod
    def create_company(company_req: CompanyReqDTO):
        # 해당 이름의 사용자가 존재하는지 확인
        user = NameRepository.find_by_name(company_req.name)
        if not user:
            raise HTTPException(
                status_code=404, 
                detail=ERROR_RESPONSES["USER_NOT_FOUND"]
            )
            
        # 이미 회사를 가지고 있는지 확인
        existing_company = CompanyRepository.find_by_name(company_req.name)
        if existing_company:
            raise HTTPException(
                status_code=400, 
                detail=ERROR_RESPONSES["USER_ALREADY_HAS_COMPANY"]
            )
            
        try:
            # 새 회사 생성
            company = CompanyEntity(
                id=str(uuid.uuid4()),
                name=company_req.name,
                company_name=company_req.company_name,
                created_at=datetime.now()
            )
            
            # 저장소에 저장
            added_company = CompanyRepository.add_company(company)
            return added_company
            
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
    
    @staticmethod
    def get_companies():
        """모든 회사 목록을 반환합니다."""
        try:
            company_list = CompanyRepository.get_companies()
            print(company_list)
            return company_list
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
            
    @staticmethod
    def get_company_by_name(name: str):
        """사용자 이름으로, 해당 사용자의 회사를 찾습니다."""
        try:
            company = CompanyRepository.find_by_name(name)
            if not company:
                raise HTTPException(
                    status_code=404, 
                    detail=ERROR_RESPONSES["COMPANY_NOT_FOUND"]
                )
            return company
        except HTTPException:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e 