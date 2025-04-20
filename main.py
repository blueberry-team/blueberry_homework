import os
from dotenv import load_dotenv
from fastapi import FastAPI
from routers.custom_router import custom_router
# 환경변수 로드
load_dotenv()

app = FastAPI()

# 라우터 등록
app.include_router(custom_router)

PORT = int(os.getenv("PORT", 8000))  # 기본값 8000

# 포트 번호를 .env로부터 가져오는 코드
if __name__ == "__main__":
    import uvicorn

    print(f"서버를 시작합니다. {PORT}번 포트에서 실행중입니다.")
    print(f"문서 : http://localhost:{PORT}/docs")
    uvicorn.run(app, host="0.0.0.0", port=PORT)
