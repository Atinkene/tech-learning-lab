from app.schemas.user import UserCreate, UserUpdate, UserResponse
from fastapi import HTTPException

_db: dict[int, dict] = {}
_counter: int = 0

def create_user(data: UserCreate) -> UserResponse:
    global _counter
    _counter += 1
    user = {"id": _counter, **data.model_dump()}
    _db[_counter] = user
    return UserResponse(**user)

def get_all_users() -> list[UserResponse]:
    return [UserResponse(**u) for u in _db.values()]

def get_user_by_id (user_id: int) -> UserResponse:
    if user_id not in _db:
        raise HTTPException(status_code=404, detail=f"Utilisateur {user_id} introuvable")
    return UserResponse(**_db[user_id])

def update_user (user_id: int, data: UserUpdate) -> UserResponse:
    if user_id not in _db:
        raise HTTPException(status_code=404, detail=f"Utilisateur {user_id} introuvable")

    updates = data.model_dump(exclude_unset=True)
    _db[user_id].update(updates)
    return UserResponse(**_db[user_id])

def delete_user(user_id: int) -> None:
    if user_id not in _db:
        raise HTTPException(status_code=404, detail=f"Utilisateur {user_id} introuvable")
    del _db[user_id]