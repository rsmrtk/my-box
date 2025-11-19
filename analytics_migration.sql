-- =============================================================================
-- FINANCIAL DASHBOARD ANALYTICS MIGRATION
-- =============================================================================
-- Author: MyBox Analytics System
-- Description: Complete analytics views and functions for financial dashboard
-- Database: PostgreSQL
-- =============================================================================

-- =============================================================================
-- SECTION 1: INCOME ANALYTICS
-- =============================================================================

-- -----------------------------------------------------------------------------
-- 1.1 Monthly Income Summary View
-- -----------------------------------------------------------------------------
CREATE OR REPLACE VIEW v_monthly_income AS
SELECT
    DATE_TRUNC('month', income_date)::DATE as month,
    COUNT(*) as transaction_count,
    SUM(income_amount) as total_amount,
    AVG(income_amount) as avg_amount,
    MIN(income_amount) as min_amount,
    MAX(income_amount) as max_amount,
    ARRAY_AGG(DISTINCT income_type) as income_types
FROM income
GROUP BY DATE_TRUNC('month', income_date)
ORDER BY month DESC;

-- -----------------------------------------------------------------------------
-- 1.2 Annual Income Summary View
-- -----------------------------------------------------------------------------
CREATE OR REPLACE VIEW v_annual_income AS
SELECT
    DATE_PART('year', income_date)::INTEGER as year,
    COUNT(*) as transaction_count,
    SUM(income_amount) as total_amount,
    AVG(income_amount) as avg_amount,
    MIN(income_amount) as min_amount,
    MAX(income_amount) as max_amount
FROM income
GROUP BY DATE_PART('year', income_date)
ORDER BY year DESC;

-- -----------------------------------------------------------------------------
-- 1.3 Income Growth Analysis Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_income_growth_analysis(
    p_start_date DATE DEFAULT NULL,
    p_end_date DATE DEFAULT NULL
)
RETURNS TABLE(
    period_date DATE,
    period_type VARCHAR(10),
    total_income DECIMAL(15,2),
    previous_period_income DECIMAL(15,2),
    growth_amount DECIMAL(15,2),
    growth_percentage DECIMAL(5,2)
) AS $$
BEGIN
    RETURN QUERY
    WITH monthly_income AS (
        SELECT
            DATE_TRUNC('month', income_date)::DATE as month_date,
            SUM(income_amount) as monthly_total
        FROM income
        WHERE
            (p_start_date IS NULL OR income_date >= p_start_date) AND
            (p_end_date IS NULL OR income_date <= p_end_date)
        GROUP BY DATE_TRUNC('month', income_date)
    )
    SELECT
        month_date as period_date,
        'monthly'::VARCHAR(10) as period_type,
        monthly_total as total_income,
        LAG(monthly_total, 1) OVER (ORDER BY month_date) as previous_period_income,
        monthly_total - LAG(monthly_total, 1) OVER (ORDER BY month_date) as growth_amount,
        CASE
            WHEN LAG(monthly_total, 1) OVER (ORDER BY month_date) > 0 THEN
                ROUND(((monthly_total - LAG(monthly_total, 1) OVER (ORDER BY month_date)) /
                LAG(monthly_total, 1) OVER (ORDER BY month_date)) * 100, 2)
            ELSE NULL
        END as growth_percentage
    FROM monthly_income
    ORDER BY month_date DESC;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 1.4 Average Income Statistics Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_income_statistics(
    p_months_back INTEGER DEFAULT 6
)
RETURNS TABLE(
    metric_name VARCHAR(50),
    metric_value DECIMAL(15,2),
    period_description VARCHAR(100)
) AS $$
BEGIN
    RETURN QUERY
    -- Average income for last N months
    SELECT
        'avg_income_' || p_months_back || '_months'::VARCHAR(50),
        AVG(income_amount),
        'Average income for last ' || p_months_back || ' months'::VARCHAR(100)
    FROM income
    WHERE income_date >= CURRENT_DATE - INTERVAL '1 month' * p_months_back

    UNION ALL

    -- Daily average income
    SELECT
        'avg_daily_income'::VARCHAR(50),
        SUM(income_amount) / GREATEST(COUNT(DISTINCT DATE(income_date)), 1),
        'Average daily income'::VARCHAR(100)
    FROM income
    WHERE income_date >= CURRENT_DATE - INTERVAL '1 month' * p_months_back

    UNION ALL

    -- Weekly average income
    SELECT
        'avg_weekly_income'::VARCHAR(50),
        SUM(income_amount) / GREATEST(COUNT(DISTINCT DATE_TRUNC('week', income_date)), 1),
        'Average weekly income'::VARCHAR(100)
    FROM income
    WHERE income_date >= CURRENT_DATE - INTERVAL '1 month' * p_months_back

    UNION ALL

    -- Monthly average income
    SELECT
        'avg_monthly_income'::VARCHAR(50),
        SUM(income_amount) / GREATEST(COUNT(DISTINCT DATE_TRUNC('month', income_date)), 1),
        'Average monthly income'::VARCHAR(100)
    FROM income
    WHERE income_date >= CURRENT_DATE - INTERVAL '1 month' * p_months_back;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 1.5 Income Forecast Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_income_forecast(
    p_forecast_months INTEGER DEFAULT 3
)
RETURNS TABLE(
    forecast_month DATE,
    predicted_income DECIMAL(15,2),
    confidence_level VARCHAR(20),
    based_on_months INTEGER
) AS $$
DECLARE
    v_avg_monthly_income DECIMAL(15,2);
    v_std_dev DECIMAL(15,2);
    v_trend DECIMAL(15,2);
