from tmp_database import tmp_db


class NameRepository:
    # 이름을 추가하는 함수
    @staticmethod
    def add_name(name: str):
        tmp_db.append(name)

    # 이름 목록을 가져오는 함수
    @staticmethod
    def get_names():
        return tmp_db
