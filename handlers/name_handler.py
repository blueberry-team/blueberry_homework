from fastapi import HTTPException
from dtos.response.name_res_dto import NameResDTO, NameListResDTO
from usecases.name_usecase import NameUseCase

from constants.error_response import ERROR_RESPONSES
from dtos.request.name_req_dto import NameReqDTO


class NameHandler:
    def get_names(self):
        try:
            name_list = NameUseCase.get_names()
            # 이름 목록을 NameListResDTO로 변환
            return NameListResDTO(names=[NameResDTO(name=n.name) for n in name_list])
        except HTTPException as e:
            raise
        except Exception as e:
            # 예상치 못한 오류가 발생한 경우
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    def delete_name_by_index(self, index: int):
        try:
            deleted_name = NameUseCase.delete_name_by_index(index)
            return {"message": "이름이 삭제되었습니다", "deleted_name": deleted_name}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    def delete_name_by_name(self, name: str):
        try:
            deleted_name = NameUseCase.delete_name_by_name(name)
            return {"message": "이름이 삭제되었습니다", "deleted_name": deleted_name}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    def get_name_by_name(self, name: str):
        try:
            user = NameUseCase.get_name_by_name(name)
            if not user:
                raise HTTPException(status_code=404, detail="존재하지 않는 이름입니다.")
            return NameResDTO(name=user.name)
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