BEGIN
    -- Calculate average and standard deviation from last 6 months
    SELECT
        AVG(monthly_total),
        STDDEV(monthly_total)
    INTO v_avg_monthly_income, v_std_dev
    FROM (
        SELECT
            DATE_TRUNC('month', income_date) as month,
            SUM(income_amount) as monthly_total
        FROM income
        WHERE income_date >= CURRENT_DATE - INTERVAL '6 months'
        GROUP BY DATE_TRUNC('month', income_date)
    ) monthly_data;

    -- Calculate trend (simple linear regression slope)
    SELECT
        COALESCE((SUM((row_num - avg_row) * (monthly_total - avg_total)) /
                  NULLIF(SUM(POWER(row_num - avg_row, 2)), 0)), 0)
    INTO v_trend
    FROM (
        SELECT
            monthly_total,
            ROW_NUMBER() OVER (ORDER BY month) as row_num,
            AVG(ROW_NUMBER() OVER (ORDER BY month)) OVER () as avg_row,
            AVG(monthly_total) OVER () as avg_total
        FROM (
            SELECT
                DATE_TRUNC('month', income_date) as month,
                SUM(income_amount) as monthly_total
            FROM income
            WHERE income_date >= CURRENT_DATE - INTERVAL '6 months'
            GROUP BY DATE_TRUNC('month', income_date)
        ) monthly_data
    ) trend_data;

    -- Generate forecast
    RETURN QUERY
    SELECT
        (DATE_TRUNC('month', CURRENT_DATE) + INTERVAL '1 month' * generate_series(1, p_forecast_months))::DATE,
        v_avg_monthly_income + (v_trend * generate_series(1, p_forecast_months)),
        CASE
            WHEN v_std_dev < v_avg_monthly_income * 0.1 THEN 'High'
            WHEN v_std_dev < v_avg_monthly_income * 0.25 THEN 'Medium'
            ELSE 'Low'
        END::VARCHAR(20),
        6
    FROM generate_series(1, p_forecast_months);
END;
$$ LANGUAGE plpgsql;

-- =============================================================================
-- SECTION 2: EXPENSE ANALYTICS
-- =============================================================================

-- -----------------------------------------------------------------------------
-- 2.1 Monthly Expense Summary View
-- -----------------------------------------------------------------------------
CREATE OR REPLACE VIEW v_monthly_expense AS
SELECT
    DATE_TRUNC('month', expense_date)::DATE as month,
    COUNT(*) as transaction_count,
    SUM(expense_amount) as total_amount,
    AVG(expense_amount) as avg_amount,
    MIN(expense_amount) as min_amount,
    MAX(expense_amount) as max_amount,
    ARRAY_AGG(DISTINCT expense_type) as expense_types
FROM expense
GROUP BY DATE_TRUNC('month', expense_date)
ORDER BY month DESC;

-- -----------------------------------------------------------------------------
-- 2.2 Expense by Category View
-- -----------------------------------------------------------------------------
CREATE OR REPLACE VIEW v_expense_by_category AS
SELECT
    expense_type as category,
    COUNT(*) as transaction_count,
    SUM(expense_amount) as total_amount,
    AVG(expense_amount) as avg_amount,
    MIN(expense_amount) as min_amount,
    MAX(expense_amount) as max_amount,
    ROUND(100.0 * SUM(expense_amount) / SUM(SUM(expense_amount)) OVER (), 2) as percentage_of_total
FROM expense
GROUP BY expense_type
ORDER BY total_amount DESC;

