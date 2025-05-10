from .base_repository import BaseRepository
from ..tmp_database import tmp_company_db
from ..models import Company
from django.core.validators import ValidationError # type: ignore

class CompanyRepository(BaseRepository):
    def __init__(self):
        super().__init__(Company)

    def get_company(self, idx=None):
        return list(self.model.objects.all())
    
    def create_company(self, company):    
        company = self.model(
            id=company.id,
            name=company.name,
            company_name=company.company_name,
            created_at=company.created_at
        )
        company.save()  
        return list(self.model.objects.all())
    
    def find_by_name(self, name: str):
        return self.model.objects.filter(name=name).exists()