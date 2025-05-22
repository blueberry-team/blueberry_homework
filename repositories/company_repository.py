from typing import List, Optional
from entities.company_entity import CompanyEntity
from db.db import ScyllaDB
import uuid


class CompanyRepository:
    @staticmethod
    def get_companies() -> List[CompanyEntity]:
        session = ScyllaDB.get_session()
        query = "SELECT id, name, company_name, created_at FROM companies"
        rows = session.execute(query)
        return [CompanyEntity(
            id=str(row["id"]),
            name=row["name"],
            company_name=row.get("company_name", ""),
            created_at=row["created_at"]
        ) for row in rows]

    @staticmethod
    def add_company(company: CompanyEntity) -> CompanyEntity:
        session = ScyllaDB.get_session()
        query = """
        INSERT INTO companies (id, name, company_name, created_at)
        VALUES (%s, %s, %s, %s)
        """
        session.execute(query, (
            uuid.UUID(company.id),
            company.name,
            company.company_name,
            company.created_at
        ))
        return company

    @staticmethod
    def find_by_name(name: str) -> Optional[CompanyEntity]:
        session = ScyllaDB.get_session()
        query = "SELECT id, name, company_name, created_at FROM companies WHERE name = %s ALLOW FILTERING"
        rows = session.execute(query, (name,))
        for row in rows:
            return CompanyEntity(
                id=str(row["id"]),
                name=row["name"],
                company_name=row.get("company_name", ""),
                created_at=row["created_at"]
            )
        return None

    @staticmethod
    def get_company_by_id(company_id: str) -> Optional[CompanyEntity]:
        session = ScyllaDB.get_session()
        query = "SELECT id, name, company_name, created_at FROM companies WHERE id = %s"
        rows = session.execute(query, (uuid.UUID(company_id),))
        for row in rows:
            return CompanyEntity(
                id=str(row["id"]),
                name=row["name"],
                company_name=row.get("company_name", ""),
                created_at=row["created_at"]
            )
        return None
