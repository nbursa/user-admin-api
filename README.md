# user-admin-api

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-blue)
![MongoDB](https://img.shields.io/badge/MongoDB-NoSQL-green?logo=mongodb)
![Docker](https://img.shields.io/badge/Dockerized-BE%2FMongo-blue?logo=docker)
![Logrus](https://img.shields.io/badge/Logrus-Logging-orange)
![Validator](https://img.shields.io/badge/Validator-go--playground-lightgrey)

Golang backend for managing usersm.

## 🔧 Tech Stack

- **Language:** Go 1.21+
- **Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** MongoDB (via official Go driver)
- **ORM/Driver:** [mongo-driver](https://github.com/mongodb/mongo-go-driver)
- **Logging:** [Logrus](https://github.com/sirupsen/logrus)
- **Validation:** [go-playground/validator](https://github.com/go-playground/validator)
- **Containerization:** Docker & Docker Compose

## ✅ Features

- `POST /users` – Create a new user
  - Validates input (name, email, age)
  - Checks for unique email before saving
  - Returns created user in JSON format
  - Unit test for `CreateUser` service

## 🚀 Getting Started

```bash
docker compose up --build
```

App will be available at:

```bash
http://localhost:8080
```

## Frontend

See [user-admin-client](https://github.com/nbursa/user-admin-client) for the corresponding Vue 3 frontend application.
