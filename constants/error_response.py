# 에러 메시지 상수 정의
ERROR_RESPONSES = {
    "DUPLICATE_NAME": {"message": "error", "error": "이미 존재하는 이름입니다"},
    "USER_NOT_FOUND": {"message": "error", "error": "존재하지 않는 사용자입니다"},
    "INVALID_INDEX": {"message": "error", "error": "유효하지 않은 인덱스입니다"},
    "NAME_VALIDATION_ERROR": {
        "message": "error",
        "error": "이름은 1~50자 사이여야 합니다",
    },
    "USER_ALREADY_HAS_COMPANY": {"message": "error", "error": "이미 회사가 존재합니다"},
    "SERVER_ERROR": {
        "message": "error",
        "error": "예상치 못한 서버 오류가 발생했습니다",
    },
    "COMPANY_NOT_FOUND": {"message": "error", "error": "존재하지 않는 회사입니다"},
}
