from typing import List, Optional
from entities.user_entity import UserEntity
from db.db import ScyllaDB
from cassandra.query import dict_factory
import uuid


class NameRepository:
    # 이름 목록을 가져오는 함수
    @staticmethod
    def get_names() -> List[UserEntity]:
        session = ScyllaDB.get_session()
        query = "SELECT id, username, created_at, updated_at FROM users"
        rows = session.execute(query)
        return [
            UserEntity(
                id=str(row["id"]),
                name=row["username"],
                created_at=row["created_at"],
                updated_at=row["updated_at"],
            )
            for row in rows
        ]

    @staticmethod
    def delete_name_by_index(index: int) -> Optional[UserEntity]:
        # 인덱스 기반 삭제는 DB에서는 비효율적이므로, 전체 목록을 가져와 해당 인덱스의 id로 삭제
        names = NameRepository.get_names()
        if index < 0 or index >= len(names):
            return None
        user = names[index]
        session = ScyllaDB.get_session()
        query = "DELETE FROM users WHERE id = %s"
        session.execute(query, (uuid.UUID(user.id),))
        return user

    @staticmethod
    def delete_name_by_name(name: str) -> Optional[UserEntity]:
        # 이름으로 모두 삭제
        session = ScyllaDB.get_session()
        query = "SELECT id, username, created_at, updated_at FROM users WHERE username = %s ALLOW FILTERING"
        rows = session.execute(query, (name,))
        deleted = None
        for row in rows:
            del_query = "DELETE FROM users WHERE id = %s"
            session.execute(del_query, (row["id"],))
            deleted = UserEntity(
                id=str(row["id"]),
                name=row["username"],
                created_at=row["created_at"],
                updated_at=row["updated_at"],
            )
        return deleted

    @staticmethod
    def find_by_name(name: str) -> Optional[UserEntity]:
        session = ScyllaDB.get_session()
        query = "SELECT id, username, created_at, updated_at FROM users WHERE username = %s ALLOW FILTERING"
        rows = session.execute(query, (name,))
        for row in rows:
            return UserEntity(
                id=str(row["id"]),
                name=row["username"],
                created_at=row["created_at"],
                updated_at=row["updated_at"],
            )
        return None
