# Financial Dashboard Analytics API Documentation

## Base URL
```
http://localhost:9595
```

## Analytics Endpoints

All analytics endpoints are grouped under `/analytics` prefix.

---

## 1. Income Analytics

### 1.1 General Income Analytics
```
GET /analytics/income
```

**Query Parameters:**
- `start_date` (optional): Start date for analysis (format: YYYY-MM-DD)
- `end_date` (optional): End date for analysis (format: YYYY-MM-DD)
- `months_back` (optional): Number of months to look back (default: 6)
- `analysis_type` (optional): Type of analysis - monthly, annual, growth, statistics, forecast

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/income?months_back=6&analysis_type=monthly"
```

### 1.2 Top Incomes
```
GET /analytics/income/top
```

**Query Parameters:**
- `limit` (optional): Number of top incomes to return (default: 5)
- `start_date` (optional): Start date filter
- `end_date` (optional): End date filter

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/income/top?limit=10"
```

### 1.3 Income Growth Analysis
```
GET /analytics/income/growth
```

**Query Parameters:**
- `start_date` (optional): Start date for analysis
- `end_date` (optional): End date for analysis

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/income/growth"
```

### 1.4 Income Forecast
```
GET /analytics/income/forecast
```

**Query Parameters:**
- `months_back` (optional): Historical months to base forecast on (default: 6)

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/income/forecast"
```

---

## 2. Expense Analytics

### 2.1 General Expense Analytics
```
GET /analytics/expense
```

**Query Parameters:**
- `start_date` (optional): Start date for analysis
- `end_date` (optional): End date for analysis
- `months_back` (optional): Number of months to look back (default: 6)
- `analysis_type` (optional): Type of analysis - monthly, categories, trends, anomalies, share_of_wallet

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/expense?analysis_type=categories"
```

### 2.2 Top Expenses
```
GET /analytics/expense/top
```

**Query Parameters:**
- `limit` (optional): Number of top expenses to return (default: 5)
- `start_date` (optional): Start date filter
- `end_date` (optional): End date filter

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/expense/top?limit=10"
```

### 2.3 Top Expense Categories
```
GET /analytics/expense/top-categories
```

**Query Parameters:**
- `limit` (optional): Number of top categories to return (default: 5)
- `start_date` (optional): Start date filter
- `end_date` (optional): End date filter

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/expense/top-categories?limit=5"
```

### 2.4 Expense Anomaly Detection
```
GET /analytics/expense/anomalies
```

**Query Parameters:**
- `threshold_factor` (optional): Threshold factor for anomaly detection (default: 1.5)
- `months_back` (optional): Lookback period in months (default: 3)

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/expense/anomalies?threshold_factor=2.0"
```

### 2.5 Expense Trends
```
GET /analytics/expense/trends
```

**Query Parameters:**
- `months_back` (optional): Number of months to analyze (default: 6)

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/expense/trends?months_back=12"
```

### 2.6 Share of Wallet Analysis
```
GET /analytics/expense/share-of-wallet
```

**Query Parameters:**
- `start_date` (optional): Start date for analysis
- `end_date` (optional): End date for analysis

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/expense/share-of-wallet"
```

---

## 3. Cash Flow Analytics

### 3.1 Cash Flow Summary
```
GET /analytics/cashflow/summary
```

**Query Parameters:**
- `start_date` (optional): Start date for analysis
- `end_date` (optional): End date for analysis

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/cashflow/summary"
```

### 3.2 Cash Flow Forecast
```
GET /analytics/cashflow/forecast
```

**Query Parameters:**
- `forecast_days` (optional): Number of days to forecast (default: 30)

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/cashflow/forecast?forecast_days=90"
```

### 3.3 Financial Stability Analysis
```
GET /analytics/cashflow/stability
```

**Query Parameters:**
- `months_back` (optional): Number of months to analyze (default: 6)

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/cashflow/stability"
```

### 3.4 Emergency Fund Analysis
```
GET /analytics/cashflow/emergency-fund
```

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/cashflow/emergency-fund"
```

---

## 4. Dashboard & Health

