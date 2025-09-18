# Stocky Go API

This is the backend service for the Stocky rewards platform, built with Go. It provides a RESTful API to manage stock rewards for users, track holdings, and maintain a double-entry ledger for all financial transactions.

## Design & Architecture

This service was built with a focus on performance, type safety, and clean architecture.

*   **Language & Framework**: Written in **Go** using the **Gin** web framework for high-performance routing.
*   **Database**: Uses **PostgreSQL** for robust data storage and **GORM** as the ORM for interacting with the database.
*   **Data Integrity**: Employs the `decimal` library and `NUMERIC` database types to ensure financial precision and avoid floating-point errors.

---

## Deliverables

### 1. API Specifications

#### `POST /reward`
Records a stock reward for a user. The `idempotency_key` ensures that the same transaction cannot be processed twice.

*   **Request Payload:**
    ```
    {
      "user_id": "user-123",
      "stock_symbol": "RELIANCE",
      "quantity": 2.5,
      "idempotency_key": "unique-transaction-id-abc123"
    }
    ```
*   **Success Response (201 Created):**
    ```
    {
      "id": 1,
      "user_id": "user-123",
      "stock_symbol": "RELIANCE",
      "quantity": 2.5,
      "rewarded_at": "2025-09-19T05:05:00.123Z"
    }
    ```

#### `GET /today-stocks/{userId}`
Returns all stock rewards a user has received on the current day.

#### `GET /historical-inr/{userId}`
Returns the total INR value of a user's portfolio for all past days (currently mocked).

#### `GET /stats/{userId}`
Returns a summary of the user's activity, including shares rewarded today and the total current value of their portfolio.

#### `GET /portfolio/{userId}`
Returns a detailed breakdown of the user's current holdings, with quantity and current INR value for each stock.

### 2. Database Schema

The schema is designed to be relational and transactional.

*   **`users`**: Stores unique user identifiers.
*   **`rewards`**: Records every reward event. It includes `user_id`, `stock_symbol`, `quantity`, `rewarded_at`, and a unique `idempotency_key`.
*   **`accounts`**: A chart of accounts for the ledger system (e.g., `USER_STOCK_EQUITY`, `COMPANY_CASH`, `FEE_BROKERAGE_EXPENSE`).
*   **`ledger_entries`**: Contains the immutable debit and credit entries. Each entry is linked to an account and a transaction, forming a complete, balanced record of every financial event.

**Data Types:**
*   Stock quantities use `NUMERIC(18, 6)` for high precision, allowing for fractional shares.
*   INR amounts use `NUMERIC(18, 4)` to accurately represent currency values.

### 3. Edge Case & Scaling Strategy

*   **Duplicate Reward Events**: Handled by the `idempotency_key` with a `UNIQUE` constraint in the `rewards` table. The API first checks for the existence of this key before creating a new record.
*   **Stock Splits/Mergers**: These are complex corporate actions that must be handled by a separate, offline administrative script. The script would update the `quantity` in the `rewards` table for all affected users and create corresponding adjustment entries in the ledger.
*   **Rounding Errors**: Using the `shopspring/decimal` library in Go and `NUMERIC` types in PostgreSQL prevents floating-point precision errors in all financial calculations.
*   **Price API Downtime**: The `price_service` caches prices for a set duration (e.g., one hour). For a production system, this would be enhanced with a persistent cache like Redis and a fallback mechanism to serve stale data if the live API is down, with a flag in the response (`"is_stale": true`) to notify clients.
*   **Scaling**: The application is stateless, allowing it to be horizontally scaled by running multiple instances behind a load balancer. Database read replicas can be used to scale read-heavy endpoints like `/stats` and `/portfolio`.

---

## How to Run This Application

Follow these steps to get the API server running on your local machine.

### 1. Prerequisites
*   **Go** (version 1.21 or newer) installed.
*   **PostgreSQL** installed and running.

### 2. Database Setup
*   Open the **SQL Shell (psql)** and run the following commands:
    ```
    CREATE USER stocky_user WITH PASSWORD 'your_secure_password';
    CREATE DATABASE stocky_db OWNER stocky_user;
    ```

### 3. Configure Your Environment
*   In the project's root directory, create a file named `.env`.
*   Copy the following into it, using the password you just created:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=stocky_user
    DB_PASSWORD=your_secure_password
    DB_NAME=stocky_db
    SERVER_PORT=8080
    ```
> **Note:** The `.gitignore` file is configured to prevent the `.env` file from being committed to version control.

### 4. Run the Initial Migration & Dependencies
*   From your terminal, run the initial migration script to set up the ledger accounts:
    ```
    psql -h localhost -U stocky_user -d stocky_db -f scripts/migrate.sql
    ```
*   Then, download the Go libraries:
    ```
    go mod tidy
    ```

### 5. Run the Application
*   Start the server from the project root:
    ```
    go run cmd/stocky/main.go
    ```
The server will start on port `8080`, and your API is now live.
