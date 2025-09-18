-- This script sets up the required accounts for the double-entry ledger.
-- Run this against your database after the tables are created by GORM.

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    type VARCHAR(50) NOT NULL
);

INSERT INTO accounts (name, type) VALUES
('USER_STOCK_EQUITY', 'STOCK'),
('COMPANY_STOCK_INVENTORY', 'STOCK'),
('COMPANY_CASH', 'INR'),
('FEE_BROKERAGE_EXPENSE', 'INR'),
('FEE_STT_EXPENSE', 'INR'),
('FEE_GST_EXPENSE', 'INR')
ON CONFLICT (name) DO NOTHING;
