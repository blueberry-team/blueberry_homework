from fastapi import HTTPException

from constants.error_response import ERROR_RESPONSES
from dtos.request.name_req_dto import NameReqDTO
from usecases.name_usecase import NameUseCase


class NameHandler:
    def create_name(self, input_name: NameReqDTO):
        try:
            added_name = NameUseCase.create_name(input_name)
            return {"message": "이름이 추가되었습니다", "name": added_name }
        except HTTPException as e:
            # HTTP 예외는 그대로 다시 발생시켜 FastAPI가 처리하도록 함
            raise
        except Exception as e:
            # 예상치 못한 오류가 발생한 경우
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

    def get_names(self):
        try:
            name_list = NameUseCase.get_names()
            # 이름이 없는 경우
            if not name_list:
                return {"message": "등록된 이름이 없습니다", "names": name_list}
            # 성공!
            return {"message": "이름 목록을 가져왔습니다", "names": name_list}
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
            deleted_names = NameUseCase.delete_name_by_name(name)
            return {"message": "이름이 삭제되었습니다", "deleted_names": deleted_names}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
    
    def change_name(self, used_id: str, new_name: str):
        try:
            updated_name = NameUseCase.change_name(used_id, new_name)
            return {"message": "이름이 수정되었습니다", "updated_name": updated_name}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e
            
    def get_name_by_name(self, name: str):
        try:
            name_by_name = NameUseCase.get_name_by_name(name)
            return {"message": "이름이 조회되었습니다", "name_by_name": name_by_name}
        except HTTPException as e:
            raise
        except Exception as e:
            raise HTTPException(
                status_code=500, detail=ERROR_RESPONSES["SERVER_ERROR"]
            ) from e

