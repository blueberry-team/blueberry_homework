from datetime import datetime
import uuid
from fastapi import HTTPException

from constants.error_response import ERROR_RESPONSES
from dtos.request.name_req_dto import NameReqDTO
from entities.user_entity import UserEntity
from repositories.name_repository import NameRepository


class NameUseCase:
    @staticmethod
    def create_name(input_name: NameReqDTO):
        name_list = NameRepository.get_names()
        # 이름이 이미 존재하는 경우
        if input_name.name in [item.name for item in name_list]:
            raise HTTPException(
                status_code=400, detail=ERROR_RESPONSES["DUPLICATE_NAME"]
            )

        # Request DTO를 Entity로 변환
        try:
            # 현재 시간 가져오기
            current_time = datetime.now()

            name_entity = UserEntity(
                id=str(uuid.uuid4()),
                name=input_name.name,
                created_at=current_time,
                updated_at=current_time,
            )
            added_name = NameRepository.add_name(name_entity)
            return added_name

        except Exception as e:
            # 예상치 못한 오류가 발생한 경우
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
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
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    @staticmethod
    def delete_name_by_index(index: int):
        name_list = NameRepository.get_names()
        if index < 0 or index >= len(name_list):
            raise HTTPException(
                status_code=400, detail=ERROR_RESPONSES["INVALID_INDEX"]
            )
        try:
            deleted_name = name_list[index]
            NameRepository.delete_name_by_index(index)
            return deleted_name
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    @staticmethod
    def delete_name_by_name(name: str):
        try:
            deleted_name = NameRepository.delete_name_by_name(name)
            # 해당 이름을 찾을 수 없는 경우
            if not deleted_name:
                raise HTTPException(
                    status_code=404, detail=ERROR_RESPONSES["USER_NOT_FOUND"]
                )
            return deleted_name
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    @staticmethod
    def change_name(user_id: str, new_name: str):
        name_list = NameRepository.get_names()

        # 해당 ID의 사용자 찾기
        user_to_update = None
        user_index = -1

        for idx, user in enumerate(name_list):
            if user.id == user_id:
                user_to_update = user
                user_index = idx
                break

        # 사용자를 찾지 못한 경우
        if not user_to_update:
            raise HTTPException(
                status_code=404, detail=ERROR_RESPONSES["USER_NOT_FOUND"]
            )

        # 이름이 변경되지 않은 경우 (자기 자신과의 중복)
        if user_to_update.name == new_name:
            raise HTTPException(
                status_code=400, detail=ERROR_RESPONSES["DUPLICATE_NAME"]
            )

        # 다른 사용자와 이름 중복 체크
        for user in name_list:
            if user.id != user_id and user.name == new_name:
                raise HTTPException(
                    status_code=400, detail=ERROR_RESPONSES["DUPLICATE_NAME"]
                )

        try:
            # 이름 변경 및 updatedAt 갱신
            updated_user = UserEntity(
                id=user_to_update.id,
                name=new_name,
                created_at=user_to_update.created_at,
                updated_at=datetime.now(),  # 현재 시간으로 업데이트
            )

            # 레포지토리에 업데이트
            NameRepository.change_name(user_index, updated_user)

            return updated_user
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    @staticmethod
    def get_name_by_name(name: str):
        name_list = NameRepository.get_names()
        if name not in [item.name for item in name_list]:
            raise HTTPException(
                status_code=404, detail=ERROR_RESPONSES["USER_NOT_FOUND"]
            )
        try:
            name_list = NameRepository.find_by_name(name)
            return name_list
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
