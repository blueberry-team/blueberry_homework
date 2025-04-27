from typing import List
from entities.user_entity import UserEntity
from tmp_database import name_db


class NameRepository:
    # 이름을 추가하는 함수
    @staticmethod
    def add_name(name_entity: UserEntity) -> UserEntity:
        name_db.append(name_entity)
        return name_db

    # 이름 목록을 가져오는 함수
    @staticmethod
    def get_names() -> List[UserEntity]:
        return name_db

    @staticmethod
    def delete_name_by_index(index: int) -> UserEntity:
        return name_db.pop(index)

    # TODO : 현재 상태로는 전체 DB를 조회함 추후에 개선 필요
    @staticmethod
    def delete_name_by_name(name: str) -> UserEntity:
        for index, item in enumerate(name_db):
            if item.name == name:
                return name_db.pop(index)
        return None

    @staticmethod
    def change_name(index: int, user_entity: UserEntity) -> UserEntity:
        name_db[index] = user_entity
        return user_entity
    
    @staticmethod
    def find_by_name(name: str) -> UserEntity:
        for item in name_db:
            if item.name == name:
                return item
        return None
