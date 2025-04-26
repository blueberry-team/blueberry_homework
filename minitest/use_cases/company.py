from ..domain.company import Company
from ..repositores.company_repository import CompanyRepository
from ..repositores.name_repository import NameRepository
from datetime import datetime
from django.core.validators import ValidationError # type: ignore
import uuid

company_repo = CompanyRepository()

def get_company(idx=None):
    company = company_repo.get_company(idx)
    return company

def create_company(name, company_name):
    name_repo = NameRepository()
    if not name_repo.find_by_name(name):
        raise ValidationError('user does not exist') 
    if company_repo.find_by_name(name):
        raise ValidationError('user already has a company')
    company = Company(id=uuid.uuid4(), name=name, company_name=company_name, created_at=datetime.now())
    print(company)
    # try:
    #     company.validate() 
    # except ValueError as e:
    #     raise ValidationError(str(e))
    print('hi')
    company = company_repo.create_company(company)
    print('hi2')
    return company

