from pydantic import BaseModel, Field

class CompanyReqDTO(BaseModel):
    name: str = Field(..., min_length=1, max_length=50)
    company_name: str = Field(..., min_length=1, max_length=100) 