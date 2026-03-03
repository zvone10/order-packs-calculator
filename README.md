# Order Packs Calculator

A full-stack application for calculating optimal packing orders. The system determines the best combination of pack sizes to fulfill customer orders efficiently.

## Overview

This project consists of:
- **Backend API**: Go-based REST API using the Gin framework
- **Frontend UI**: Modern React application built with Vite and TypeScript
- **Containerization**: Docker & Docker Compose for easy local development and deployment

## Features

- 📦 Knapsack dynamic programming pack size calculation algorithms
- 🎨 Responsive web interface with React
- 🔄 RESTful API with CORS support
- 🐳 Docker & Docker Compose support for easy deployment
- 📊 Real-time calculation results
- 💾 State management with Zustand

## Quick Start with Docker Compose

The easiest way to run the entire application locally:

```bash
docker-compose up
```

The application will be available at:
- **Frontend**: http://localhost:3000
- **API**: http://localhost:8080
- **API Health**: http://localhost:8080/health

To stop the application:

```bash
docker-compose down
```

## Local Development Setup

### Prerequisites

- **Node.js** 18+ and npm (for frontend development)
- **Go** 1.25+ (for backend development)
- **Docker & Docker Compose** (optional, for containerized development)

### Backend Setup

```bash
cd api

# Install dependencies
go mod download

# Run the application
go run ./cmd/main.go
```

The API will start on `http://localhost:8080`

**Environment Variables:**
- `PORT`: Server port (default: 8080)

### Frontend Setup

```bash
cd web-ui

# Install dependencies
npm install

# Development server
npm run dev
```

The UI will typically be available at `http://localhost:5000`

To build for production:

```bash
npm run build
npm start  # Serve the built application
```

## Project Structure

```
.
├── api/                          # Go backend
│   ├── cmd/
│   │   └── main.go              # Application entry point
│   ├── internal/
│   │   ├── calculator/          # Packing algorithms
│   │   ├── handler/             # HTTP request handlers
│   │   ├── model/               # Data models
│   │   ├── repository/          # Data access layer
│   │   └── service/             # Business logic
│   ├── go.mod & go.sum          # Go dependencies
│   ├── Dockerfile               # Backend container
│   └── Procfile                 # Heroku deployment config
│
├── web-ui/                      # React frontend
│   ├── src/
│   │   ├── components/          # React components
│   │   ├── store/               # Zustand state management
│   │   ├── utils/               # Utility functions
│   │   └── main.tsx             # Entry point
│   ├── package.json             # Node dependencies
│   ├── vite.config.ts           # Vite configuration
│   ├── tsconfig.json            # TypeScript configuration
│   ├── Dockerfile               # Frontend container
│   └── Procfile                 # Heroku deployment config
│
└── docker-compose.yml           # Multi-container setup
```

## API Endpoints

### Health Check
```
GET /health
```
Returns the health status of the API.

### Pack Calculation
```
POST /api/v1/pack-calculation
Content-Type: application/json

{
  "order_quantity": 100,
  "pack_sizes": [5, 10, 20, 50]
}
```

## Docker & Docker Compose

### Running with Docker Compose

```bash
# Start all services
docker-compose up

# Start in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

### Building Individual Containers

**Backend:**
```bash
cd api
docker build -t order-packs-calculator-api:latest .
docker run -p 8080:8080 order-packs-calculator-api:latest
```

**Frontend:**
```bash
cd web-ui
docker build -t order-packs-calculator-ui:latest .
docker run -p 3000:5000 -e VITE_API_URL=http://localhost:8080 order-packs-calculator-ui:latest
```

## Environment Configuration

### Frontend Environment Variables
- `VITE_API_URL`: Backend API URL (default: `http://localhost:8080`)

### Backend Environment Variables
- `PORT`: Server port (default: `8080`)

## Testing

Run unit tests in api folder with 
```
go test ./...
```
Run benchmark tests
```
go test -bench=. -benchmem ./internal/calculator
```

### Backend Tests
```bash
cd api
go test ./...
```

### Frontend Linting
```bash
cd web-ui
npm run lint
```

## Deployment

Both services are configured for Heroku deployment:
- Backend API: See api/Procfile
- Frontend UI: See web-ui/Procfile

## Technology Stack

### Backend
- **Go 1.25.5**: Programming language
- **Gin**: Web framework
- **CORS**: Cross-origin resource sharing

### Frontend
- **React 18**: UI framework
- **Vite**: Build tool and dev server
- **TypeScript**: Type-safe JavaScript
- **Zustand**: State management
- **Serve**: Static file server

## Troubleshooting

### Docker Compose issues
- If ports are already in use, modify the ports in `docker-compose.yml`
- Ensure Docker daemon is running
- Use `docker-compose logs` to check service logs

### Frontend can't connect to API
- Verify the API is running and accessible
- Check that `VITE_API_URL` environment variable is set correctly
- Ensure CORS is properly configured in the backend

### Port conflicts
- View all running containers: `docker ps`
- Stop specific service: `docker stop <container_id>`

## Contributing

When contributing to this project:
1. Test changes locally with Docker Compose
2. Ensure both frontend and backend work together
3. Update documentation as needed

## License

This project is available for use.


