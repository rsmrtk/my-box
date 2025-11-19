-- Analytics Test Data
-- Run this SQL script to insert sample data for testing analytics endpoints

-- Clear existing test data (optional - be careful in production!)
-- DELETE FROM expense WHERE expense_id LIKE 'test_%';
-- DELETE FROM income WHERE income_id LIKE 'test_%';

-- Insert test income data for current and previous month
INSERT INTO income (income_id, income_name, income_amount, income_type, income_date, created_at) VALUES
-- Current month (November 2024)
('test_i001', 'November Salary', 8500.00, 'Employment', CURRENT_DATE, NOW()),
('test_i002', 'Freelance Project', 2000.00, 'Freelance', CURRENT_DATE - INTERVAL '5 days', NOW()),
('test_i003', 'Investment Dividend', 500.00, 'Investment', CURRENT_DATE - INTERVAL '10 days', NOW()),
('test_i004', 'Bonus', 1500.00, 'Employment', CURRENT_DATE - INTERVAL '7 days', NOW()),

-- Previous month (October 2024)
('test_i005', 'October Salary', 8500.00, 'Employment', CURRENT_DATE - INTERVAL '1 month', NOW()),
('test_i006', 'Side Business', 1300.00, 'Business', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '5 days', NOW()),
('test_i007', 'Freelance Work', 1500.00, 'Freelance', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '10 days', NOW()),

-- 2 months ago (September 2024)
('test_i008', 'September Salary', 8500.00, 'Employment', CURRENT_DATE - INTERVAL '2 months', NOW()),
('test_i009', 'Consulting', 3000.00, 'Freelance', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '5 days', NOW()),

-- 3 months ago (August 2024)
('test_i010', 'August Salary', 8500.00, 'Employment', CURRENT_DATE - INTERVAL '3 months', NOW()),
('test_i011', 'Project Payment', 2500.00, 'Freelance', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '5 days', NOW());

