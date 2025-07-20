Certainly! Here's a **professionally written and structured `README.md`** tailored for a production-grade Go project using GraphQL, PostgreSQL, `gqlgen`, and a `Makefile` for automation:

---

````markdown
# 📦 GraphQL Demo - Clean Architecture in Go

This project demonstrates a clean, modular architecture for building scalable GraphQL APIs using **Go**, **PostgreSQL**, and **gqlgen**. It emphasizes separation of concerns, testability, and maintainability, following Clean Architecture principles.

---

## 📁 Project Structure

```bash
graphql-demo/
├── internal/                         # Domain-specific packages
│   ├── user/                         # User subdomain
│   │   └── delivery/graphql/         # GraphQL schema and resolvers
│   └── payment/                      # Payment subdomain
│       ├── delivery/graphql/
│       ├── repository/
│       ├── usecase/
│       └── payment.go                # Domain interfaces
├── models/                           # Database models
├── dto/                              # Data transfer objects
├── gqlgen.yml                        # gqlgen configuration
├── go.mod / go.sum                   # Go modules
├── Makefile                          # Automation commands
└── main.go                           # Application entrypoint
````

---

## 🛠 Features

* ✅ Clean Architecture principles (Domain → Usecase → Delivery)
* ✅ GraphQL API using gqlgen
* ✅ PostgreSQL integration
* ✅ Automated GraphQL codegen with resolver preservation
* ✅ Modular and extensible structure
* ✅ `Makefile` for streamlined developer workflow

---

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/yourusername/graphql-demo.git
cd graphql-demo
cp .env.example .env
make setup
```

### 2. Run the Application

```bash
make start
```

Visit GraphQL Playground at [http://localhost:8080](http://localhost:8080)

---

## 🧱 Adding a New Subdomain (e.g., `payment`)

### 1. Create the Structure

```bash
mkdir -p internal/payment/{delivery/graphql,repository,usecase}
touch internal/payment/payment.go
```

### 2. Define Domain Interfaces

```go
// internal/payment/payment.go
package payment

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

```graphql
# internal/payment/delivery/graphql/schema.graphqls
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

### 4. Update `gqlgen.yml`

```yaml
schema:
  - internal/user/delivery/graphql/schema.graphqls
  - internal/payment/delivery/graphql/schema.graphqls
```

### 5. Generate Code

```bash
make generate
```

---

## 📘 Makefile Overview

The included `Makefile` provides commonly used commands to speed up development:

| Command         | Description                                   |
| --------------- | --------------------------------------------- |
| `make setup`    | Install dependencies and start PostgreSQL     |
| `make generate` | Regenerate GraphQL code (preserves resolvers) |
| `make migrate`  | Run database migrations                       |
| `make fmt`      | Format Go code                                |
| `make lint`     | Lint the codebase                             |
| `make start`    | Start the application                         |
| `make build`    | Build binary to `bin/` folder                 |
| `make test`     | Run unit tests                                |
| `make clean`    | Remove generated and backup files             |
| `make init`     | Full setup: DB + generate + migrate + fmt     |

---

## 🧪 Testing

Run all unit tests with verbose output:

```bash
make test
```

---

## 🧹 Clean Generated Files

To remove generated GraphQL and backup files:

```bash
make clean
```

---

## 🔄 Code Generation Logic

Custom logic in the `Makefile` ensures:

* Your custom resolver code is **preserved** between generations
* `.bak` files used during merging are cleaned up automatically
* No duplicate closing braces (`}`) are appended

> ✅ You don’t need to manually edit `schema.resolvers.go` after codegen!

---

## 📐 Architecture Flow

```
GraphQL Request
      ↓
GraphQL Handler (Delivery)
      ↓
Usecase (Business Logic)
      ↓
Repository (Data Access)
      ↓
PostgreSQL
```

---

## 🧑‍💻 Requirements

* Go 1.16+
* Docker (for PostgreSQL container)
* PostgreSQL (optional if already installed)
* golangci-lint (for linting)

---

## 🙌 Contributions

Feel free to open issues or submit PRs to improve this boilerplate or add new features.

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE).

```

---

Let me know if you’d like me to also include:

- CI/CD setup (e.g., GitHub Actions)
- Dockerfile + `docker-compose.yml`
- GraphQL query examples
- Swagger/OpenAPI for REST fallback

I'm happy to extend the documentation further.
```
