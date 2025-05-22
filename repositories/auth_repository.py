from db.db import ScyllaDB
from entities.user_entity import UserEntity
from passlib.hash import bcrypt
import uuid
from datetime import datetime


class AuthRepository:
    @staticmethod
    def sign_up(user_req):
        session = ScyllaDB.get_session()

        # 이메일 중복 체크
        query = "SELECT id FROM users WHERE email = %s ALLOW FILTERING"
        rows = session.execute(query, (user_req["email"],))
        if any(rows):
            raise Exception("이미 존재하는 이메일입니다.")
        # 비밀번호 해싱
        hashed_pw = bcrypt.hash(user_req["password"])
        user_id = str(uuid.uuid4())
        now = datetime.now()
        # DB 저장
        insert_query = """
        INSERT INTO users (id, username, email, password, address, role, created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s)
        """
        session.execute(
            insert_query,
            (
                uuid.UUID(user_id),
                user_req["name"],
                user_req["email"],
                hashed_pw,
                user_req.get("address", ""),
                user_req.get("role", "worker"),
                now,
                now,
            ),
        )
        return {"message": "회원가입 성공", "user_id": user_id}

    @staticmethod
    def log_in(login_req):
        session = ScyllaDB.get_session()
        query = "SELECT * FROM users WHERE email = %s ALLOW FILTERING"
        rows = session.execute(query, (login_req["email"],))
        user = None
        for row in rows:
            user = row
            break
        if not user:
            raise Exception("존재하지 않는 이메일입니다.")
        if not bcrypt.verify(login_req["password"], user["password"]):
            raise Exception("비밀번호가 일치하지 않습니다.")
        # 로그인 성공 시 최소 정보 반환
        return {
            "message": "로그인 성공",
            "user_id": str(user["id"]),
            "name": user["username"],
            "email": user["email"],
            "role": user["role"],
        }

    @staticmethod
    def change_user(change_req):
        session = ScyllaDB.get_session()
        user_id = change_req["user_id"]
        # 기존 유저 조회
        query = "SELECT * FROM users WHERE id = %s"
        rows = session.execute(query, (uuid.UUID(user_id),))
        user = None
        for row in rows:
            user = row
            break
        if not user:
            raise Exception("존재하지 않는 유저입니다.")
        # 업데이트할 필드만 추출
        update_fields = []
        update_values = []
        for field in ["name", "email", "address", "role", "password"]:
            if field in change_req:
                if field == "password":
                    update_fields.append("password = %s")
                    update_values.append(bcrypt.hash(change_req["password"]))
                elif field == "name":
                    update_fields.append("username = %s")
                    update_values.append(change_req["name"])
                else:
                    update_fields.append(f"{field} = %s")
                    update_values.append(change_req[field])
        update_fields.append("updated_at = %s")
        update_values.append(datetime.now())
        update_values.append(uuid.UUID(user_id))
        update_query = f"UPDATE users SET {', '.join(update_fields)} WHERE id = %s"
        session.execute(update_query, tuple(update_values))
        return {"message": "유저정보 수정 성공", "user_id": user_id}

    @staticmethod
    def get_user(user_id):
        session = ScyllaDB.get_session()
        query = "SELECT * FROM users WHERE id = %s"
        rows = session.execute(query, (uuid.UUID(user_id),))
        for row in rows:
            return {
                "user_id": str(row["id"]),
                "name": row["username"],
                "email": row["email"],
                "address": row.get("address", ""),
                "role": row.get("role", "worker"),
                "created_at": row["created_at"],
                "updated_at": row["updated_at"],
            }
        raise Exception("존재하지 않는 유저입니다.")

    @staticmethod
    def find_by_id(user_id):
        # 내부용: get_user와 동일하게 동작
        return AuthRepository.get_user(user_id)
