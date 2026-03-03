# Go User REST API

A lightweight RESTful API built with Go's standard `net/http` package. Demonstrates basic CRUD operations for managing a users collection.

## 🚀 Features
- **GET /users** - Retrieve all users
- **POST /users** - Add a new user (JSON body: `{"name": "John"}`)
- **DELETE /users** - Delete the last user
- In-memory storage
- JSON request/response handling

## 🛠️ Tech Stack
- Go (standard library only)
- `net/http` for routing
- `encoding/json` for serialization

## 📦 Quick Start

1. **Clone the repo**
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-user-rest-api.git
   cd go-user-rest-api