### 4.1 Comprehensive Dashboard Summary
```
GET /analytics/dashboard
```

**Query Parameters:**
- `period_start` (optional): Start date for the period
- `period_end` (optional): End date for the period

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/dashboard"
```

**Response Example:**
```json
{
  "success": true,
  "data": [
    {
      "section": "Income",
      "metric": "Total Income",
      "value": 15000.00,
      "percentage_change": 10.5,
      "trend": "Up",
      "description": "Total income for the period"
    },
    {
      "section": "Expense",
      "metric": "Total Expenses",
      "value": 10000.00,
      "percentage_change": -5.2,
      "trend": "Down",
      "description": "Total expenses for the period"
    },
    {
      "section": "CashFlow",
      "metric": "Net Cash Flow",
      "value": 5000.00,
      "trend": "Positive",
      "description": "Income minus expenses"
    }
  ]
}
```

### 4.2 Financial Health Check
```
GET /analytics/financial-health
```

**Query Parameters:**
- `start_date` (optional): Start date for analysis
- `end_date` (optional): End date for analysis
- `months_back` (optional): Number of months to analyze (default: 6)

**Example Request:**
```bash
curl -X GET "http://localhost:9595/analytics/financial-health"
```

**Response Example:**
```json
{
  "success": true,
  "data": {
    "overall_score": 75,
    "health_status": "Good",
    "metrics": [
      {
        "name": "Income to Expense Ratio",
        "value": 1.5,
        "score": 100,
        "status": "Healthy",
        "description": "Excellent financial stability - Income significantly exceeds expenses"
      },
      {
        "name": "Savings Rate (%)",
        "value": 33.33,
        "score": 100,
        "status": "Excellent",
        "description": "Excellent savings rate - Strong financial future"
      }
    ],
    "recommendations": [
      "Continue maintaining your current financial habits",
      "Consider investing surplus funds for long-term growth"
    ],
    "alerts": []
  }
}
```

---

## Error Responses

All endpoints return standard error responses:

```json
{
  "success": false,
  "error": "Error description"
}
```

Common HTTP status codes:
- `200 OK` - Successful request
- `400 Bad Request` - Invalid parameters
- `500 Internal Server Error` - Server error

---

## Database Migration

Before using the analytics endpoints, run the database migration:

```bash
psql -U your_username -d your_database -f analytics_migration.sql
```

This will create all necessary views and functions for the analytics system.

---

## Testing Examples

### Complete Income Analytics Test
```bash
# Get monthly income summary
curl -X GET "http://localhost:9595/analytics/income?analysis_type=monthly"

# Get income growth analysis
curl -X GET "http://localhost:9595/analytics/income/growth"

# Get income forecast
curl -X GET "http://localhost:9595/analytics/income/forecast"

# Get top 10 incomes
curl -X GET "http://localhost:9595/analytics/income/top?limit=10"
```

### Complete Expense Analytics Test
```bash
# Get expense by categories
curl -X GET "http://localhost:9595/analytics/expense?analysis_type=categories"

# Get expense trends
curl -X GET "http://localhost:9595/analytics/expense/trends?months_back=12"

# Detect expense anomalies
curl -X GET "http://localhost:9595/analytics/expense/anomalies"

# Get share of wallet analysis
curl -X GET "http://localhost:9595/analytics/expense/share-of-wallet"

# Get top 5 expense categories
curl -X GET "http://localhost:9595/analytics/expense/top-categories"
```

### Complete Cash Flow Test
```bash
# Get cash flow summary
curl -X GET "http://localhost:9595/analytics/cashflow/summary"

# Get 90-day forecast
curl -X GET "http://localhost:9595/analytics/cashflow/forecast?forecast_days=90"

# Check financial stability
curl -X GET "http://localhost:9595/analytics/cashflow/stability"

# Analyze emergency fund
curl -X GET "http://localhost:9595/analytics/cashflow/emergency-fund"
```

### Dashboard & Health Check
```bash
# Get complete dashboard summary
curl -X GET "http://localhost:9595/analytics/dashboard"

# Get financial health assessment
curl -X GET "http://localhost:9595/analytics/financial-health"
```