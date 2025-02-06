import app.repositories.main as repo

def get_user(id: int):
    res = repo.get_user(id)
    return res
