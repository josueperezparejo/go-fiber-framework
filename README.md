# Go Project Management API

A RESTful API built with Go, Fiber, and PostgreSQL for managing projects and tasks.

## 🚀 Features

- Project management (CRUD operations)
- Task management with project association
- Swagger documentation
- PostgreSQL database integration
- Input validation
- Error handling
- CORS enabled
- Security middleware (Helmet)

## 📋 Prerequisites

- Go 1.24 or higher
- PostgreSQL
- Swag CLI (for Swagger documentation)

## 🛠️ Installation

1. Clone the repository:

```bash
git clone https://github.com/josueperezparejo/my-go.git
cd my-go
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database
DB_PORT=5432
DB_SSLMODE=disable
```

4. Generate Swagger documentation:

```bash
swag init
```

5. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:3000`

## 📚 API Documentation

Once the server is running, you can access the Swagger documentation at:

```
http://localhost:3000/swagger/
```

## 🏗️ Project Structure

```
.
├── config/         # Configuration files
├── controllers/    # HTTP request handlers
├── docs/          # Swagger documentation
├── dtos/          # Data Transfer Objects
├── models/        # Database models
├── routes/        # Route definitions
├── services/      # Business logic
└── utils/         # Utility functions
```

## 🔧 API Endpoints

### Projects

- `GET /api/projects` - Get all projects
- `GET /api/projects/:id` - Get project by ID
- `POST /api/projects` - Create new project
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Delete project

### Tasks

- `GET /api/tasks` - Get all tasks
- `GET /api/tasks/:id` - Get task by ID
- `POST /api/tasks` - Create new task
- `PUT /api/tasks/:id` - Update task
- `DELETE /api/tasks/:id` - Delete task
- `GET /api/projects/:projectId/tasks` - Get tasks by project

## 🧪 Testing

Run the tests with:

```bash
go test ./...
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📫 Contact

Josue Perez Parejo - [@josueperezparejo](https://github.com/josueperezparejo)

Project Link: [https://github.com/josueperezparejo/my-go](https://github.com/josueperezparejo/my-go)
