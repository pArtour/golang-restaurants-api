# Restaurant Management System API

This project is a RESTful API for managing restaurant information. It is built using Golang, Fiber, Gorm, Postgres, Docker, Docker Compose and Air for live reload.
## Features
- CRUD operations for restaurants, menus and menu items
- Search and filter restaurants by name, location, and cuisine
- Authentication and Authorization using JWT
- Pagination and sorting for the list of restaurants
- Swagger documentation for all endpoints

## Getting Started
1. Clone the repository
```bash
git clone https://github.com/pArtour/restaurant-management-system.git
```
2. Build the Docker image
```bash
docker-compose build
```
3. Run the Docker container
```bash
docker-compose up
```
4. The API will be available at http://localhost:3000
5. You can access the Swagger documentation at http://localhost:3000/swagger/index.html

## Built With
- Golang - The programming language used
- Fiber - Web framework for Golang
- Gorm - ORM library for Golang
- Postgres - Database used
- Docker - Containerization platform
- Docker Compose - Tool for defining and running multi-container Docker applications
- Air - Live Reload tool