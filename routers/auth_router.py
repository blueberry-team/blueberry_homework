from fastapi import APIRouter, Depends
from handlers.auth_handler import AuthHandler
from dtos.request.auth_req_dto import SignUpReqDTO, LogInReqDTO, ChangeUserReqDTO
from dtos.response.auth_res_dto import AuthResDTO, UserResDTO

auth_router = APIRouter(prefix="/auth", tags=["auth"])


def get_auth_handler():
    return AuthHandler()


@auth_router.post("/sign-up", response_model=AuthResDTO)
def sign_up(user_req: SignUpReqDTO, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.sign_up(user_req)


@auth_router.post("/log-in", response_model=AuthResDTO)
def log_in(login_req: LogInReqDTO, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.log_in(login_req)


@auth_router.put("/change-user", response_model=AuthResDTO)
def change_user(
    change_req: ChangeUserReqDTO, handler: AuthHandler = Depends(get_auth_handler)
):
    return handler.change_user(change_req)


@auth_router.get("/get-user/{user_id}", response_model=UserResDTO)
def get_user(user_id: str, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.get_user(user_id)
