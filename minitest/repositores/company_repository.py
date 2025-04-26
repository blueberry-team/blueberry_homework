from .base_repository import BaseRepository
from ..tmp_database import tmp_company_db
from ..models import Company
from django.core.validators import ValidationError # type: ignore

class CompanyRepository(BaseRepository):
    def __init__(self):
        super().__init__(Company)

    def get_company(self, idx=None):
        return tmp_company_db
    
    def create_company(self, company):    
        tmp_company_db.append(company)    
        return tmp_company_db
    
    def find_by_name(self, name: str):
        for company in tmp_company_db:
            if company.name==name:
                return True
        return False