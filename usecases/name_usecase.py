from datetime import datetime
from fastapi import HTTPException

from dtos.request.name_req_dto import NameReqDTO
from entities.name_entity import NameEntity
from repositories.name_repository import NameRepository


class NameUseCase:
    @staticmethod
    def create_name(input_name: NameReqDTO):
        print(input_name)
        print(input_name.name)
        name_list = NameRepository.get_names()
        # 이름이 이미 존재하는 경우
        if input_name.name in [item.name for item in name_list]:
            raise HTTPException(status_code=400, detail="이름이 이미 존재합니다")

        # Request DTO를 Entity로 변환
        try:
            print(datetime.now())
            name_entity = NameEntity(name=input_name.name, created_at=datetime.now())
            added_name_entity = NameRepository.add_name(name_entity)
            return {"name": added_name_entity.name}
        except Exception as e:
            # 예상치 못한 오류가 발생한 경우
            raise HTTPException(
                status_code=500, detail="서버 오류가 발생했습니다"
            ) from e

    @staticmethod
    def get_names():
        # 이름을 레포지토리를 사용해서 가져오기
        try:
            name_list = NameRepository.get_names()
            return name_list
        except Exception as e:
            # 예상치 못한 오류가 발생한 경우
            raise HTTPException(
                status_code=500, detail="서버 오류가 발생했습니다"
            ) from e

    @staticmethod
    def delete_name_by_index(index: int):
        name_list = NameRepository.get_names()
        if index < 0 or index >= len(name_list):
            raise HTTPException(status_code=400, detail="유효하지 않은 인덱스입니다")
        try:
            deleted_name = name_list[index]
            NameRepository.delete_name_by_index(index)
            return deleted_name
        except Exception as e:
            raise HTTPException(
                status_code=500, detail="서버 오류가 발생했습니다"
            ) from e

    @staticmethod
    def delete_name_by_name(name: str):
        try:
            deleted_name = NameRepository.delete_name_by_name(name)
            # 해당 이름을 찾을 수 없는 경우
            if not deleted_name:
                raise HTTPException(
                    status_code=404, detail="해당 이름을 찾을 수 없습니다"
                )
            return deleted_name
        except Exception as e:
            raise HTTPException(
                status_code=500, detail="서버 오류가 발생했습니다"
            ) from e
