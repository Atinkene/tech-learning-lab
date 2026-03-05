from fastapi import APIRouter, status
from app.schemas.user import UserCreate, UserUpdate, UserResponse
from app.services import user_service

router = APIRouter(
    prefix="/users",
    tags=["Users"],  # Groupe dans Swagger
)

@router.post("/", response_model=UserResponse, status_code=status.HTTP_201_CREATED)
def create_user(user: UserCreate):
    return user_service.create_user(user)

@router.get("/", response_model=list[UserResponse])
def list_users():
    return user_service.get_all_users()

@router.get("/{user_id}", response_model=UserResponse)
def get_user(user_id: int):
    return user_service.get_user_by_id(user_id)

@router.patch("/{user_id}", response_model=UserResponse)
def update_user(user_id: int, data: UserUpdate):
    return user_service.update_user(user_id, data)

@router.delete("/{user_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_user(user_id: int):
    user_service.delete_user(user_id)