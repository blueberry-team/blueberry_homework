from fastapi import HTTPException
from dtos.request.company_req_dto import CreateCompanyReqDTO, ChangeCompanyReqDTO
from dtos.response.company_res_dto import CompanyResDTO, CompanyListResDTO
from usecases.company_usecase import CompanyUseCase


class CompanyHandler:
    def create_company(self, company_req: CreateCompanyReqDTO) -> CompanyResDTO:
        """회사 생성"""
        try:
            return CompanyUseCase.create_company(company_req)
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

    def get_companies(self) -> CompanyListResDTO:
        """회사 전체 목록 조회"""
        try:
            return CompanyUseCase.get_companies()
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

    def get_company_by_id(self, id: str) -> CompanyResDTO:
        """회사 단일 조회"""
        try:
            return CompanyUseCase.get_company_by_id(id)
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

    def change_company(self, id: str, change_req: ChangeCompanyReqDTO) -> CompanyResDTO:
        """회사 정보 수정"""
        try:
            return CompanyUseCase.change_company(id, change_req)
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

    def delete_company(self, id: str) -> CompanyResDTO:
        """회사 삭제"""
        try:
            return CompanyUseCase.delete_company(id)
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))
