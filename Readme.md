# GraphQL Demo - Clean Architecture

A Go project demonstrating Clean Architecture with GraphQL, PostgreSQL, and gqlgen.

## Project Structure

```
graphql-demo/
├── internal/
│   ├── payment/                    # New subdomain
│   │   ├── delivery/
│   │   │   └── graphql/
│   │   ├── repository/
│   │   ├── usecase/
│   │   └── payment.go
│   └── user/                       # Existing subdomain
│       └── ...                     # Same structure as payment
├── models/
├── dto/
├── go.mod
├── gqlgen.yml
└── main.go
```

## Adding a New Subdomain (e.g., Payment)

### 1. Create Subdomain Structure

```bash
mkdir -p internal/payment/{delivery/graphql,repository,usecase}
touch internal/payment/payment.go
```

### 2. Define Interfaces (`payment.go`)

```go
package payment

import (
	"context"
	"graphql-demo/dto"
	"graphql-demo/models"
)

type Repository interface {
	Create(ctx context.Context, payment *models.Payment) error
	FindByID(ctx context.Context, id uint) (*models.Payment, error)
}

type Usecase interface {
	CreatePayment(ctx context.Context, input dto.PaymentInput) (*dto.PaymentResponse, error)
	GetPayment(ctx context.Context, id uint) (*dto.PaymentResponse, error)
}
```

### 3. Add GraphQL Schema

Create `internal/payment/delivery/graphql/schema.graphqls`:

```graphql
type Payment {
  id: ID!
  amount: Float!
  status: String!
  createdAt: String!
}

input PaymentInput {
  amount: Float!
  userId: ID!
}

extend type Mutation {
  createPayment(input: PaymentInput!): Payment!
}

extend type Query {
  payment(id: ID!): Payment!
}
```

### 4. Update gqlgen.yml

```yaml
schema:
  - internal/user/delivery/graphql/schema.graphqls
  - internal/payment/delivery/graphql/schema.graphqls  # Added
# ... rest of config remains same
```

### 5. Generate GraphQL Code

```bash
# Clean existing generated files
find internal -name "*.go" -not -path "*delivery/graphql/model*" -delete

# Regenerate
go run github.com/99designs/gqlgen generate
```

## Running the Project

### Prerequisites
- Go 1.16+
- PostgreSQL

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/graphql-demo.git
   cd graphql-demo
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up database:
   ```bash
   createdb graphql_demo
   ```

4. Configure environment variables:
   ```bash
   cp .env.example .env
   ```

### Running

```bash
go run main.go
```

Access GraphQL playground at `http://localhost:8080`

```gitignore
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/

# Test binary
*.test

# Dependency directories
vendor/
node_modules/

# Environment files
.env
.env.local

# IDE specific
.idea/
.vscode/
*.swp
*.swo

# gqlgen generated files
internal/*/delivery/graphql/generated.go
internal/*/delivery/graphql/schema.resolvers.go

# Database files
*.db
*.sqlite

# Logs
*.log

# OS generated
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db
```

## Key Commands

| Command | Description |
|---------|-------------|
| `go run github.com/99designs/gqlgen generate` | Generate GraphQL code |
| `go run main.go` | Start development server |
| `go test ./...` | Run all tests |

## Architecture Flow

1. **Request** → GraphQL Handler (Delivery)
2. **Handler** → Calls Usecase
3. **Usecase** → Business Logic → Calls Repository
4. **Repository** → Database Operations
