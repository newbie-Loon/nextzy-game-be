# Nextzy-Game-BE

# Installation Guide

## Prerequisites

Before running the project, make sure the following software is installed:

- Node.js (v22+)
- Golang (v1.26+)
- PostgreSQL
- Docker (Optional)
- Docker Compose (Optional)

---

## Backend Setup

### Navigate to backend folder
cd nextzy-game-be

### Install dependencies

go mod tidy

### Create .env

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=nextzy_game
DB_SSLMODE=disable

### Run application
go run main.go

### Backend Server
http://localhost:8080

# Docker
Build : docker build -t nextzy-game-be .
Run : docker run -d \
--name nextzy-game-be \
--env-file .env \
-p 3050:3050 \
nextzy-game-be

Verify : docker logs -f nextzy-game-be


# Nextzy Game Backend

Backend service for the Nextzy Game application.

Built with:

- Golang
- Gin Framework
- GORM
- PostgreSQL
- Docker

---

# Features

## User Management

- Create Guest User
- Retrieve User Data
- Retrieve Home Page Data

## Point System

- Store play history
- Calculate accumulated points
- Validate point values

Allowed points:

- 300
- 500
- 1000
- 3000

## Reward System

Reward checkpoints:

| Reward | Required Score |
|---------|----------------|
| A | 5,000 |
| B | 7,500 |
| C | 10,000 |

Features:

- Claim reward
- Prevent duplicate claims
- Validate minimum score requirement

## Reset Progress

Tester can:

- Reset play history
- Reset reward history
- Start testing from scratch

---

# Architecture

```text
Client (Next.js)
        │
        ▼
   Gin Backend
        │
        ▼
    PostgreSQL
```
## Layers

handlers
    │
    ▼
services/business logic
    │
    ▼
database (gorm)

# Database Design

## Users
id
user_id
created_at
updated_at
deleted_at

## Point Histories
id
user_id
point
earned_date
created_at
updated_at
deleted_at

## Reward Histories

id
user_id
reward
created_at
updated_at
deleted_at

# API Endpoints

## GET /users/getUserData/:userId
Response
```json
{
  "user": {
    "userId": "uuid",
    "point": 8500
  },
  "reward": {
    "a": true,
    "b": true,
    "c": false
  },
  "playHistory": [],
  "rewardHistory": []
}
```

## POST /point-history/create
Request
```json
{
  "userId": "uuid",
  "point": 1000
}
```
Response
```json
{
  "earnedPoint": 1000,
  "totalPoint": 8500
}
```

## POST /reward-history/create
Request
```json
{
  "userId": "uuid",
  "reward": "A"
}
```
Response
```json
```

## POST /game/resetProgress
Request
```json
{
  "userId": "uuid",
  "reward": "A"
}
```
Response
```json
```

