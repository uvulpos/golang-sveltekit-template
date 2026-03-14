# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Architecture Overview

This is a full-stack application template that combines Golang backend with SvelteKit frontend. The project supports two deployment modes:
1. **Single binary mode**: Backend serves the frontend as embedded static files
2. **Microservices mode**: Backend and frontend run as separate services

### Tech Stack
- **Backend**: Go 1.25+ with Fiber web framework, SQLX for database, PostgreSQL
- **Frontend**: SvelteKit with TypeScript, Vite, SvelteUI components
- **Infrastructure**: Docker Compose for development, Traefik reverse proxy, optional Authentik for auth

### Project Structure
- `/services/backend/` - Go backend service
  - `/src/cmd/` - CLI commands (run, migrate)
  - `/src/resources/` - API endpoints and handlers
  - `/src/helper/` - Utility packages
  - `/src/web-app/` - Embedded frontend handling
  - `/src/migrator/` - Database migrations
  - `/swagger-docs/` - Auto-generated API documentation
- `/services/frontend/` - SvelteKit frontend
  - `/src/routes/` - SvelteKit pages and routing
  - `/src/lib/` - Shared components and utilities
- `/devops/` - Build scripts and Docker configurations

## Common Commands

### Development
```bash
# Start all services in Docker (backend, frontend, database, reverse proxy)
make dev

# Run backend locally (requires database)
cd services/backend
go run src/main.go run

# Run frontend locally
cd services/frontend
npm install
npm run dev
```

### Testing
```bash
# Run backend tests in Docker
make test-be

# Run backend integration tests
make test-be-integrationtest

# Run frontend tests
make test-fe

# Run a single Go test
cd services/backend
go test ./src/helper/customerrors -run TestHttpDatabaseErrorHandling
```

### Building
```bash
# Build single binary (frontend embedded in backend)
make build-full

# Build Docker images
make build-dockerfile-binary    # Single container with everything
make build-dockerfile-backend   # Backend microservice only
make build-dockerfile-frontend  # Frontend microservice only

# Local release build (all platforms)
make local-release
```

### Database Operations
```bash
# Run database migrations
cd services/backend
go run src/main.go migrate-db up
```

### Code Quality
```bash
# Quick quality checks before committing
make check-all         # Run all tests, linting, and formatting checks
make check-quick       # Run only linting and formatting (faster)
make fix-all          # Auto-fix linting and formatting issues

# Individual checks
make lint-be          # Run Go linting (golangci-lint)
make lint-fe          # Run frontend ESLint
make format-check-fe  # Check frontend formatting (Prettier)
make format-fe        # Fix frontend formatting
make lint-fix-fe      # Fix frontend linting issues

# Setup pre-commit hooks (one-time setup)
make setup-hooks      # Enables automatic checks before commits

# Check frontend TypeScript
cd services/frontend
npm run check

# Generate Swagger docs (run from backend directory)
cd services/backend
swag init -g ./src/main.go -o ./swagger-docs --parseDependency --parseInternal

# License checks
make license-check-be  # Go dependencies
make license-check-fe  # NPM dependencies
```

## Backend Architecture

The backend uses a command-based structure with Cobra CLI:
- **Main entry**: `services/backend/src/main.go` → `cmd.Execute()`
- **Commands**: `run` (start server), `migrate-db` (database migrations)
- **API Structure**: Resources organized by domain in `/src/resources/`
- **Error Handling**: Custom error types in `/src/helper/customerrors/` with HTTP mapping
- **Database**: PostgreSQL with SQLX, migrations in `/src/migrator/database-migrations/`
- **Authentication**: JWT-based with optional OAuth2 integration

## Frontend Architecture

SvelteKit application with:
- **Routing**: File-based routing in `/src/routes/`
- **State Management**: SvelteKit stores and context
- **API Client**: Ky HTTP client for backend communication
- **UI Components**: SvelteUI component library
- **Internationalization**: svelte-i18n for multi-language support
- **Build Modes**: Static adapter for embedding, Node adapter for microservice
- **Testing**: Vitest with Testing Library for component tests
- **Code Quality**: ESLint for linting, Prettier for formatting

## API Documentation

Swagger UI is available at `/swagger` when running the backend. The API follows RESTful conventions with JSON payloads.