from fastapi import FastAPI
from name_model import NameModel
from name_repository import add_name
from tmp_database import tmp_db

app = FastAPI()


@app.post("/createName")
def create_name(input_name: NameModel):
    add_name(input_name.name)
    return {"message": "이름이 추가되었습니다", "name": input_name.name}


@app.get("/getName")
def get_names():
    if not tmp_db:
        return {"message": "등록된 이름이 없습니다", "names": tmp_db}
    return {"message": "이름 목록을 가져왔습니다", "names": tmp_db}
