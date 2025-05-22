from fastapi import HTTPException
from usecases.auth_usecase import AuthUseCase


class AuthHandler:
    def sign_up(self, user_req):
        return AuthUseCase.sign_up(user_req)

    def log_in(self, login_req):
        return AuthUseCase.log_in(login_req)

    def change_user(self, change_req):
        return AuthUseCase.change_user(change_req)

    def get_user(self, user_id):
        return AuthUseCase.get_user(user_id)
