from fastapi import APIRouter, Depends
from dtos.response.name_res_dto import NameResDTO, NameListResDTO
from handlers.name_handler import NameHandler

name_router = APIRouter(prefix="/names", tags=["names"])


# 의존성 주입을 위한 함수
def get_name_handler():
    return NameHandler()


@name_router.get("/", response_model=NameListResDTO)
def get_names(handler: NameHandler = Depends(get_name_handler)):
    return handler.get_names()


@name_router.get("/{name}", response_model=NameResDTO)
def get_name_by_name(name: str, handler: NameHandler = Depends(get_name_handler)):
    return handler.get_name_by_name(name)


@name_router.delete("/index/{index}")
def delete_name_by_index(index: int, handler: NameHandler = Depends(get_name_handler)):
    return handler.delete_name_by_index(index)


@name_router.delete("/name/{name}")
def delete_name_by_name(name: str, handler: NameHandler = Depends(get_name_handler)):
    return handler.delete_name_by_name(name)
