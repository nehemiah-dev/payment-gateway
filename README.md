# FicMart Payment Gateway

A payment gateway built in Go for FicMart, a fictional e-commerce platform. It sits between FicMart's order service and a mock bank API, handling the payment lifecycle with strict state management, idempotency, and resilient failure handling.

## Architecture

```text
FicMart → Gateway API → Service Layer → Bank Client → Mock Bank
                ↓
           PostgreSQL
```

The gateway is organized into distinct layers:

* **Domain** — payment types and state transition rules with no external dependencies
* **Repository** — database access and queries
* **Bank** — HTTP client for the mock bank, including retry handling
* **Service** — payment operations and business rules
* **Handlers** — HTTP request and response handling
* **Middleware** — idempotency, request logging, and panic recovery

## Payment Lifecycle

```text
PENDING → AUTHORIZED → CAPTURED → REFUNDED
              ↓
            VOIDED
```

Payment transitions are enforced at the domain layer, regardless of what the bank API accepts.

For example:

* A `PENDING` payment can be authorized.
* An `AUTHORIZED` payment can be captured or voided.
* A `CAPTURED` payment can be refunded.
* Invalid transitions are rejected.

## Key Features

### Idempotency

All write operations require an `Idempotency-Key`.

If the same request is submitted again with the same key, the gateway returns the original response without processing the payment again.

Replayed responses include:

```text
X-Idempotent-Replayed: true
```

### Resilient Bank Communication

Transient failures such as network errors, timeouts, and 5xx responses are retried using exponential backoff with jitter.

Permanent failures such as expired cards and insufficient funds are not retried.

### Crash Safety

The gateway creates a `PENDING` payment record before calling the bank. This makes interrupted operations detectable and provides a basis for recovery.

## API

| Method | Endpoint                    | Description                   |
| ------ | --------------------------- | ----------------------------- |
| `POST` | `/payments/authorize`       | Reserve funds on a card       |
| `POST` | `/payments/{id}/capture`    | Capture an authorized payment |
| `POST` | `/payments/{id}/void`       | Cancel an authorization       |
| `POST` | `/payments/{id}/refund`     | Refund a captured payment     |
| `GET`  | `/payments?order_id=xxx`    | Get payments for an order     |
| `GET`  | `/payments?customer_id=xxx` | Get payments for a customer   |
| `GET`  | `/health`                   | Health check                  |

### Example

```bash
curl -X POST http://localhost:8088/payments/authorize \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: unique-key-001" \
  -d '{
    "order_id": "order-123",
    "customer_id": "customer-456",
    "amount": 5000,
    "currency": "USD",
    "card_number": "4111111111111111",
    "card_expiry": "12/2030",
    "card_cvv": "123"
  }'
```

Response:

```json
{
  "payment_id": "4262730c-3a1a-4489-b7a1-b761b7a7eae3",
  "order_id": "order-123",
  "customer_id": "customer-456",
  "amount": 5000,
  "currency": "USD",
  "state": "AUTHORIZED",
  "bank_auth_id": "auth_b424a765-...",
  "created_at": "2026-06-01T17:57:06Z",
  "authorized_at": "2026-06-01T17:57:06Z"
}
```

Amounts are represented in cents, so `5000` means `$50.00`.

## Test Cards

| Card Number        |   CVV | Expiry  | Balance | Use Case           |
| ------------------ | ----: | ------- | ------: | ------------------ |
| `4111111111111111` | `123` | 12/2030 | $10,000 | Happy path         |
| `4242424242424242` | `456` | 06/2030 |    $500 | Limited balance    |
| `5555555555554444` | `789` | 09/2030 |      $0 | Insufficient funds |
| `5105105105105100` | `321` | 03/2020 |  $5,000 | Expired card       |

## Project Structure

```text
payment-gateway/
├── bank/          # Mock bank API
├── docker/        # Docker Compose configuration
└── gateway/       # Payment gateway
```

## Running Locally

### Prerequisites

* Docker
* Docker Compose
* Go 1.25+
* PostgreSQL database

### 1. Clone

```bash
git clone https://github.com/Oliveszn/Payment-Gateway.git
cd Payment-Gateway
```

### 2. Configure

```bash
cp gateway/.env.example gateway/.env
```

Set `DATABASE_URL` in `gateway/.env`.

### 3. Start the mock bank

```bash
cd docker
docker compose up -d
```

### 4. Run migrations

```bash
cd ../gateway
make migrate-up
```

### 5. Start the gateway

```bash
make run
```

The gateway runs on:

```text
http://localhost:8088
```

The mock bank API documentation is available at:

```text
http://localhost:8787/docs
```

## Commands

```bash
make run             # Start the gateway
make test            # Run all tests
make test-short      # Run unit tests
make test-cover      # Run tests with coverage
make migrate-up      # Apply migrations
make migrate-down    # Roll back the last migration
make docker-up       # Start containers
make docker-down     # Stop containers
make docker-logs     # Tail gateway logs
make smoke-authorize # Run an authorization smoke test
make health          # Check gateway health
```

## Testing

```bash
make test
```

Current test coverage includes:

* **Domain tests** — payment state transitions and invalid transition handling
* **Service tests** — payment operations using mocked bank and repository dependencies

## Stack

* **Go 1.25**
* **PostgreSQL (Neon)**
* **golang-migrate**
* **Docker + Docker Compose**
* **testify**

## Design Decisions

See [`gateway/TRADEOFFS.md`](gateway/TRADEOFFS.md) for details on the architectural decisions, payment state management, idempotency approach, retry strategy, and considerations for a production implementation.
