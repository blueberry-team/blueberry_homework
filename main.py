from fastapi import FastAPI, HTTPException
from name_model import NameModel
from name_repository import add_name
from tmp_database import tmp_db

app = FastAPI()


@app.post("/createName")
def create_name(input_name: NameModel):
    # 입력값 검증
    if not input_name.name or input_name.name.strip() == "":
        raise HTTPException(status_code=400, detail="이름은 비어있을 수 없습니다")

    if len(input_name.name) > 50:
        raise HTTPException(status_code=400, detail="이름이 너무 깁니다 (최대 50자)")

    try:
        add_name(input_name.name)
        return {"message": "이름이 추가되었습니다", "name": input_name.name}
    except Exception as e:
        raise HTTPException(
            status_code=500, detail="서버 오류가 발생했습니다 : " + str(e)
        )


@app.get("/getName")
def get_names():

    names = get_names()

    try:
        if not tmp_db:
            return {"message": "등록된 이름이 없습니다", "names": names}
        return {"message": "이름 목록을 가져왔습니다", "names": names}
    except Exception as e:
        raise HTTPException(
            status_code=500, detail="서버 오류가 발생했습니다 : " + str(e)
        )
