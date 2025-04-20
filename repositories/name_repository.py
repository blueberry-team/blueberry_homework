from typing import List
from entities.name_entity import NameEntity
from tmp_database import tmp_db


class NameRepository:
    # 이름을 추가하는 함수
    @staticmethod
    def add_name(name_entity: NameEntity) -> NameEntity:
        tmp_db.append(name_entity)
        return name_entity

    # 이름 목록을 가져오는 함수
    @staticmethod
    def get_names() -> List[NameEntity]:
        return tmp_db

    @staticmethod
    def delete_name_by_index(index: int) -> NameEntity:
        return tmp_db.pop(index)

    # TODO : 현재 상태로는 전체 DB를 조회함 추후에 개선 필요
    @staticmethod
    def delete_name_by_name(name: str) -> NameEntity:
        for item in tmp_db:
            if item.name == name:
                return tmp_db.pop(item)
        return None
