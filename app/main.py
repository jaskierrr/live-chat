from fastapi import FastAPI
from app.handlers.main import router

app = FastAPI()

app.include_router(router)
