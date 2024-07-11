# Go Task Server

This server is designed to manage CRUD operations with an integrated authorization mechanism, developed as part of an interview assignment for a company.

## Technologies Used

- **Go** v1.21
- **gin** v1.10.0
- **gin-swagger** v1.16.3
- **logrus** v1.9.3
- **testify** v1.9.0
- **go-cache** v2.1.0+incompatible

## How to Run

To build and run the Docker image, use the following commands:

```sh
docker build -t gotaskserver .
docker run -p 8080:8080 gotaskserver
```

## How to Use

### Swagger Documentation

Access the API documentation at: [Swagger UI](http://localhost:8080/swagger/index.html)

### API Endpoints

#### Authorization Mechanism

1. **Create a User**
   - User passwords are hashed before being saved to storage.
   - `POST /user`
```cURL
curl --location 'localhost:8080/api/v1/user' \
--header 'Content-Type: application/json' \
--data '{
    "name": "andy",
    "password": "123"
}'
```

2. **Login to Get API Key**
   - The API key expires after 1 minute.
   - `POST /user/login`
```cURL
curl --location 'localhost:8080/api/v1/user/login' \
--header 'Content-Type: application/json' \
--data '{
    "name": "andy",
    "password": "123"
}'
```
```cURL
{
    "result": "d501d827-5e6d-4c84-ae3c-1248608ac084"
}
```

#### CRUD Tasks

All task-related API requests must include the header `H-API-KEY`. Requests without this header will be rejected with a 401 status code.

1. **List Tasks**
   - `GET /tasks`
```cURL
curl --location 'localhost:8080/api/v1/tasks' \
--header 'H-API-KEY: afd66253-a36e-4034-988d-83b9e9d81f19'
```

2. **Create Task**
   - `POST /task`
```cURL
curl --location 'localhost:8080/api/v1/task' \
--header 'H-API-KEY: afd66253-a36e-4034-988d-83b9e9d81f19' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AAA"
}'
```

3. **Put Task**
   - `PUT /task/:id`
```cURL
curl --location --request PUT 'localhost:8080/api/v1/task/1' \
--header 'H-API-KEY: afd66253-a36e-4034-988d-83b9e9d81f19' \
--header 'Content-Type: application/json' \
--data '{
        "id": 1,
        "name": "BBB",
        "status": 1
    }'
```

4. **Delete Task**
   - `DELETE /task/:id`
```cURL
curl --location --request DELETE 'localhost:8080/api/v1/task/3' \
--header 'H-API-KEY: afd66253-a36e-4034-988d-83b9e9d81f19'
```

## Introduction

- **api**: Handles HTTP requests.
- **module**: Defines interfaces and business logic for authentication, user management, and tasks.
  - **Auth module**: Handles API key creation and validation for authentication.
  - **User module**: Manages user creation and login.
  - **Task module**: Manages CRUD operations for tasks.
- **dataaccess**: Implements storage for authentication, user management, and tasks.
  - The assignment requires storing data in memory.
  - Memory cache is chosen for in-memory storage.
  - Data stored in memory cannot be shared with other pods.
  - Interfaces are used to facilitate easy changes to different storage implementations.
- **log**: Utilizes Logrus for logging.
  - Each log entry includes a TraceID and UserID to aid in tracing and debugging.
- **error**: Defines error codes to simplify debugging and integration.
