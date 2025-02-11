from fastapi import APIRouter
import app.usecases.main as usecases

router = APIRouter()


@router.get("/user/{id}")
async def get_user(id: int):
    res = usecases.get_user(id)
    return res


router2 = APIRouter()


@router2.get("/task/{id}")
async def get_task(id: int):
    return "task №" + str(id)


router3 = APIRouter()


@router3.get("/profile/{id}")
async def get_profile(id: int):
    return "profile №" + str(id)
