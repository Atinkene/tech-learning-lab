from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_methods=["*"],
    allow_headers=["*"],
)

Students = [
    {
        "id": 0,
        "firstname": "John",
        "lastname": "Doe",
        "age": 25,
        "email": "john.doe@example.com",
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "state": "CA",
            "zip": "12345"
        },
        "courses": [
            {
                "name": "Mathematics",
                "grade": "A"
            },
            {
                "name": "English",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 1,
        "firstname": "Jane",
        "lastname": "Smith",
        "age": 30,
        "email": "jane.smith@example.com",
        "address": {
            "street": "456 Oak Ave",
            "city": "Somewhere",
            "state": "NY",
            "zip": "67890"
        },
        "courses": [
            {
                "name": "Physics",
                "grade": "A"
            },
            {
                "name": "Chemistry",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
     {
        "id": 1,
        "firstname": "Jane",
        "lastname": "Smith",
        "age": 30,
        "email": "jane.smith@example.com",
        "address": {
            "street": "456 Oak Ave",
            "city": "Somewhere",
            "state": "NY",
            "zip": "67890"
        },
        "courses": [
            {
                "name": "Physics",
                "grade": "A"
            },
            {
                "name": "Chemistry",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
     {
        "id": 1,
        "firstname": "Jane",
        "lastname": "Smith",
        "age": 30,
        "email": "jane.smith@example.com",
        "address": {
            "street": "456 Oak Ave",
            "city": "Somewhere",
            "state": "NY",
            "zip": "67890"
        },
        "courses": [
            {
                "name": "Physics",
                "grade": "A"
            },
            {
                "name": "Chemistry",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
     {
        "id": 1,
        "firstname": "Jane",
        "lastname": "Smith",
        "age": 30,
        "email": "jane.smith@example.com",
        "address": {
            "street": "456 Oak Ave",
            "city": "Somewhere",
            "state": "NY",
            "zip": "67890"
        },
        "courses": [
            {
                "name": "Physics",
                "grade": "A"
            },
            {
                "name": "Chemistry",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    },
    {
        "id": 2,
        "firstname": "Alice",
        "lastname": "Johnson",
        "age": 22,
        "email": "alice.johnson@example.com",
        "address": {
            "street": "789 Pine Rd",
            "city": "Elsewhere",
            "state": "TX",
            "zip": "54321"
        },
        "courses": [
            {
                "name": "Biology",
                "grade": "A"
            },
            {
                "name": "Geography",
                "grade": "B"
            }
        ]       
    }
]

@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.get("/Students")
def read_students():
    return {"students": Students}  # Utilisez "students" en minuscules pour correspondre à votre React

@app.get("/Students/{item_id}")
def read_item(item_id: int, q: str | None = None):
    # Correction : Students est une liste, pas un dict
    if 0 <= item_id < len(Students):
        return {"item": Students[item_id], "q": q}
    raise HTTPException(status_code=404, detail="Student not found")

@app.get("/docs")
def read_docs():
    return {"docs": "This is the documentation page"}