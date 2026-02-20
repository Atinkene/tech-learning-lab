from fastapi import FastAPI
from typing import Optional
from pydantic import BaseModel, EmailStr


app = FastAPI(
    title="My First FastAPI App",
    description="This is a simple FastAPI application to demonstrate basic functionality.",
    version="1.0.0" 
)


# Create an Api that allows us to apply crud operations on items and users. We will use path parameters, query parameters, and request bodies.  


class UserCreate(BaseModel):
    name: str
    email: str
    role: Optional[str] = "user"


class UserResponse(BaseModel):
    id: int
    name: str
    email: str
    role: str

users = [
    {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "role": "admin"
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane.smith@example.com",
        "role": "user"
    }
]


@app.get("/")
def root():
    return {"message": "Bienvenue dans FastAPI 🚀"}


@app.get('/users')
def get_users():
    return users


@app.get('/users/{user_id}')
def get_user(user_id: int):
    for u in users:
        if u["id"] == user_id:
            return u
    return {"message": "User not found"}, 404


@app.post('/users')
def create_user(user_create: UserCreate):
    new_user = {
        "id": len(users) + 1,
        "name": user_create.name,
        "email": user_create.email,
        "role": user_create.role
    }
    if  user_create.email in [user["email"] for user in users]:
        return {"message": "Email already exists"}, 400
    users.append(new_user)
    return users, 201


@app.put('/users/{user_id}')
def update_user(user_id: int, user_update: UserCreate):
    for u in users:
        if u["id"] == user_id:
            u["name"] = user_update.name
            u["email"] = user_update.email
            u["role"] = user_update.role
            return u
    return {"message": "User not found"}, 404


@app.delete('/users/{user_id}')
def delete_user(user_id: int):
    for u in users:
        if u["id"] == user_id:
            users.remove(u)
            return {"message": "User deleted"}
    return {"message": "User not found"}, 404   
    