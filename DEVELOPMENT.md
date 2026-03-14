# Development Guide

## Setting Up Development Environment

### Prerequisites
- Go 1.25+
- Node.js 18+
- Docker & Docker Compose
- golangci-lint (optional, for linting)

### Initial Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/uvulpos/golang-sveltekit-template.git
   cd golang-sveltekit-template
   ```

2. **Install dependencies**
   ```bash
   # Install frontend dependencies
   cd services/frontend && npm install
   cd ../..

   # Install Go dependencies
   cd services/backend && go mod download
   cd ../..
   ```

3. **Setup Git hooks (recommended)**
   ```bash
   make setup-hooks
   ```
   This will enable automatic checks before each commit.

## Quality Checks

### Available Make Commands

#### Quick Commands
- `make check-all` - Run all tests and checks (recommended before committing)
- `make check-quick` - Run only linting and formatting checks (faster)
- `make fix-all` - Automatically fix linting and formatting issues

#### Individual Checks
- `make lint-be` - Run Go linting
- `make lint-fe` - Run frontend ESLint
- `make format-check-fe` - Check frontend formatting (Prettier)
- `make test-be-local` - Run Go tests locally
- `make test-fe-local` - Run frontend tests locally

#### Fix Commands
- `make format-fe` - Fix frontend formatting issues
- `make lint-fix-fe` - Fix frontend linting issues (auto-fixable ones)

### Pre-commit Hooks

After running `make setup-hooks`, the following checks will run automatically before each commit:
- Backend linting (golangci-lint)
- Backend tests
- Frontend linting (ESLint)
- Frontend formatting (Prettier)
- Frontend tests (Vitest)

To skip pre-commit checks (not recommended):
```bash
SKIP_CHECKS=1 git commit -m "your message"
```

To disable hooks:
```bash
git config --unset core.hooksPath
```

To re-enable hooks:
```bash
make setup-hooks
```

## Development Workflow

### Recommended Workflow

1. **Before starting work:**
   ```bash
   git pull
   cd services/frontend && npm install
   cd ../backend && go mod download
   ```

2. **During development:**
   ```bash
   # Run the development environment
   make dev

   # Or run services individually
   make dev:frontend  # Frontend only
   make dev:backend   # Backend only
   ```

3. **Before committing:**
   ```bash
   # Run all checks
   make check-all

   # If there are issues, fix them
   make fix-all

   # Check again
   make check-all
   ```

4. **Commit your changes:**
   ```bash
   git add .
   git commit -m "feat: your feature description"
   ```
   The pre-commit hooks will run automatically.

## Testing

### Backend Testing
```bash
# Run all Go tests
make test-be-local

# Run with Docker (includes database)
make test-be

# Run specific test
cd services/backend
go test -v ./src/helper/customerrors -run TestHttpDatabaseErrorHandling
```

### Frontend Testing
```bash
# Run unit tests
make test-fe-local

# Run tests in watch mode
cd services/frontend
npm run test

# Run with coverage
cd services/frontend
npm run test:coverage
```

## Code Quality

### Linting Configuration
- **Backend**: `.golangci.yml` - Configures golangci-lint
- **Frontend**: `.eslintrc.cjs` - ESLint configuration
- **Formatting**: `.prettierrc` - Prettier configuration

### Common Issues and Fixes

#### Frontend ESLint errors
```bash
# See all errors
make lint-fe

# Auto-fix what's possible
make lint-fix-fe
```

#### Frontend formatting issues
```bash
# Check formatting
make format-check-fe

# Auto-fix formatting
make format-fe
```

#### Backend linting errors
```bash
# See all errors
make lint-be

# Most backend issues need manual fixes
# Common issues:
# - Unchecked errors: Add error checking
# - Unused variables: Remove or use them
# - Style issues: Follow Go conventions
```

## Tips

1. **Set up your editor:**
   - Install ESLint and Prettier extensions for frontend
   - Install Go extension with golangci-lint integration
   - Enable format-on-save

2. **Quick iteration:**
   - Use `make check-quick` for fast feedback during development
   - Run `make check-all` before pushing changes

3. **Debugging failed checks:**
   - Run individual make commands to see detailed output
   - Check the specific line numbers in error messages
   - Use `make fix-all` to fix common issues automatically