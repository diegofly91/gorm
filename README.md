# API user management with authentication and authorization

## Description
this is api example arquiteture hexagonal with golang, this api is for user management with 
with modules user


## Project structure
```bash
        myapp/                                                                                      
        ├── src/
        │   ├── main.go
        │   ├── config/
        │   │   └── database.go
        │   ├── common/
        │   │   └── handlers
        │   │       └── response.go
        │   ├──modules/
        |   │   ├── modules.go
        │   │   ├── user/
        │   │   │   ├── user_module.go
        |   │   │   ├── controller/
        |   │   │   |   └── user_controller.go
        |   │   │   ├── middlawares/
        |   │   │   │   └── body_decode_user.go
        |   │   │   │   └── exits_user.go
        │   │   │   ├── models/
        │   │   │   │   └── user.go
        |   │   │   ├── repository/
        |   │   │   │   └── user_repository.go
        |   │   │   ├── routes/
        |   │   │   │   └── user_router.go
        │   │   │   ├── service/
        |   │   │   │   └── user_service.go
        ├── .env
        |── .gitignore
        |── .fresh.yaml
        ├── go.mod
        ├── go.sum
        └── README.md
```

## Cómo Ejecutar

1. Instalar dependencias:
    ```sh
    go mod tidy
    ```

2. Ejecutar la aplicación:
    ```sh
    go run src/main.go
    ```
    fresh
    ```

La aplicación se ejecutará en `http://localhost:3000`.

## Endpoints
   ### User
    - **GET** `/api/user/`  // get all users
    - **GET** `/api/user/[0-9]+`  // get user by id
    - **POST** `/api/user/` // register
    - **PUT** `/api/user/` // update user
    - **DELETE** `/api/user/` // delete user

