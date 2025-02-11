from fastapi import FastAPI
from app.handlers.main import router, router2, router3

app = FastAPI()
app.include_router(router)
app.include_router(router2)
app.include_router(router3)
