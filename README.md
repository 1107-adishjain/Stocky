# Stocky Go API

This is the backend service for the Stocky rewards platform, built with Go. It provides a RESTful API to manage stock rewards for users, track holdings, and maintain a double-entry ledger for all financial transactions.

## How This Project Works

This service was built with a focus on performance, type safety, and clean architecture.

*   **Language & Framework**: Written in **Go** using the **Gin** web framework for high-performance routing.
*   **Database**: Uses **PostgreSQL** for robust data storage and **GORM** as the ORM for interacting with the database.
*   **Data Integrity**: Employs the `decimal` library and `NUMERIC` database types to ensure financial precision and avoid floating-point errors.
*   **Idempotency**: The `/reward` endpoint uses an `idempotency_key` to prevent accidental duplicate transactions.
*   **Ledger System**: A double-entry accounting system is implemented to track the flow of value (both stock and INR) for every reward, ensuring the books are always balanced.

## How to Run This Application

Follow these steps to get the API server running on your local machine.

### 1. Prerequisites
*   **Go** (version 1.21 or newer) installed.
*   **PostgreSQL** installed and running.

### 2. Database Setup
First, you need to create a dedicated user and database for the application.

*   Open the **SQL Shell (psql)** or your preferred PostgreSQL client and run the following commands:

    ```
    -- Create a user for the application
    CREATE USER stocky_user WITH PASSWORD 'your_secure_password';

    -- Create the database and assign ownership to the new user
    CREATE DATABASE stocky_db OWNER stocky_user;
    ```

### 3. Configure Your Environment
The application reads database credentials from an environment file.

*   In the project's root directory, create a file named `.env`.
*   Copy the following into it, making sure to use the password you just created:

    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=stocky_user
    DB_PASSWORD=your_secure_password
    DB_NAME=stocky_db
    SERVER_PORT=8080
    ```
> **IMPORTANT:** The `.gitignore` file is configured to prevent the `.env` file from ever being committed. **Never** share this file or commit it to version control.

### 4. Install Dependencies
Navigate to the project's root directory in your terminal and run:

This will download all the necessary libraries defined in the `go.mod` file.

### 5. Run the Application
You are now ready to start the server. From the project root, run:

The server will start, and you'll see a message indicating it's listening on port `8080`. Your API is now live and ready to accept requests. You can view the automatically generated API documentation by navigating to `http://localhost:8080/docs` in your browser.
