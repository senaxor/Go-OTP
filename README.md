# 🛡️ Go-OTP

Go-OTP is a high-performance, containerized One-Time Password (OTP) generator and authentication service written in Go. It provides secure OTP-based verification suitable for modern web and API-based applications.

## 🚀 Features

- ✅ OTP generation and verification
- 🔒 Time-based (TOTP) or random OTP support
- ⚙️ RESTful API endpoints
- 🧱 MongoDB for persistent storage
- 🐳 Docker and Docker Compose ready
- 🧪 Ready for integration with web or mobile apps

---

## 📦 Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Go](https://golang.org/) for development and debugging

---

## 🛠️ Getting Started

### 1. Build & Start Containers:
```bash
swag init   --dir ./cmd/server,./internal/database,./internal/handlers,./internal/models
```
```bash
docker-compose -f docker-compose.yml build --no-cache
```
```bash
docker-compose up
```

## Swagger
visit
http://localhost:8080/swagger/index.html