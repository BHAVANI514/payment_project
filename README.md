# payment_project


# Internal Transfers System (HTTP Interface)

This project is a Golang-based financial transactions application that enables internal transfers between accounts. It provides HTTP endpoints to:
- Create accounts
- Query account balances
- Process transactions

The system uses PostgreSQL to persist account states and transaction logs.

---

## Features
- Create accounts with initial balances.
- Query account balances via HTTP GET.
- Transfer amounts between accounts while maintaining data integrity.
- Error handling for invalid inputs and insufficient funds.
- Clean code, modular structure, and easy-to-read logic.

---

## API Endpoints

### 1. Account Creation (POST /accounts)

**Request Body**
```json
{
  "account_id": 123,
  "initial_balance": "100.23344"
}
```

**Success Response**
```json
{}
```

**Error Response**
```json
{
  "error": "account already exists"
}
```

---

### 2. Account Query (GET /accounts/{account_id})

**Example Request**
```
GET /accounts/123
```

**Response**
```json
{
  "account_id": 123,
  "balance": "100.23344"
}
```

---

### 3. Transaction Submission (POST /transactions)

**Request Body**
```json
{
  "source_account_id": 123,
  "destination_account_id": 456,
  "amount": "100.12345"
}
```

**Success Response**
```json
{
  "message": "transaction successful"
}
```

**Error Responses**
```json
{
  "error": "insufficient funds"
}
```

or

```json
{
  "error": "account not found"
}
```

---

## Setup Instructions

### Prerequisites
- Go 1.22 or later
- PostgreSQL installed and running

### 1. Clone Repository
```bash
git clone <your-public-repo-url>
cd internal-transfers
```

### 2. Configure Environment
Create `.env` file:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=internal_transfers
```

### 3. Initialize Database
```bash
psql -U postgres -c "CREATE DATABASE internal_transfers;"
psql -U postgres internal_transfers < schema.sql
```

**schema.sql**
```sql
CREATE TABLE accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    source_account_id BIGINT NOT NULL,
    destination_account_id BIGINT NOT NULL,
    amount NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 4. Run Application
```bash
go mod tidy
go run main.go
```

---

## Example Console Output

### Application Startup
```
$ go run main.go
[INFO] Connecting to database...
[INFO] Database connection established.
[INFO] Applying database migrations...
[INFO] Migrations complete.
[INFO] Server started at http://localhost:8080
```

### Account Creation
```
[INFO] POST /accounts
[DEBUG] Request Body: {"account_id":123,"initial_balance":"100.23344"}
[INFO] Account 123 created with balance 100.23344
[INFO] Response: {}
```

### Account Query
```
[INFO] GET /accounts/123
[INFO] Account 123 found: Balance = 100.23344
[INFO] Response: {"account_id":123,"balance":"100.23344"}
```

### Transaction Submission
```
[INFO] POST /transactions
[DEBUG] Request Body: {"source_account_id":123,"destination_account_id":456,"amount":"50.12345"}
[INFO] Transaction started: 123 -> 456 amount 50.12345
[INFO] Updated Balances: Account 123 = 50.10999, Account 456 = 50.12345
[INFO] Transaction successful
[INFO] Response: {"message":"transaction successful"}
```

### Error Case: Insufficient Funds
```
[INFO] POST /transactions
[DEBUG] Request Body: {"source_account_id":123,"destination_account_id":456,"amount":"500"}
[ERROR] Insufficient funds in account 123: Available = 50.10999, Required = 500
[INFO] Response: {"error":"insufficient funds"}
```

---

## Assumptions
- Same currency for all accounts
- No authentication or authorization
- Numeric type for precise balances
- Transactions handled atomically

---

## License
MIT
