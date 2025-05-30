from typing import List, Optional
from entities.company_entity import CompanyEntity
from db.db import ScyllaDB
import uuid
from datetime import datetime


class CompanyRepository:
    @staticmethod
    def add_company(company: CompanyEntity) -> CompanyEntity:
        """회사 추가"""
        session = ScyllaDB.get_session()
        query = """
        INSERT INTO companies (id, user_id, company_name, company_address, total_staff, created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        """
        session.execute(
            query,
            (
                uuid.UUID(company.id),
                company.user_id,
                company.company_name,
                company.company_address or "",
                company.total_staff or 0,
                company.created_at,
                company.updated_at,
            ),
        )
        return company

    @staticmethod
    def get_companies() -> List[CompanyEntity]:
        """회사 전체 목록 조회"""
        session = ScyllaDB.get_session()
        query = "SELECT * FROM companies"
        rows = session.execute(query)
        return [
            CompanyEntity(
                id=str(row["id"]),
                user_id=row["user_id"],
                company_name=row["company_name"],
                company_address=row.get("company_address", ""),
                total_staff=row.get("total_staff", 0),
                created_at=row["created_at"],
                updated_at=row.get("updated_at"),
            )
            for row in rows
        ]

    @staticmethod
    def get_company_by_id(id: str) -> Optional[CompanyEntity]:
        """회사 단일 조회"""
        session = ScyllaDB.get_session()
        query = "SELECT * FROM companies WHERE id = %s"
        rows = session.execute(query, (uuid.UUID(id),))
        for row in rows:
            return CompanyEntity(
                id=str(row["id"]),
                user_id=row["user_id"],
                company_name=row["company_name"],
                company_address=row.get("company_address", ""),
                total_staff=row.get("total_staff", 0),
                created_at=row["created_at"],
                updated_at=row.get("updated_at"),
            )
        return None

    @staticmethod
    def update_company(company: CompanyEntity) -> CompanyEntity:
        """회사 정보 수정"""
        session = ScyllaDB.get_session()
        query = """
        UPDATE companies SET user_id = %s, company_name = %s, company_address = %s, total_staff = %s, created_at = %s, updated_at = %s WHERE id = %s
        """
        session.execute(
            query,
            (
                company.user_id,
                company.company_name,
                company.company_address or "",
                company.total_staff or 0,
                company.created_at,
                company.updated_at,
                uuid.UUID(company.id),
            ),
        )
        return company

    @staticmethod
    def delete_company(id: str) -> bool:
        """회사 삭제"""
        session = ScyllaDB.get_session()
        query = "DELETE FROM companies WHERE id = %s"
        session.execute(query, (uuid.UUID(id),))
        return True

    @staticmethod
    def find_by_name(name: str) -> Optional[CompanyEntity]:
        """회사 이름으로 회사 조회 (색인용)"""
        session = ScyllaDB.get_session()
        query = "SELECT * FROM companies WHERE company_name = %s ALLOW FILTERING"
        rows = session.execute(query, (name,))
        for row in rows:
            return CompanyEntity(
                id=str(row["id"]),
                user_id=row["user_id"],
                company_name=row["company_name"],
                company_address=row.get("company_address", ""),
                total_staff=row.get("total_staff", 0),
                created_at=row["created_at"],
                updated_at=row.get("updated_at"),
            )
        return None
