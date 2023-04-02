### GET ALL USERS
GET http://localhost:8000/users

### GET USER
GET http://localhost:8000/users/5

### INSERT USER
POST http://localhost:8000/users
Content-Type: application/json

{
    "name": "Aulia",
    "email": "aul@gmail.com",
    "password": "10020"
}

### DELETE USER
DELETE  http://localhost:8000/users/2

### PUT USER
PUT http://localhost:8000/users/2
Content-Type: application/json

{
    "name": "Aul",
    "email": "aul@gmail.com",
    "password": "7777"
}

### GET ALL BOOKS
GET http://localhost:8000/books

### GET BOOK
GET http://localhost:8000/books/1

### INSERT BOOK
POST http://localhost:8000/books
Content-Type: application/json

{
    "title": "Laut Bercerita",
    "author": "Leila",
    "description": "Novels"
}

### DELETE BOOK
DELETE  http://localhost:8000/books/1

### PUT BOOK
PUT http://localhost:8000/books/2
Content-Type: application/json

{
    "title": "Ikan di laut",
    "author": "Ziggy",
    "description": "Novels"
}