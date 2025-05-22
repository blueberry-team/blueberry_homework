from repositories.auth_repository import AuthRepository


class AuthUseCase:
    @staticmethod
    def sign_up(user_req):
        return AuthRepository.sign_up(user_req)

    @staticmethod
    def log_in(login_req):
        return AuthRepository.log_in(login_req)

    @staticmethod
    def change_user(change_req):
        return AuthRepository.change_user(change_req)

    @staticmethod
    def get_user(user_id):
        return AuthRepository.get_user(user_id)
