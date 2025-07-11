# Go Docker App

A simple RESTful Go web application with Docker support.

This app provides three endpoints:

- `GET /` – Welcome message
- `GET /health` – Health check
- `GET /greet?name=YourName` – Personalized greeting

---

## 📦 Features

- Minimal Go REST API
- Multi-stage Docker build
- Lightweight and fast
- Compatible with Ubuntu-based runtime

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/rishugkp688/Go-Docker_Project.git
cd Go-Docker_Project
```

```bash
docker build -t go-docker-app .
```

```bash
docker run -p 8080:8080 go-docker-app
```
