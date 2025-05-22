from datetime import datetime
import uuid
from fastapi import HTTPException

from constants.error_response import ERROR_RESPONSES
from dtos.request.name_req_dto import NameReqDTO
from entities.user_entity import UserEntity
from repositories.name_repository import NameRepository


class NameUseCase:
    @staticmethod
    def get_names():
        # 이름을 레포지토리를 사용해서 가져오기
        try:
            name_list = NameRepository.get_names()
            return name_list  # 핸들러에서 NameListResDTO로 변환
        except Exception as e:
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
    def get_name_by_name(name: str):
        name_list = NameRepository.get_names()
        if name not in [item.name for item in name_list]:
            raise HTTPException(
                status_code=404, detail=ERROR_RESPONSES["USER_NOT_FOUND"]
            )
        try:
            user = NameRepository.find_by_name(name)
            return user  # 핸들러에서 NameResDTO로 변환
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
