# user-admin-api

Golang backend for managing users in the ABC company system.

## 🔧 Tech Stack

- **Language:** Go 1.24+
- **Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** MongoDB (via official Go driver)
- **ORM/Driver:** [mongo-driver](https://github.com/mongodb/mongo-go-driver)
- **Logging:** [Logrus](https://github.com/sirupsen/logrus)
- **Validation:** [go-playground/validator](https://github.com/go-playground/validator)
- **Containerization:** Docker & Docker Compose

## 🚀 Getting Started

```bash
docker compose up --build
```

App will be available at:

```bash
http://localhost:8080
```

## ✅ Features

- [x] `POST /users` – Create a new user
  - Validates input (name, email, age)
  - Checks for unique email before saving
  - Returns created user in JSON format
- [x] Unit test for `CreateUser` service
