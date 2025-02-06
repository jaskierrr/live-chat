from fastapi import FastAPI, APIRouter

app = FastAPI()

router = APIRouter()

@router.get('/')
async def read_root():
    return {"Hello": "World"}

@router.get('/lol')
async def lol():
    return {"LoL": 'что еще нужно'}

app.include_router(router)