-- Insert test expense data with various categories and amounts
INSERT INTO expense (expense_id, expense_name, expense_amount, expense_type, expense_date, created_at) VALUES
-- Current month expenses (November 2024)
('test_e001', 'November Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '1 day', NOW()),
('test_e002', 'Groceries Week 1', 450.00, 'Food', CURRENT_DATE - INTERVAL '2 days', NOW()),
('test_e003', 'Gas Fill-up', 65.00, 'Transportation', CURRENT_DATE - INTERVAL '3 days', NOW()),
('test_e004', 'Electric Bill', 180.00, 'Utilities', CURRENT_DATE - INTERVAL '5 days', NOW()),
('test_e005', 'Restaurant Dinner', 85.00, 'Food', CURRENT_DATE - INTERVAL '6 days', NOW()),
('test_e006', 'Internet Bill', 70.00, 'Utilities', CURRENT_DATE - INTERVAL '7 days', NOW()),
('test_e007', 'Car Insurance', 200.00, 'Insurance', CURRENT_DATE - INTERVAL '8 days', NOW()),
('test_e008', 'Groceries Week 2', 380.00, 'Food', CURRENT_DATE - INTERVAL '9 days', NOW()),
('test_e009', 'Coffee Shop', 25.00, 'Food', CURRENT_DATE - INTERVAL '10 days', NOW()),
('test_e010', 'Gym Membership', 50.00, 'Health', CURRENT_DATE - INTERVAL '11 days', NOW()),
('test_e011', 'Phone Bill', 85.00, 'Utilities', CURRENT_DATE - INTERVAL '12 days', NOW()),
('test_e012', 'Streaming Services', 45.00, 'Entertainment', CURRENT_DATE - INTERVAL '13 days', NOW()),
('test_e013', 'Gas Fill-up 2', 70.00, 'Transportation', CURRENT_DATE - INTERVAL '14 days', NOW()),
('test_e014', 'Groceries Week 3', 420.00, 'Food', CURRENT_DATE - INTERVAL '15 days', NOW()),
('test_e015', 'Clothing', 150.00, 'Shopping', CURRENT_DATE - INTERVAL '16 days', NOW()),

-- Anomalies for current month
('test_e016', 'Annual Insurance Premium', 3600.00, 'Insurance', CURRENT_DATE - INTERVAL '5 days', NOW()), -- Major anomaly
('test_e017', 'Emergency Car Repair', 2500.00, 'Transportation', CURRENT_DATE - INTERVAL '10 days', NOW()), -- Major anomaly
('test_e018', 'Holiday Shopping Spree', 1200.00, 'Shopping', CURRENT_DATE - INTERVAL '7 days', NOW()), -- Medium anomaly
('test_e019', 'Premium Restaurant', 450.00, 'Food', CURRENT_DATE - INTERVAL '3 days', NOW()), -- Minor anomaly

-- Previous month expenses (October 2024)
('test_e020', 'October Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '1 day', NOW()),
('test_e021', 'Groceries', 420.00, 'Food', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '2 days', NOW()),
('test_e022', 'Gas', 60.00, 'Transportation', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '3 days', NOW()),
('test_e023', 'Electric Bill', 165.00, 'Utilities', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '5 days', NOW()),
('test_e024', 'Dining Out', 120.00, 'Food', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '6 days', NOW()),
('test_e025', 'Internet', 70.00, 'Utilities', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '7 days', NOW()),
('test_e026', 'Insurance', 200.00, 'Insurance', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '8 days', NOW()),
('test_e027', 'Groceries 2', 385.00, 'Food', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '9 days', NOW()),
('test_e028', 'Coffee', 30.00, 'Food', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '10 days', NOW()),
('test_e029', 'Entertainment', 75.00, 'Entertainment', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '11 days', NOW()),
('test_e030', 'Phone', 85.00, 'Utilities', CURRENT_DATE - INTERVAL '1 month' - INTERVAL '12 days', NOW()),

-- 2 months ago expenses (September 2024)
('test_e031', 'September Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '1 day', NOW()),
('test_e032', 'Groceries', 460.00, 'Food', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '2 days', NOW()),
('test_e033', 'Transportation', 180.00, 'Transportation', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '3 days', NOW()),
('test_e034', 'Utilities Total', 320.00, 'Utilities', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '5 days', NOW()),
('test_e035', 'Food & Dining', 550.00, 'Food', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '6 days', NOW()),
('test_e036', 'Insurance', 200.00, 'Insurance', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '8 days', NOW()),
('test_e037', 'Shopping', 200.00, 'Shopping', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '10 days', NOW()),
('test_e038', 'Health & Fitness', 100.00, 'Health', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '12 days', NOW()),
('test_e039', 'Entertainment', 90.00, 'Entertainment', CURRENT_DATE - INTERVAL '2 months' - INTERVAL '14 days', NOW()),

-- 3 months ago expenses (August 2024)
('test_e040', 'August Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '1 day', NOW()),
('test_e041', 'Groceries', 440.00, 'Food', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '2 days', NOW()),
('test_e042', 'Transportation', 150.00, 'Transportation', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '3 days', NOW()),
('test_e043', 'Utilities', 310.00, 'Utilities', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '5 days', NOW()),
('test_e044', 'Food', 480.00, 'Food', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '6 days', NOW()),
('test_e045', 'Insurance', 200.00, 'Insurance', CURRENT_DATE - INTERVAL '3 months' - INTERVAL '8 days', NOW()),

-- 4-6 months ago expenses (for trend analysis)
('test_e046', 'July Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '4 months', NOW()),
('test_e047', 'July Expenses', 2300.00, 'Food', CURRENT_DATE - INTERVAL '4 months', NOW()),
('test_e048', 'June Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '5 months', NOW()),
('test_e049', 'June Expenses', 2100.00, 'Food', CURRENT_DATE - INTERVAL '5 months', NOW()),
('test_e050', 'May Rent', 2500.00, 'Housing', CURRENT_DATE - INTERVAL '6 months', NOW()),
('test_e051', 'May Expenses', 2000.00, 'Food', CURRENT_DATE - INTERVAL '6 months', NOW());

-- Verify data insertion
SELECT 'Income Records Created:' as info, COUNT(*) as count FROM income WHERE income_id LIKE 'test_%'
UNION ALL
SELECT 'Expense Records Created:', COUNT(*) FROM expense WHERE expense_id LIKE 'test_%'
UNION ALL
SELECT 'Expense Categories:', COUNT(DISTINCT expense_type) FROM expense WHERE expense_id LIKE 'test_%';

-- Show category breakdown
SELECT
    expense_type as category,
    COUNT(*) as transactions,
    SUM(expense_amount) as total_amount,
    AVG(expense_amount) as avg_amount
FROM expense
WHERE expense_id LIKE 'test_%'
GROUP BY expense_type
ORDER BY total_amount DESC;

-- Show monthly summary
SELECT
    DATE_TRUNC('month', expense_date) as month,
    COUNT(*) as transactions,
    SUM(expense_amount) as total_expenses
FROM expense
WHERE expense_id LIKE 'test_%'
GROUP BY DATE_TRUNC('month', expense_date)
ORDER BY month DESC;