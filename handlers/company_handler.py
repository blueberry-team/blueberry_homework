from fastapi import HTTPException

from constants.error_response import ERROR_RESPONSES
from dtos.request.company_req_dto import CompanyReqDTO
from usecases.company_usecase import CompanyUseCase

class CompanyHandler:
    def create_company(self, company_req: CompanyReqDTO):
        try:
            CompanyUseCase.create_company(company_req)
            return {"message": "회사가 추가되었습니다"}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
            
    def get_companies(self):
        try:
            companies = CompanyUseCase.get_companies()
            if not companies:
                return {"message": "등록된 회사가 없습니다", "data": []}
            return {"message": "회사 목록을 가져왔습니다", "data": companies}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
            
    def get_company_by_name(self, name: str):
        try:
            company = CompanyUseCase.get_company_by_name(name)
            return {"message": "회사가 조회되었습니다", "data": [company]}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, 
                detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e 