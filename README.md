# Regma

**Regulated Markets Architecture** — A compliance-first RWA tokenization platform in Go.

## Overview

Regma enables compliant tokenization of real-world assets (real estate, equities, debt instruments) with a focus on regulatory compliance. Inspired by ERC-3643.

## Architecture

Single Go binary with clean internal packages. Monolith first, microservices never (until proven necessary).

```
regma/
├── cmd/regma/           # Application entry point
├── internal/
│   ├── api/             # HTTP server & handlers
│   ├── config/          # Configuration management
│   ├── domain/          # Core domain models
│   ├── investor/        # Investor registry & KYC
│   ├── compliance/      # Rule engine & transfer validation
│   ├── asset/           # Asset registry & metadata
│   └── storage/         # Database implementations
├── pkg/                 # Shared utilities
└── migrations/          # SQL migrations
```

## Tech Stack

| Component | Choice |
|-----------|--------|
| Language | Go 1.24 |
| HTTP | net/http (stdlib) |
| Database | PostgreSQL |
| Migrations | golang-migrate |
| Config | Environment variables |
| Logging | slog (stdlib) |

## Quick Start

```bash
# Run the server
go run ./cmd/regma

# Health check
curl http://localhost:8080/health
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | 8080 | Server port |
| `ENVIRONMENT` | development | Environment name |
| `DATABASE_URL` | postgres://localhost:5432/regma | PostgreSQL connection |
| `JWT_SECRET` | change-me-in-production | JWT signing secret |

## Domain Models

### Investor
Wallet-based identity with KYC status and jurisdiction tracking.

### Asset
Tokenized real-world asset with compliance requirements.

### Compliance Rules
Transfer validation rules: jurisdiction checks, investor type restrictions, transfer limits.

## Roadmap

- [ ] PostgreSQL storage implementation
- [ ] Investor CRUD API
- [ ] Asset registry API
- [ ] Compliance rule engine
- [ ] Transfer validation
- [ ] JWT authentication
- [ ] Admin endpoints
- [ ] Ethereum connector (ERC-3643)

## Standards & Inspiration

- [ERC-3643: Permissioned Token Standard](https://github.com/ERC-3643)
- [W3C Verifiable Credentials](https://www.w3.org/TR/vc-data-model/)

## License

MIT
