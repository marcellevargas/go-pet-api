# Pet API

Pet API build with go

## Features

- CRUD operations for pets
- In-memory storage (can be extended to use a database)
- RESTful endpoints
- Built with Gin Framework

## Requirements

- Go 1.21 or higher
- Git

## Local Development Setup

1. Clone the repository:
```bash
https://github.com/marcellevargas/go-pet-api
cd go-pet-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Create a Pet
```bash
curl -X POST http://localhost:8080/pets \
-H "Content-Type: application/json" \
-d '{
    "name": "Rex",
    "age": 2,
    "tutor": "John Smith",
    "photo": "https://example.com/rex-photo.jpg"
}'
```

### List all Pets
```bash
curl http://localhost:8080/pets
```

### Get Pet by ID
```bash
curl http://localhost:8080/pets/{id}
```

### Update a Pet
```bash
curl -X PUT http://localhost:8080/pets/{id} \
-H "Content-Type: application/json" \
-d '{
    "name": "Rex",
    "age": 3,
    "tutor": "John Smith",
    "photo": "https://example.com/rex-new-photo.jpg"
}'
```

### Delete a Pet
```bash
curl -X DELETE http://localhost:8080/pets/{id}
```

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── domain/             # Domain models
│   ├── handlers/           # HTTP handlers
│   ├── repository/         # Persistence layer
│   └── service/            # Business logic
├── go.mod                  # Go dependencies
├── go.sum                  # Dependencies checksums
└── README.md              # This file
```

## Data Model

```json
{
    "id": "string",
    "name": "string",
    "age": int,
    "tutor": "string",
    "photo": "string"
}
``` 