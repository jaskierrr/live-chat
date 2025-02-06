# Загружать переменные из .env файла
set dotenv-load := true

run:
  uvicorn app.main:app --reload

venv:
  source .venv/bin/activate

install:
  pip install -r requirements.txt
