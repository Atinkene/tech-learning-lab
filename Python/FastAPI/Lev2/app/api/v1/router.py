from fastapi import APIRouter
from app.api.v1.routes import users

api_router = APIRouter()
api_router.include_router(users.router)

# Plus tard : api_router.include_router(tasks.router), etc.