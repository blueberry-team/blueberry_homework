from ..domain.company import Company
from ..repositores.company_repository import CompanyRepository
from ..repositores.name_repository import NameRepository
from datetime import datetime
from django.core.validators import ValidationError 
import uuid

company_repo = CompanyRepository()

def get_company(idx=None):
    company = company_repo.get_company(idx)
    return company

def create_company(user_id, company_name, company_address, total_staff):
    name_repo = NameRepository()
    user = name_repo.get_user_by_uuid(user_id)
    if user.role != 'boss':
        raise ValidationError('사용자는 관리자여야 합니다')
    if company_repo.find_by_user_id(user_id):
        raise ValidationError('사용자는 이미 회사에 속해 있습니다')
    company = Company(
        id=uuid.uuid4(), 
        user_id=user, 
        company_name=company_name, 
        company_address=company_address,
        total_staff=total_staff,
        created_at=datetime.now(),
        updated_at=None
        )
    company = company_repo.create_company(company, user)
    return company

def change_company(uuid, company_name, company_address, total_staff):
    company = company_repo.change_info(uuid, company_name, company_address, total_staff)
    company.updated_at = datetime.now()
    return company

def delete_company_by_index(idx):
    company = company_repo.delete_index(idx)
    return company

