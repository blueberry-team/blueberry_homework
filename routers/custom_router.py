from fastapi import APIRouter, Depends
from dtos.request.name_req_dto import NameReqDTO
from handlers.name_handler import NameHandler

custom_router = APIRouter(prefix="/names", tags=["names"])


# 의존성 주입을 위한 함수
def get_name_handler():
    return NameHandler()


@custom_router.post("/")
def create_name(
    input_name: NameReqDTO, handler: NameHandler = Depends(get_name_handler)
):
    return handler.create_name(input_name)


@custom_router.get("/")
def get_names(handler: NameHandler = Depends(get_name_handler)):
    return handler.get_names()


@custom_router.delete("/index/{index}")
def delete_name_by_index(index: int, handler: NameHandler = Depends(get_name_handler)):
    return handler.delete_name_by_index(index)

@custom_router.delete("/name/{name}")
def delete_name_by_name(name: str, handler: NameHandler = Depends(get_name_handler)):
    return handler.delete_name_by_name(name)

