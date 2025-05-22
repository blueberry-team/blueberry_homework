from fastapi import APIRouter, Depends
from dtos.request.name_req_dto import NameReqDTO
from handlers.name_handler import NameHandler

name_router = APIRouter(prefix="/names", tags=["names"])


# 의존성 주입을 위한 함수
def get_name_handler():
    return NameHandler()


@name_router.post("/")
def create_name(
    input_name: NameReqDTO, handler: NameHandler = Depends(get_name_handler)
):
    return handler.create_name(input_name)


@name_router.get("/")
def get_names(handler: NameHandler = Depends(get_name_handler)):
    return handler.get_names()


@name_router.get("/{name}")
def get_name_by_name(name: str, handler: NameHandler = Depends(get_name_handler)):
    return handler.get_name_by_name(name)


@name_router.put("/{used_id}")
def change_name(
    used_id: str, new_name: str, handler: NameHandler = Depends(get_name_handler)
):
    return handler.change_name(used_id, new_name)


@name_router.delete("/index/{index}")
def delete_name_by_index(index: int, handler: NameHandler = Depends(get_name_handler)):
    return handler.delete_name_by_index(index)


@name_router.delete("/name/{name}")
def delete_name_by_name(name: str, handler: NameHandler = Depends(get_name_handler)):
    return handler.delete_name_by_name(name)
