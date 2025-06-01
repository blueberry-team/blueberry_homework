from .base_repository import BaseRepository
from ..models import Company
from django.core.validators import ValidationError 

class CompanyRepository(BaseRepository):
    def __init__(self):
        super().__init__(Company)

    def get_company(self, idx=None):
        return list(self.model.objects.all())
    
    def create_company(self, company, user):    
        company = self.model(
            id=company.id,
            user_id=user,
            company_name=company.company_name,
            company_address=company.company_address,
            total_staff=company.total_staff,
            created_at=company.created_at,
            updated_at=company.updated_at
        )
        company.save()  
        return list(self.model.objects.all())
    
    def find_by_user_id(self, user_id: str):
        return self.model.objects.filter(user_id=user_id).exists()
    
    def change_info(self, uuid, company_name, company_address, total_staff):
        try:
            company = self.model.objects.get(id=uuid)
            company.company_name = company_name
            company.company_address = company_address
            company.total_staff = total_staff
            company.save()
            return company
        except self.model.DoesNotExist:
            raise ValidationError('회사가 존재하지 않습니다')
        
    def delete_index(self, idx):
        all_companies = list(self.model.objects.all())
        if idx >= len(all_companies):
            raise ValidationError('해당 인덱스에 값이 없습니다')
        
        company_to_delete = all_companies[idx]
        company_to_delete.delete()
        
        return list(self.model.objects.all())