-- -----------------------------------------------------------------------------
-- 2.3 Top Expenses Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_top_expenses(
    p_limit INTEGER DEFAULT 5,
    p_start_date DATE DEFAULT NULL,
    p_end_date DATE DEFAULT NULL
)
RETURNS TABLE(
    expense_id UUID,
    expense_name VARCHAR(255),
    expense_amount DECIMAL(15,2),
    expense_type VARCHAR(100),
    expense_date TIMESTAMP,
    rank_position INTEGER
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        e.expense_id,
        e.expense_name,
        e.expense_amount,
        e.expense_type,
        e.expense_date,
        ROW_NUMBER() OVER (ORDER BY e.expense_amount DESC)::INTEGER as rank_position
    FROM expense e
    WHERE
        (p_start_date IS NULL OR e.expense_date >= p_start_date) AND
        (p_end_date IS NULL OR e.expense_date <= p_end_date)
    ORDER BY e.expense_amount DESC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 2.4 Top Expense Categories Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_top_expense_categories(
    p_limit INTEGER DEFAULT 5,
    p_start_date DATE DEFAULT NULL,
    p_end_date DATE DEFAULT NULL
)
RETURNS TABLE(
    category VARCHAR(100),
    total_amount DECIMAL(15,2),
    transaction_count BIGINT,
    avg_amount DECIMAL(15,2),
    percentage_of_total DECIMAL(5,2),
    rank_position INTEGER
) AS $$
BEGIN
    RETURN QUERY
    WITH category_totals AS (
        SELECT
            expense_type,
            SUM(expense_amount) as total,
            COUNT(*) as count,
            AVG(expense_amount) as avg
        FROM expense
        WHERE
            (p_start_date IS NULL OR expense_date >= p_start_date) AND
            (p_end_date IS NULL OR expense_date <= p_end_date)
        GROUP BY expense_type
    )
    SELECT
        expense_type as category,
        total as total_amount,
        count as transaction_count,
        avg as avg_amount,
        ROUND(100.0 * total / SUM(total) OVER (), 2) as percentage_of_total,
        ROW_NUMBER() OVER (ORDER BY total DESC)::INTEGER as rank_position
    FROM category_totals
    ORDER BY total DESC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 2.5 Anomaly Detection Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_expense_anomalies(
    p_threshold_factor DECIMAL DEFAULT 1.5,
    p_lookback_months INTEGER DEFAULT 3
)
RETURNS TABLE(
    expense_id UUID,
    expense_name VARCHAR(255),
    expense_amount DECIMAL(15,2),
    expense_type VARCHAR(100),
    expense_date TIMESTAMP,
    category_avg DECIMAL(15,2),
    deviation_factor DECIMAL(5,2),
    anomaly_score VARCHAR(20)
) AS $$
BEGIN
    RETURN QUERY
    WITH category_stats AS (
        SELECT
            expense_type,
            AVG(expense_amount) as avg_amount,
            STDDEV(expense_amount) as std_dev
        FROM expense
        WHERE expense_date >= CURRENT_DATE - INTERVAL '1 month' * p_lookback_months
        GROUP BY expense_type
    )
    SELECT
        e.expense_id,
        e.expense_name,
        e.expense_amount,
        e.expense_type,
        e.expense_date,
        cs.avg_amount as category_avg,
        ROUND(e.expense_amount / NULLIF(cs.avg_amount, 0), 2) as deviation_factor,
        CASE
            WHEN e.expense_amount > cs.avg_amount + (cs.std_dev * 3) THEN 'Critical'
            WHEN e.expense_amount > cs.avg_amount + (cs.std_dev * 2) THEN 'High'
            WHEN e.expense_amount > cs.avg_amount + (cs.std_dev * p_threshold_factor) THEN 'Medium'
            ELSE 'Normal'
        END::VARCHAR(20) as anomaly_score
    FROM expense e
    JOIN category_stats cs ON e.expense_type = cs.expense_type
    WHERE
        e.expense_date >= CURRENT_DATE - INTERVAL '1 month' AND
        e.expense_amount > cs.avg_amount * p_threshold_factor
    ORDER BY deviation_factor DESC;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 2.6 Expense Trend Analysis Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_expense_trend_analysis(
    p_months_back INTEGER DEFAULT 6
)
RETURNS TABLE(
    month_date DATE,
    total_expense DECIMAL(15,2),
    previous_month_expense DECIMAL(15,2),
    change_amount DECIMAL(15,2),
    change_percentage DECIMAL(5,2),
    trend_direction VARCHAR(20),
    moving_avg_3m DECIMAL(15,2)
) AS $$
BEGIN
    RETURN QUERY
    WITH monthly_expenses AS (
        SELECT
            DATE_TRUNC('month', expense_date)::DATE as month,
            SUM(expense_amount) as total
        FROM expense
        WHERE expense_date >= CURRENT_DATE - INTERVAL '1 month' * (p_months_back + 3)
        GROUP BY DATE_TRUNC('month', expense_date)
    )
    SELECT
        month as month_date,
        total as total_expense,
        LAG(total, 1) OVER (ORDER BY month) as previous_month_expense,
        total - LAG(total, 1) OVER (ORDER BY month) as change_amount,
        CASE
            WHEN LAG(total, 1) OVER (ORDER BY month) > 0 THEN
                ROUND(((total - LAG(total, 1) OVER (ORDER BY month)) /
                LAG(total, 1) OVER (ORDER BY month)) * 100, 2)
            ELSE NULL
        END as change_percentage,
        CASE
            WHEN total > LAG(total, 1) OVER (ORDER BY month) THEN 'Increasing'
            WHEN total < LAG(total, 1) OVER (ORDER BY month) THEN 'Decreasing'
            ELSE 'Stable'
        END::VARCHAR(20) as trend_direction,
        AVG(total) OVER (ORDER BY month ROWS BETWEEN 2 PRECEDING AND CURRENT ROW) as moving_avg_3m
    FROM monthly_expenses
    WHERE month >= CURRENT_DATE - INTERVAL '1 month' * p_months_back
    ORDER BY month DESC;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 2.7 Share of Wallet Analysis Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_share_of_wallet(
    p_start_date DATE DEFAULT NULL,
    p_end_date DATE DEFAULT NULL
)
RETURNS TABLE(
    category VARCHAR(100),
    total_amount DECIMAL(15,2),
    transaction_count BIGINT,
    avg_transaction DECIMAL(15,2),
    share_percentage DECIMAL(5,2),
    cumulative_percentage DECIMAL(5,2),
    category_rank INTEGER
) AS $$
BEGIN
    RETURN QUERY
    WITH category_summary AS (
        SELECT
            expense_type,
            SUM(expense_amount) as total,
            COUNT(*) as count,
            AVG(expense_amount) as avg_amount
        FROM expense
        WHERE
            (p_start_date IS NULL OR expense_date >= p_start_date) AND
            (p_end_date IS NULL OR expense_date <= p_end_date)
        GROUP BY expense_type
    ),
    with_percentages AS (
        SELECT
            expense_type,
            total,
            count,
            avg_amount,
            ROUND(100.0 * total / SUM(total) OVER (), 2) as share_pct,
            ROW_NUMBER() OVER (ORDER BY total DESC) as rank
        FROM category_summary
    )
    SELECT
        expense_type as category,
        total as total_amount,
        count as transaction_count,
        avg_amount as avg_transaction,
        share_pct as share_percentage,
        SUM(share_pct) OVER (ORDER BY rank) as cumulative_percentage,
        rank::INTEGER as category_rank
    FROM with_percentages
    ORDER BY rank;
END;
$$ LANGUAGE plpgsql;

-- =============================================================================
-- SECTION 3: CASH FLOW ANALYTICS
-- =============================================================================

-- -----------------------------------------------------------------------------
-- 3.1 Cash Flow Summary View
-- -----------------------------------------------------------------------------
CREATE OR REPLACE VIEW v_cash_flow_summary AS
WITH income_monthly AS (
    SELECT
        DATE_TRUNC('month', income_date)::DATE as month,
        SUM(income_amount) as income_total
    FROM income
    GROUP BY DATE_TRUNC('month', income_date)
),
expense_monthly AS (
    SELECT
        DATE_TRUNC('month', expense_date)::DATE as month,
        SUM(expense_amount) as expense_total
    FROM expense
    GROUP BY DATE_TRUNC('month', expense_date)
)
SELECT
    COALESCE(i.month, e.month) as month,
    COALESCE(i.income_total, 0) as total_income,
    COALESCE(e.expense_total, 0) as total_expense,
    COALESCE(i.income_total, 0) - COALESCE(e.expense_total, 0) as net_cash_flow,
    CASE
        WHEN COALESCE(e.expense_total, 0) > 0 THEN
            ROUND(COALESCE(i.income_total, 0) / e.expense_total, 2)
        ELSE NULL
    END as income_expense_ratio,
    CASE
        WHEN COALESCE(i.income_total, 0) > 0 THEN
            ROUND(((COALESCE(i.income_total, 0) - COALESCE(e.expense_total, 0)) / i.income_total) * 100, 2)
        ELSE NULL
    END as savings_rate_percentage
FROM income_monthly i
FULL OUTER JOIN expense_monthly e ON i.month = e.month
ORDER BY month DESC;

-- -----------------------------------------------------------------------------
-- 3.2 Daily Cash Flow View
-- -----------------------------------------------------------------------------
CREATE OR REPLACE VIEW v_daily_cash_flow AS
WITH daily_income AS (
    SELECT
        DATE(income_date) as date,
        SUM(income_amount) as income_total
    FROM income
    GROUP BY DATE(income_date)
),
daily_expense AS (
    SELECT
        DATE(expense_date) as date,
        SUM(expense_amount) as expense_total
    FROM expense
    GROUP BY DATE(expense_date)
)
SELECT
    COALESCE(i.date, e.date) as date,
    COALESCE(i.income_total, 0) as income,
    COALESCE(e.expense_total, 0) as expense,
    COALESCE(i.income_total, 0) - COALESCE(e.expense_total, 0) as net_flow
FROM daily_income i
FULL OUTER JOIN daily_expense e ON i.date = e.date
ORDER BY date DESC;

-- -----------------------------------------------------------------------------
-- 3.3 Cash Flow Forecast Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_cash_flow_forecast(
    p_forecast_days INTEGER DEFAULT 30
)
RETURNS TABLE(
    forecast_date DATE,
    predicted_income DECIMAL(15,2),
    predicted_expense DECIMAL(15,2),
    predicted_net_flow DECIMAL(15,2),
    confidence_level VARCHAR(20),
    cumulative_net_flow DECIMAL(15,2)
) AS $$
DECLARE
    v_avg_daily_income DECIMAL(15,2);
    v_avg_daily_expense DECIMAL(15,2);
    v_income_std_dev DECIMAL(15,2);
    v_expense_std_dev DECIMAL(15,2);
    v_current_balance DECIMAL(15,2);
BEGIN
    -- Calculate average daily income and expense from last 90 days
    SELECT
        SUM(income_amount) / GREATEST(COUNT(DISTINCT DATE(income_date)), 1),
        STDDEV(daily_total)
    INTO v_avg_daily_income, v_income_std_dev
    FROM (
        SELECT DATE(income_date) as day, SUM(income_amount) as daily_total
        FROM income
        WHERE income_date >= CURRENT_DATE - INTERVAL '90 days'
        GROUP BY DATE(income_date)
    ) daily_income;

    SELECT
        SUM(expense_amount) / GREATEST(COUNT(DISTINCT DATE(expense_date)), 1),
        STDDEV(daily_total)
    INTO v_avg_daily_expense, v_expense_std_dev
    FROM (
        SELECT DATE(expense_date) as day, SUM(expense_amount) as daily_total
        FROM expense
        WHERE expense_date >= CURRENT_DATE - INTERVAL '90 days'
        GROUP BY DATE(expense_date)
    ) daily_expense;

    -- Calculate current balance (total income - total expense)
    SELECT
        COALESCE(SUM(income_amount), 0) - COALESCE((SELECT SUM(expense_amount) FROM expense), 0)
    INTO v_current_balance
    FROM income;

    -- Generate forecast
    RETURN QUERY
    WITH forecast_data AS (
        SELECT
            CURRENT_DATE + generate_series(1, p_forecast_days) as f_date,
            v_avg_daily_income as f_income,
            v_avg_daily_expense as f_expense,
            v_avg_daily_income - v_avg_daily_expense as f_net_flow,
            CASE
                WHEN v_income_std_dev < v_avg_daily_income * 0.2 AND
                     v_expense_std_dev < v_avg_daily_expense * 0.2 THEN 'High'
                WHEN v_income_std_dev < v_avg_daily_income * 0.4 AND
                     v_expense_std_dev < v_avg_daily_expense * 0.4 THEN 'Medium'
                ELSE 'Low'
            END::VARCHAR(20) as f_confidence
        FROM generate_series(1, p_forecast_days)
    )
    SELECT
        f_date,
        f_income,
        f_expense,
        f_net_flow,
        f_confidence,
        v_current_balance + SUM(f_net_flow) OVER (ORDER BY f_date) as cumulative_net_flow
    FROM forecast_data
    ORDER BY f_date;
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 3.4 Financial Stability Coefficient Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_financial_stability_coefficient(
    p_months_back INTEGER DEFAULT 6
)
RETURNS TABLE(
    metric_name VARCHAR(100),
    metric_value DECIMAL(15,2),
    interpretation VARCHAR(200),
    status VARCHAR(20)
) AS $$
DECLARE
    v_avg_income DECIMAL(15,2);
    v_avg_expense DECIMAL(15,2);
    v_income_variance DECIMAL(15,2);
    v_expense_variance DECIMAL(15,2);
    v_stability_ratio DECIMAL(15,2);
    v_savings_rate DECIMAL(15,2);
    v_coverage_ratio DECIMAL(15,2);
BEGIN
    -- Calculate averages
    SELECT AVG(monthly_income), VARIANCE(monthly_income)
    INTO v_avg_income, v_income_variance
    FROM (
        SELECT DATE_TRUNC('month', income_date) as month, SUM(income_amount) as monthly_income
        FROM income
        WHERE income_date >= CURRENT_DATE - INTERVAL '1 month' * p_months_back
        GROUP BY DATE_TRUNC('month', income_date)
    ) income_data;

    SELECT AVG(monthly_expense), VARIANCE(monthly_expense)
    INTO v_avg_expense, v_expense_variance
    FROM (
        SELECT DATE_TRUNC('month', expense_date) as month, SUM(expense_amount) as monthly_expense
        FROM expense
        WHERE expense_date >= CURRENT_DATE - INTERVAL '1 month' * p_months_back
        GROUP BY DATE_TRUNC('month', expense_date)
    ) expense_data;

    -- Calculate ratios
    v_stability_ratio := CASE
        WHEN v_avg_expense > 0 THEN v_avg_income / v_avg_expense
        ELSE NULL
    END;

    v_savings_rate := CASE
        WHEN v_avg_income > 0 THEN ((v_avg_income - v_avg_expense) / v_avg_income) * 100
        ELSE 0
    END;

    v_coverage_ratio := CASE
        WHEN v_avg_expense > 0 THEN (v_avg_income - v_avg_expense) / v_avg_expense
        ELSE NULL
    END;

    RETURN QUERY
    -- Income to Expense Ratio
    SELECT
        'Income to Expense Ratio'::VARCHAR(100),
        v_stability_ratio,
        CASE
            WHEN v_stability_ratio >= 1.5 THEN 'Excellent financial stability - Income significantly exceeds expenses'
            WHEN v_stability_ratio >= 1.2 THEN 'Good financial stability - Comfortable margin'
            WHEN v_stability_ratio >= 1.0 THEN 'Adequate stability - Income covers expenses'
            WHEN v_stability_ratio >= 0.9 THEN 'Warning - Expenses nearly exceed income'
            ELSE 'Critical - Expenses exceed income'
        END::VARCHAR(200),
        CASE
            WHEN v_stability_ratio >= 1.2 THEN 'Healthy'
            WHEN v_stability_ratio >= 1.0 THEN 'Stable'
            WHEN v_stability_ratio >= 0.9 THEN 'Warning'
            ELSE 'Critical'
        END::VARCHAR(20)

    UNION ALL

    -- Savings Rate
    SELECT
        'Savings Rate (%)'::VARCHAR(100),
        v_savings_rate,
        CASE
            WHEN v_savings_rate >= 30 THEN 'Excellent savings rate - Strong financial future'
            WHEN v_savings_rate >= 20 THEN 'Good savings rate - Building wealth'
            WHEN v_savings_rate >= 10 THEN 'Adequate savings rate - Meeting minimum recommendations'
            WHEN v_savings_rate >= 5 THEN 'Low savings rate - Consider reducing expenses'
            WHEN v_savings_rate >= 0 THEN 'Minimal savings - Financial vulnerability'
            ELSE 'Negative savings - Spending exceeds income'
        END::VARCHAR(200),
        CASE
            WHEN v_savings_rate >= 20 THEN 'Excellent'
            WHEN v_savings_rate >= 10 THEN 'Good'
            WHEN v_savings_rate >= 5 THEN 'Fair'
            WHEN v_savings_rate >= 0 THEN 'Poor'
            ELSE 'Critical'
        END::VARCHAR(20)

    UNION ALL

    -- Income Volatility
    SELECT
        'Income Volatility Index'::VARCHAR(100),
        CASE
            WHEN v_avg_income > 0 THEN SQRT(v_income_variance) / v_avg_income * 100
            ELSE NULL
        END,
        CASE
            WHEN v_avg_income > 0 AND SQRT(v_income_variance) / v_avg_income < 0.1 THEN 'Very stable income'
            WHEN v_avg_income > 0 AND SQRT(v_income_variance) / v_avg_income < 0.25 THEN 'Moderate income stability'
            WHEN v_avg_income > 0 AND SQRT(v_income_variance) / v_avg_income < 0.5 THEN 'Variable income'
            ELSE 'Highly volatile income'
        END::VARCHAR(200),
        CASE
            WHEN v_avg_income > 0 AND SQRT(v_income_variance) / v_avg_income < 0.25 THEN 'Stable'
            WHEN v_avg_income > 0 AND SQRT(v_income_variance) / v_avg_income < 0.5 THEN 'Moderate'
            ELSE 'Volatile'
        END::VARCHAR(20)

    UNION ALL

    -- Coverage Ratio
    SELECT
        'Expense Coverage Ratio'::VARCHAR(100),
        v_coverage_ratio,
        CASE
            WHEN v_coverage_ratio >= 0.5 THEN 'Strong buffer - Can handle 50%+ expense increase'
            WHEN v_coverage_ratio >= 0.25 THEN 'Good buffer - Can handle 25% expense increase'
            WHEN v_coverage_ratio >= 0.1 THEN 'Small buffer - Limited flexibility'
            WHEN v_coverage_ratio >= 0 THEN 'Minimal buffer - No room for unexpected expenses'
            ELSE 'Deficit - Unable to cover current expenses'
        END::VARCHAR(200),
        CASE
            WHEN v_coverage_ratio >= 0.25 THEN 'Strong'
            WHEN v_coverage_ratio >= 0.1 THEN 'Adequate'
            WHEN v_coverage_ratio >= 0 THEN 'Weak'
            ELSE 'Deficit'
        END::VARCHAR(20);
END;
$$ LANGUAGE plpgsql;

-- -----------------------------------------------------------------------------
-- 3.5 Emergency Fund Analysis Function
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_emergency_fund_analysis()
RETURNS TABLE(
    metric_name VARCHAR(100),
    metric_value DECIMAL(15,2),
    recommendation VARCHAR(200)
) AS $$
DECLARE
    v_monthly_expense DECIMAL(15,2);
    v_current_savings DECIMAL(15,2);
    v_months_covered DECIMAL(5,2);
BEGIN
    -- Calculate average monthly expense
    SELECT AVG(monthly_expense)
    INTO v_monthly_expense
    FROM (
        SELECT DATE_TRUNC('month', expense_date) as month, SUM(expense_amount) as monthly_expense
        FROM expense
        WHERE expense_date >= CURRENT_DATE - INTERVAL '6 months'
        GROUP BY DATE_TRUNC('month', expense_date)
    ) expense_data;

    -- Calculate current savings (total income - total expense)
    SELECT
        COALESCE(SUM(income_amount), 0) - COALESCE((SELECT SUM(expense_amount) FROM expense), 0)
    INTO v_current_savings
    FROM income;

    -- Calculate months covered
    v_months_covered := CASE
        WHEN v_monthly_expense > 0 THEN v_current_savings / v_monthly_expense
        ELSE NULL
    END;

    RETURN QUERY
    SELECT
        'Current Savings'::VARCHAR(100),
        v_current_savings,
        'Total accumulated savings'::VARCHAR(200)

    UNION ALL

    SELECT
        'Monthly Expense Average'::VARCHAR(100),
        v_monthly_expense,
        'Average monthly expenses (last 6 months)'::VARCHAR(200)

    UNION ALL

    SELECT
        'Emergency Fund Coverage (months)'::VARCHAR(100),
        v_months_covered,
        CASE
            WHEN v_months_covered >= 6 THEN 'Excellent - Recommended 6+ months coverage achieved'
            WHEN v_months_covered >= 3 THEN 'Good - Minimum 3 months coverage achieved'
            WHEN v_months_covered >= 1 THEN 'Fair - Build towards 3 months minimum'
            ELSE 'Critical - Immediate action needed to build emergency fund'
        END::VARCHAR(200)

    UNION ALL

    SELECT
        'Recommended Emergency Fund'::VARCHAR(100),
        v_monthly_expense * 6,
        'Target: 6 months of expenses'::VARCHAR(200)

    UNION ALL

    SELECT
        'Emergency Fund Gap'::VARCHAR(100),
        GREATEST((v_monthly_expense * 6) - v_current_savings, 0),
        'Amount needed to reach 6-month target'::VARCHAR(200);
END;
$$ LANGUAGE plpgsql;

-- =============================================================================
-- SECTION 4: COMPREHENSIVE DASHBOARD FUNCTION
-- =============================================================================

-- -----------------------------------------------------------------------------
-- 4.1 Complete Financial Dashboard Summary
-- -----------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION fn_financial_dashboard_summary(
    p_period_start DATE DEFAULT NULL,
    p_period_end DATE DEFAULT NULL
)
RETURNS TABLE(
    section VARCHAR(50),
    metric VARCHAR(100),
    value DECIMAL(15,2),
    percentage_change DECIMAL(5,2),
    trend VARCHAR(20),
    description VARCHAR(200)
) AS $$
BEGIN
    -- If no dates provided, use current month
    IF p_period_start IS NULL THEN
        p_period_start := DATE_TRUNC('month', CURRENT_DATE);
    END IF;

    IF p_period_end IS NULL THEN
        p_period_end := CURRENT_DATE;
    END IF;

    RETURN QUERY
    -- Income Section
    WITH current_period_income AS (
        SELECT
            SUM(income_amount) as total,
            AVG(income_amount) as average,
            COUNT(*) as count
        FROM income
        WHERE income_date BETWEEN p_period_start AND p_period_end
    ),
    previous_period_income AS (
        SELECT
            SUM(income_amount) as total
        FROM income
        WHERE income_date BETWEEN
            p_period_start - (p_period_end - p_period_start + INTERVAL '1 day') AND
            p_period_start - INTERVAL '1 day'
    ),
    -- Expense Section
    current_period_expense AS (
        SELECT
            SUM(expense_amount) as total,
            AVG(expense_amount) as average,
            COUNT(*) as count
        FROM expense
        WHERE expense_date BETWEEN p_period_start AND p_period_end
    ),
    previous_period_expense AS (
        SELECT
            SUM(expense_amount) as total
        FROM expense
        WHERE expense_date BETWEEN
            p_period_start - (p_period_end - p_period_start + INTERVAL '1 day') AND
            p_period_start - INTERVAL '1 day'
    )
    -- Income Metrics
    SELECT
        'Income'::VARCHAR(50),
        'Total Income'::VARCHAR(100),
        cpi.total,
        CASE
            WHEN ppi.total > 0 THEN ROUND(((cpi.total - ppi.total) / ppi.total) * 100, 2)
            ELSE NULL
        END,
        CASE
            WHEN cpi.total > COALESCE(ppi.total, 0) THEN 'Up'
            WHEN cpi.total < COALESCE(ppi.total, 0) THEN 'Down'
            ELSE 'Stable'
        END::VARCHAR(20),
        'Total income for the period'::VARCHAR(200)
    FROM current_period_income cpi, previous_period_income ppi

    UNION ALL

    SELECT
        'Income'::VARCHAR(50),
        'Average Transaction'::VARCHAR(100),
        cpi.average,
        NULL,
        NULL::VARCHAR(20),
        'Average income per transaction'::VARCHAR(200)
    FROM current_period_income cpi

    UNION ALL

    SELECT
        'Income'::VARCHAR(50),
        'Transaction Count'::VARCHAR(100),
        cpi.count::DECIMAL,
        NULL,
        NULL::VARCHAR(20),
        'Number of income transactions'::VARCHAR(200)
    FROM current_period_income cpi

    UNION ALL

    -- Expense Metrics
    SELECT
        'Expense'::VARCHAR(50),
        'Total Expenses'::VARCHAR(100),
        cpe.total,
        CASE
            WHEN ppe.total > 0 THEN ROUND(((cpe.total - ppe.total) / ppe.total) * 100, 2)
            ELSE NULL
        END,
        CASE
            WHEN cpe.total > COALESCE(ppe.total, 0) THEN 'Up'
            WHEN cpe.total < COALESCE(ppe.total, 0) THEN 'Down'
            ELSE 'Stable'
        END::VARCHAR(20),
        'Total expenses for the period'::VARCHAR(200)
    FROM current_period_expense cpe, previous_period_expense ppe

    UNION ALL

    SELECT
        'Expense'::VARCHAR(50),
        'Average Transaction'::VARCHAR(100),
        cpe.average,
        NULL,
        NULL::VARCHAR(20),
        'Average expense per transaction'::VARCHAR(200)
    FROM current_period_expense cpe

    UNION ALL

    -- Cash Flow Metrics
    SELECT
        'CashFlow'::VARCHAR(50),
        'Net Cash Flow'::VARCHAR(100),
        cpi.total - cpe.total,
        NULL,
        CASE
            WHEN cpi.total - cpe.total > 0 THEN 'Positive'
            WHEN cpi.total - cpe.total < 0 THEN 'Negative'
            ELSE 'Neutral'
        END::VARCHAR(20),
        'Income minus expenses'::VARCHAR(200)
    FROM current_period_income cpi, current_period_expense cpe

    UNION ALL

    SELECT
        'CashFlow'::VARCHAR(50),
        'Savings Rate'::VARCHAR(100),
        CASE
            WHEN cpi.total > 0 THEN ROUND(((cpi.total - cpe.total) / cpi.total) * 100, 2)
            ELSE 0
        END,
        NULL,
        CASE
            WHEN cpi.total > 0 AND ((cpi.total - cpe.total) / cpi.total) > 0.2 THEN 'Excellent'
            WHEN cpi.total > 0 AND ((cpi.total - cpe.total) / cpi.total) > 0.1 THEN 'Good'
            WHEN cpi.total > 0 AND ((cpi.total - cpe.total) / cpi.total) > 0 THEN 'Fair'
            ELSE 'Poor'
        END::VARCHAR(20),
        'Percentage of income saved'::VARCHAR(200)
    FROM current_period_income cpi, current_period_expense cpe

    UNION ALL

    SELECT
        'CashFlow'::VARCHAR(50),
        'Financial Stability'::VARCHAR(100),
        CASE
            WHEN cpe.total > 0 THEN ROUND(cpi.total / cpe.total, 2)
            ELSE NULL
        END,
        NULL,
        CASE
            WHEN cpe.total > 0 AND cpi.total / cpe.total >= 1.2 THEN 'Strong'
            WHEN cpe.total > 0 AND cpi.total / cpe.total >= 1.0 THEN 'Stable'
            WHEN cpe.total > 0 AND cpi.total / cpe.total >= 0.9 THEN 'Warning'
            ELSE 'Critical'
        END::VARCHAR(20),
        'Income to expense ratio'::VARCHAR(200)
    FROM current_period_income cpi, current_period_expense cpe

    ORDER BY section, metric;
END;
$$ LANGUAGE plpgsql;

-- =============================================================================
-- SECTION 5: INDEXES FOR PERFORMANCE
-- =============================================================================

-- Create indexes if they don't exist
CREATE INDEX IF NOT EXISTS idx_income_date_amount ON income(income_date, income_amount);
CREATE INDEX IF NOT EXISTS idx_income_type_date ON income(income_type, income_date);
CREATE INDEX IF NOT EXISTS idx_expense_date_amount ON expense(expense_date, expense_amount);
CREATE INDEX IF NOT EXISTS idx_expense_type_date ON expense(expense_type, expense_date);

-- Create composite indexes for common queries
CREATE INDEX IF NOT EXISTS idx_income_month_year ON income(
    DATE_TRUNC('month', income_date),
    DATE_PART('year', income_date)
);

CREATE INDEX IF NOT EXISTS idx_expense_month_year ON expense(
    DATE_TRUNC('month', expense_date),
    DATE_PART('year', expense_date)
);

-- =============================================================================
-- SECTION 6: PERMISSIONS (Adjust as needed)
-- =============================================================================

-- Grant permissions to application user (replace 'app_user' with your actual user)
-- GRANT SELECT ON ALL TABLES IN SCHEMA public TO app_user;
-- GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO app_user;

-- =============================================================================
-- END OF MIGRATION
-- =============================================================================

-- To execute this migration:
-- psql -U your_username -d your_database -f analytics_migration.sql

-- To rollback this migration, save a rollback script with:
-- DROP VIEW IF EXISTS v_monthly_income CASCADE;
-- DROP VIEW IF EXISTS v_annual_income CASCADE;
-- DROP VIEW IF EXISTS v_monthly_expense CASCADE;
-- DROP VIEW IF EXISTS v_expense_by_category CASCADE;
-- DROP VIEW IF EXISTS v_cash_flow_summary CASCADE;
-- DROP VIEW IF EXISTS v_daily_cash_flow CASCADE;
-- DROP FUNCTION IF EXISTS fn_income_growth_analysis CASCADE;
-- DROP FUNCTION IF EXISTS fn_income_statistics CASCADE;
-- DROP FUNCTION IF EXISTS fn_income_forecast CASCADE;
-- DROP FUNCTION IF EXISTS fn_top_expenses CASCADE;
-- DROP FUNCTION IF EXISTS fn_top_expense_categories CASCADE;
-- DROP FUNCTION IF EXISTS fn_expense_anomalies CASCADE;
-- DROP FUNCTION IF EXISTS fn_expense_trend_analysis CASCADE;
-- DROP FUNCTION IF EXISTS fn_share_of_wallet CASCADE;
-- DROP FUNCTION IF EXISTS fn_cash_flow_forecast CASCADE;
-- DROP FUNCTION IF EXISTS fn_financial_stability_coefficient CASCADE;
-- DROP FUNCTION IF EXISTS fn_emergency_fund_analysis CASCADE;
-- DROP FUNCTION IF EXISTS fn_financial_dashboard_summary CASCADE;