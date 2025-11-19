-- =============================================================================
-- TABLES CREATION FOR FINANCIAL DASHBOARD
-- =============================================================================
-- Database: PostgreSQL
-- =============================================================================

-- Drop existing tables if needed
DROP TABLE IF EXISTS expense CASCADE;
DROP TABLE IF EXISTS income CASCADE;

-- =============================================================================
-- INCOME TABLE
-- =============================================================================
CREATE TABLE income (
    income_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    income_name VARCHAR(255) NOT NULL,
    income_amount DECIMAL(15,2) NOT NULL CHECK (income_amount > 0),
    income_type VARCHAR(100),
    income_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX idx_income_date ON income(income_date);
CREATE INDEX idx_income_type ON income(income_type);
CREATE INDEX idx_income_amount ON income(income_amount);
CREATE INDEX idx_income_created_at ON income(created_at);

-- =============================================================================
-- EXPENSE TABLE
-- =============================================================================
CREATE TABLE expense (
    expense_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    expense_name VARCHAR(255) NOT NULL,
    expense_amount DECIMAL(15,2) NOT NULL CHECK (expense_amount > 0),
    expense_type VARCHAR(100),
    expense_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX idx_expense_date ON expense(expense_date);
CREATE INDEX idx_expense_type ON expense(expense_type);
CREATE INDEX idx_expense_amount ON expense(expense_amount);
CREATE INDEX idx_expense_created_at ON expense(created_at);

-- =============================================================================
-- TRIGGER FOR UPDATED_AT
-- =============================================================================

-- Create function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers for income table
CREATE TRIGGER update_income_updated_at
    BEFORE UPDATE ON income
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Create triggers for expense table
CREATE TRIGGER update_expense_updated_at
    BEFORE UPDATE ON expense
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- =============================================================================
-- SAMPLE DATA FOR TESTING (Optional - remove in production)
-- =============================================================================

-- Sample income data
INSERT INTO income (income_name, income_amount, income_type, income_date) VALUES
('Salary January', 5000.00, 'Salary', '2024-01-15 10:00:00'),
('Freelance Project', 1500.00, 'Freelance', '2024-01-20 14:30:00'),
('Salary February', 5000.00, 'Salary', '2024-02-15 10:00:00'),
('Investment Returns', 300.00, 'Investment', '2024-02-25 16:00:00'),
('Salary March', 5200.00, 'Salary', '2024-03-15 10:00:00'),
('Bonus Q1', 1000.00, 'Bonus', '2024-03-31 17:00:00');

-- Sample expense data
INSERT INTO expense (expense_name, expense_amount, expense_type, expense_date) VALUES
('Rent January', 1500.00, 'Housing', '2024-01-05 09:00:00'),
('Groceries', 450.00, 'Food', '2024-01-10 11:00:00'),
('Utilities', 200.00, 'Utilities', '2024-01-15 12:00:00'),
('Transportation', 150.00, 'Transport', '2024-01-20 08:00:00'),
('Entertainment', 100.00, 'Entertainment', '2024-01-25 19:00:00'),
('Rent February', 1500.00, 'Housing', '2024-02-05 09:00:00'),
('Groceries', 500.00, 'Food', '2024-02-10 11:00:00'),
('Utilities', 180.00, 'Utilities', '2024-02-15 12:00:00'),
('Healthcare', 300.00, 'Healthcare', '2024-02-20 14:00:00'),
('Rent March', 1500.00, 'Housing', '2024-03-05 09:00:00'),
('Groceries', 480.00, 'Food', '2024-03-10 11:00:00'),
('Shopping', 350.00, 'Shopping', '2024-03-15 15:00:00');

-- =============================================================================
-- VERIFICATION QUERIES
-- =============================================================================

-- Check income table
SELECT COUNT(*) as income_count FROM income;

-- Check expense table
SELECT COUNT(*) as expense_count FROM expense;

-- Show income summary
SELECT
    DATE_TRUNC('month', income_date) as month,
    COUNT(*) as transactions,
    SUM(income_amount) as total_income
FROM income
GROUP BY DATE_TRUNC('month', income_date)
ORDER BY month;

-- Show expense summary by category
SELECT
    expense_type,
    COUNT(*) as transactions,
    SUM(expense_amount) as total_expense
FROM expense
GROUP BY expense_type
ORDER BY total_expense DESC;