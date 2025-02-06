from fastapi import APIRouter
import app.usecases.main as usecases

router = APIRouter()

@router.get('/user/{id}')
async def get_user(id: int):
    res = usecases.get_user(id)
    return res
