from fastapi import FastAPI

app = FastAPI()

# 일단 이름을 저장할 공간을 생성
names = []

# 이름을 추가하는 함수
@app.post("/createName")
def create_name(name: str):
    names.append(name)
    return {"message": "이름이 추가되었습니다", "name": name}

# 이름을 가져오는 함수
@app.get("/getName")
def get_names():
    return names