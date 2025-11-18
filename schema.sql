-- PostgreSQL Schema for Financial Dashboard

-- Create database if not exists (run as superuser)
-- CREATE DATABASE fd;

-- UUID extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Income table
CREATE TABLE IF NOT EXISTS income (
    income_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    income_name VARCHAR(255),
    income_amount DECIMAL(15, 2),
    income_type VARCHAR(100),
    income_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Expense table
CREATE TABLE IF NOT EXISTS expense (
    expense_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    expense_name VARCHAR(255),
    expense_amount DECIMAL(15, 2),
    expense_type VARCHAR(100),
    expense_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_income_date ON income(income_date);
CREATE INDEX IF NOT EXISTS idx_income_type ON income(income_type);
CREATE INDEX IF NOT EXISTS idx_income_created_at ON income(created_at);

CREATE INDEX IF NOT EXISTS idx_expense_date ON expense(expense_date);
CREATE INDEX IF NOT EXISTS idx_expense_type ON expense(expense_type);
CREATE INDEX IF NOT EXISTS idx_expense_created_at ON expense(created_at);

-- Trigger function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Triggers for updating updated_at on UPDATE
CREATE TRIGGER update_income_updated_at BEFORE UPDATE ON income
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_expense_updated_at BEFORE UPDATE ON expense
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Sample data (optional, uncomment to insert)
/*
-- Sample income records
INSERT INTO income (income_name, income_amount, income_type, income_date) VALUES
    ('Salary', 5000.00, 'Employment', '2024-01-01 09:00:00'),
    ('Freelance Project', 1500.00, 'Freelance', '2024-01-15 14:30:00'),
    ('Investment Returns', 250.50, 'Investment', '2024-01-20 12:00:00');

-- Sample expense records
INSERT INTO expense (expense_name, expense_amount, expense_type, expense_date) VALUES
    ('Rent', 1200.00, 'Housing', '2024-01-01 08:00:00'),
    ('Groceries', 350.75, 'Food', '2024-01-05 18:30:00'),
    ('Internet Bill', 59.99, 'Utilities', '2024-01-10 10:00:00'),
    ('Gas', 85.50, 'Transportation', '2024-01-12 16:45:00');
*/

-- Useful queries

-- Get all income for a specific month
/*
SELECT * FROM income
WHERE EXTRACT(YEAR FROM income_date) = 2024
AND EXTRACT(MONTH FROM income_date) = 1
ORDER BY income_date DESC;
*/

-- Get all expenses for a specific month
/*
SELECT * FROM expense
WHERE EXTRACT(YEAR FROM expense_date) = 2024
AND EXTRACT(MONTH FROM expense_date) = 1
ORDER BY expense_date DESC;
*/

-- Calculate total income by type
/*
SELECT income_type, SUM(income_amount) as total
FROM income
GROUP BY income_type
ORDER BY total DESC;
*/

-- Calculate total expenses by type
/*
SELECT expense_type, SUM(expense_amount) as total
FROM expense
GROUP BY expense_type
ORDER BY total DESC;
*/

-- Get monthly income/expense summary
/*
SELECT
    DATE_TRUNC('month', income_date) as month,
    SUM(income_amount) as total_income
FROM income
GROUP BY month
ORDER BY month DESC;

SELECT
    DATE_TRUNC('month', expense_date) as month,
    SUM(expense_amount) as total_expenses
FROM expense
GROUP BY month
ORDER BY month DESC;
*/