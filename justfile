# Загружать переменные из .env файла
set dotenv-load := true
set shell := ['fish', '-c']

run:
  uvicorn app.main:app --reload

venv:
  source venv/bin/activate.fish

install:
  pip install -r requirements.txt
