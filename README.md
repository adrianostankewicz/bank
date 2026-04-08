# Bank API

A banking wallet API built with Go, focused on clean architecture, domain separation, and scalability.

## Architecture

The project follows a domain-driven design approach with clean architecture principles, organized in the following layers:

- **Domain** — entities, value objects, and repository interfaces
- **Usecases** — business logic orchestration
- **Event** — event dispatcher and handlers
- **Infra** — HTTP transport layer

### Architecture Decision Records

- [ADR-001](docs/adr/ADR-001.md) — Balance as Account attribute
- [ADR-002](docs/adr/ADR-002.md) — Transaction as domain entity
- [ADR-003](docs/adr/ADR-003.md) — Separation between EventType and TransactionType
- [ADR-004](docs/adr/ADR-004.md) — Post-transfer reconciliation

## Key Design Decisions

- **Money as Value Object** — immutable, int64 in cents, no floating point
- **Transaction as source of truth** — balance is always derived from transactions, never stored directly
- **Event Dispatcher** — decoupled event routing, no switch statements in HTTP handlers
- **In-memory repositories** — clean interfaces ready for database implementation

## Requirements

- Go 1.21+
- [chi](https://github.com/go-chi/chi) — HTTP router

## Getting Started

```bash
# clone the repository
git clone https://github.com/adrianostankewicz/bank
cd bank

# install dependencies
go mod tidy

# run the server
go run cmd/main.go
```

Server will start on port `3000`.

## Running Tests

```bash
# run all tests
go test ./...

# run tests with coverage
go test ./... -cover
```