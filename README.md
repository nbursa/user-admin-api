# User Admin API

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-blue)
![MongoDB](https://img.shields.io/badge/MongoDB-NoSQL-green?logo=mongodb)
![Docker](https://img.shields.io/badge/Dockerized-BE%2FMongo-blue?logo=docker)
![Logrus](https://img.shields.io/badge/Logrus-Logging-orange)
![Validator](https://img.shields.io/badge/Validator-go--playground-lightgrey)

This is the backend service for managing users. Built with Go and Gin, it exposes a JSON REST API to perform CRUD operations with validation, pagination, and logging. The system is designed to be modular, testable, and database-agnostic.

## Tech Stack

- **Language:** Go 1.21+
- **Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** MongoDB (swappagle, Go driver)
- **ORM/Driver:** [mongo-driver](https://github.com/mongodb/mongo-go-driver)
- **Logging:** [Logrus](https://github.com/sirupsen/logrus)
- **Validation:** [go-playground/validator](https://github.com/go-playground/validator)
- **Containerization:** Docker & Docker Compose

## Features

- `GET /users` – List users with pagination and optional search
- `POST /users` – Create new user (email must be unique, age > 18)
- `GET /users/:id` – Get single user by ID
- `PUT /users/:id` – Update user by ID (validates email uniqueness)
- `DELETE /users/:id` – Delete user by ID
- Modular architecture for easy database replacement
- Unit tests for core service logic
- Request logging via Logrus
- Input validation using `validator`

## Getting Started

To start the API with Docker:

```bash
docker compose up --build
```

App will be available at:

```bash
http://localhost:8080
```

## Running Tests

To run unit tests for service layer logic:

```bash
go test ./...
```

## Frontend

See [user-admin-client](https://github.com/nbursa/user-admin-client) for the corresponding Vue 3 frontend application.
