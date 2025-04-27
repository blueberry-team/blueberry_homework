from typing import List, Optional
from entities.company_entity import CompanyEntity
from tmp_database import company_db


class CompanyRepository:
    @staticmethod
    def get_companies() -> List[CompanyEntity]:
        print(company_db)
        """모든 회사 목록을 반환합니다."""
        return company_db

    @staticmethod
    def add_company(company: CompanyEntity) -> CompanyEntity:
        """새로운 회사를 추가합니다."""
        company_db.append(company)
        return company

    @staticmethod
    def find_by_name(name: str) -> Optional[CompanyEntity]:
        """사용자 이름으로 회사를 찾습니다."""
        for company in company_db:
            if company.name == name:
                return company
        return None

    @staticmethod
    def get_company_by_id(company_id: str) -> Optional[CompanyEntity]:
        """회사 ID로 회사를 찾습니다."""
        for company in company_db:
            if company.id == company_id:
                return company
        return None
