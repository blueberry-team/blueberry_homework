from fastapi import APIRouter, Depends
from handlers.auth_handler import AuthHandler

auth_router = APIRouter(prefix="/auth", tags=["auth"])


def get_auth_handler():
    return AuthHandler()


@auth_router.post("/sign-up")
def sign_up(user_req: dict, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.sign_up(user_req)


@auth_router.post("/log-in")
def log_in(login_req: dict, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.log_in(login_req)


@auth_router.put("/change-user")
def change_user(change_req: dict, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.change_user(change_req)


@auth_router.get("/get-user/{user_id}")
def get_user(user_id: str, handler: AuthHandler = Depends(get_auth_handler)):
    return handler.get_user(user_id)
