# Analytics API Documentation

## Base URL
```
http://localhost:9595
```

## Authentication
All endpoints require appropriate authentication headers if configured.

---

## 1. Dashboard Summary
Get comprehensive financial dashboard metrics including income, expenses, cash flow, and top categories.

### Endpoint
```
GET /analytics/dashboard
```

### Request
```bash
curl -X GET http://localhost:9595/analytics/dashboard \
  -H "Content-Type: application/json"
```

### Response Example
```json
{
  "success": true,
  "data": {
    "total_income": [
      {
        "amount": 125000.50,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "monthly_income": [
      {
        "amount": 10500.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "daily_avg_income": [
      {
        "amount": 350.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "total_expense": [
      {
        "amount": 95000.25,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "monthly_expense": [
      {
        "amount": 7500.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "daily_avg_expense": [
      {
        "amount": 250.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "net_cash_flow": [
      {
        "amount": 30000.25,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "savings_rate": 24.0,
    "stability_ratio": 1.32,
    "stability_status": "Good",
    "top_expense_categories": [
      {
        "category": "Housing",
        "total": [
          {
            "amount": 35000.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 12,
        "percentage": 36.84
      },
      {
        "category": "Food",
        "total": [
          {
            "amount": 25000.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 150,
        "percentage": 26.32
      },
      {
        "category": "Transportation",
        "total": [
          {
            "amount": 15000.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 45,
        "percentage": 15.79
      }
    ]
  }
}
```

### Metrics Explained
- **stability_status**:
  - "Excellent" - ratio >= 1.5
  - "Good" - ratio >= 1.2
  - "Stable" - ratio >= 1.0
  - "Warning" - ratio >= 0.9
  - "Critical" - ratio < 0.9

---

## 2. Top Expense Categories
Get the top N expense categories by total amount.

### Endpoint
```
GET /analytics/expenses/top?limit=3
```

### Query Parameters
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| limit | int | 3 | Number of top categories to return |

### Request Examples
```bash
# Get top 3 expense categories (default)
curl -X GET http://localhost:9595/analytics/expenses/top \
  -H "Content-Type: application/json"

# Get top 5 expense categories
curl -X GET http://localhost:9595/analytics/expenses/top?limit=5 \
  -H "Content-Type: application/json"

# Get top 10 expense categories
curl -X GET http://localhost:9595/analytics/expenses/top?limit=10 \
  -H "Content-Type: application/json"
```

### Response Example
```json
{
  "success": true,
  "data": {
    "categories": [
      {
        "category": "Housing",
        "total": [
          {
            "amount": 35000.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 12,
        "percentage": 43.75
      },
      {
        "category": "Food",
        "total": [
          {
            "amount": 25000.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 150,
        "percentage": 31.25
      },
      {
        "category": "Transportation",
        "total": [
          {
            "amount": 20000.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 45,
        "percentage": 25.00
      }
    ]
  }
}
```

---

## 3. Expense Trends
Analyze expense trends over the specified number of months.

### Endpoint
```
GET /analytics/expenses/trends?months=6
```

### Query Parameters
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| months | int | 6 | Number of months to analyze |

### Request Examples
```bash
# Get expense trends for last 6 months (default)
curl -X GET http://localhost:9595/analytics/expenses/trends \
  -H "Content-Type: application/json"

# Get expense trends for last 3 months
curl -X GET http://localhost:9595/analytics/expenses/trends?months=3 \
  -H "Content-Type: application/json"

# Get expense trends for last 12 months
curl -X GET http://localhost:9595/analytics/expenses/trends?months=12 \
  -H "Content-Type: application/json"
```

### Response Example
```json
{
  "success": true,
  "data": {
    "trends": [
      {
        "month": "2024-06-01",
        "total": [
          {
            "amount": 7200.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 45,
        "change": null,
        "change_percentage": 0
      },
      {
        "month": "2024-07-01",
        "total": [
          {
            "amount": 7500.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 52,
        "change": [
          {
            "amount": 300.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "change_percentage": 4.17
      },
      {
        "month": "2024-08-01",
        "total": [
          {
            "amount": 6800.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 48,
        "change": [
          {
            "amount": -700.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "change_percentage": -9.33
      },
      {
        "month": "2024-09-01",
        "total": [
          {
            "amount": 8200.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 55,
        "change": [
          {
            "amount": 1400.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "change_percentage": 20.59
      },
      {
        "month": "2024-10-01",
        "total": [
          {
            "amount": 7900.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 50,
        "change": [
          {
            "amount": -300.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "change_percentage": -3.66
      },
      {
        "month": "2024-11-01",
        "total": [
          {
            "amount": 7500.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "count": 49,
        "change": [
          {
            "amount": -400.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "change_percentage": -5.06
      }
    ]
  }
}
```

---

## 4. Income Growth Analysis
Analyze income growth between current and previous month.

### Endpoint
```
GET /analytics/income/growth
```

### Request
```bash
curl -X GET http://localhost:9595/analytics/income/growth \
  -H "Content-Type: application/json"
```

### Response Example
```json
{
  "success": true,
  "data": {
    "current_month": [
      {
        "amount": 10500.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "previous_month": [
      {
        "amount": 9800.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "growth_amount": [
      {
        "amount": 700.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "growth_percentage": 7.14
  }
}
```

### Growth Scenarios
- **Positive Growth**: Current > Previous (percentage > 0)
- **Negative Growth**: Current < Previous (percentage < 0)
- **New Income**: Previous = 0, Current > 0 (percentage = 100)
- **No Income**: Both = 0 (percentage = 0)

---

## 5. Expense Anomaly Detection
Detect expense anomalies based on deviation from category averages.

### Endpoint
```
GET /analytics/expenses/anomalies?threshold=1.5
```

### Query Parameters
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| threshold | float | 1.5 | Deviation factor threshold for anomaly detection |

### Request Examples
```bash
# Detect anomalies with default threshold (1.5x average)
curl -X GET http://localhost:9595/analytics/expenses/anomalies \
  -H "Content-Type: application/json"

# Detect anomalies with 2x average threshold
curl -X GET http://localhost:9595/analytics/expenses/anomalies?threshold=2.0 \
  -H "Content-Type: application/json"

# Detect anomalies with 1.2x average threshold (more sensitive)
curl -X GET http://localhost:9595/analytics/expenses/anomalies?threshold=1.2 \
  -H "Content-Type: application/json"
```

### Response Example
```json
{
  "success": true,
  "data": {
    "anomalies": [
      {
        "id": "e1234567-89ab-cdef-0123-456789abcdef",
        "name": "Annual Insurance Premium",
        "amount": [
          {
            "amount": 3600.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "type": "Insurance",
        "date": "2024-11-15",
        "category_average": [
          {
            "amount": 300.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "deviation_factor": 12.0,
        "status": "Critical"
      },
      {
        "id": "f2345678-9abc-def0-1234-56789abcdef0",
        "name": "Emergency Car Repair",
        "amount": [
          {
            "amount": 2500.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "type": "Transportation",
        "date": "2024-11-10",
        "category_average": [
          {
            "amount": 500.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "deviation_factor": 5.0,
        "status": "Critical"
      },
      {
        "id": "a3456789-abcd-ef01-2345-6789abcdef01",
        "name": "Holiday Shopping",
        "amount": [
          {
            "amount": 1200.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "type": "Shopping",
        "date": "2024-11-05",
        "category_average": [
          {
            "amount": 400.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "deviation_factor": 3.0,
        "status": "Critical"
      },
      {
        "id": "b4567890-bcde-f012-3456-789abcdef012",
        "name": "Premium Restaurant Dinner",
        "amount": [
          {
            "amount": 450.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "type": "Food",
        "date": "2024-11-01",
        "category_average": [
          {
            "amount": 200.00,
            "currency_code": "USD",
            "currency_symbol": "$"
          }
        ],
        "deviation_factor": 2.25,
        "status": "High"
      }
    ]
  }
}
```

### Anomaly Status Levels
- **Critical**: deviation_factor > 3.0 (expense is 3x higher than average)
- **High**: deviation_factor > 2.0 (expense is 2x higher than average)
- **Medium**: deviation_factor > threshold (expense exceeds threshold)

---

## Error Responses

All endpoints return consistent error responses:

### 500 Internal Server Error
```json
{
  "success": false,
  "error": "Failed to calculate income metrics."
}
```

### Common Error Messages
- "Failed to calculate income metrics."
- "Failed to calculate expense metrics."
- "Failed to get expense categories."
- "Failed to get expense trends."
- "Failed to get top expense categories."
- "Failed to calculate income growth."
- "Failed to find expense anomalies."

---

## Testing with Sample Data

### 1. Insert Sample Income Data
```sql
-- Insert sample income data for testing
INSERT INTO income (income_id, income_name, income_amount, income_type, income_date, created_at) VALUES
('i001', 'Salary', 8500.00, 'Employment', '2024-11-01', NOW()),
('i002', 'Freelance Project', 2000.00, 'Freelance', '2024-11-05', NOW()),
('i003', 'Investment Dividend', 500.00, 'Investment', '2024-11-10', NOW()),
('i004', 'Salary', 8500.00, 'Employment', '2024-10-01', NOW()),
('i005', 'Side Business', 1300.00, 'Business', '2024-10-15', NOW());
```

### 2. Insert Sample Expense Data
```sql
-- Insert sample expense data for testing
INSERT INTO expense (expense_id, expense_name, expense_amount, expense_type, expense_date, created_at) VALUES
('e001', 'Rent', 2500.00, 'Housing', '2024-11-01', NOW()),
('e002', 'Groceries', 450.00, 'Food', '2024-11-02', NOW()),
('e003', 'Gas', 120.00, 'Transportation', '2024-11-03', NOW()),
('e004', 'Electric Bill', 180.00, 'Utilities', '2024-11-05', NOW()),
('e005', 'Restaurant', 85.00, 'Food', '2024-11-06', NOW()),
('e006', 'Internet', 70.00, 'Utilities', '2024-11-07', NOW()),
('e007', 'Car Insurance', 200.00, 'Insurance', '2024-11-08', NOW()),
('e008', 'Groceries', 380.00, 'Food', '2024-11-09', NOW()),
('e009', 'Coffee Shop', 25.00, 'Food', '2024-11-10', NOW()),
('e010', 'Emergency Car Repair', 2500.00, 'Transportation', '2024-11-10', NOW()), -- Anomaly
('e011', 'Annual Insurance Premium', 3600.00, 'Insurance', '2024-11-15', NOW()), -- Anomaly
('e012', 'Holiday Shopping', 1200.00, 'Shopping', '2024-11-05', NOW()), -- Anomaly
('e013', 'Rent', 2500.00, 'Housing', '2024-10-01', NOW()),
('e014', 'Groceries', 420.00, 'Food', '2024-10-02', NOW()),
('e015', 'Gas', 100.00, 'Transportation', '2024-10-03', NOW());
```

### 3. Test All Endpoints
```bash
# Test Dashboard
curl -X GET http://localhost:9595/analytics/dashboard

# Test Top Expenses (default top 3)
curl -X GET http://localhost:9595/analytics/expenses/top

# Test Top 5 Expenses
curl -X GET http://localhost:9595/analytics/expenses/top?limit=5

# Test Expense Trends (last 6 months)
curl -X GET http://localhost:9595/analytics/expenses/trends

# Test Expense Trends (last 3 months)
curl -X GET http://localhost:9595/analytics/expenses/trends?months=3

# Test Income Growth
curl -X GET http://localhost:9595/analytics/income/growth

# Test Anomaly Detection (default threshold)
curl -X GET http://localhost:9595/analytics/expenses/anomalies

# Test Anomaly Detection (stricter threshold)
curl -X GET http://localhost:9595/analytics/expenses/anomalies?threshold=2.0
```

---

## Performance Considerations

1. **Dashboard Endpoint**: Executes multiple queries, consider caching for production use
2. **Anomaly Detection**: Scans last month's expenses and calculates 3-month averages
3. **Trends Analysis**: Aggregates data by month, performance depends on data volume
4. **All queries use indexes**: Ensure proper indexes on expense_date, income_date, expense_type

---

## Notes for Development

1. All amounts are returned in Amount array format for multi-currency support
2. Dates are returned in YYYY-MM-DD format
3. Percentages are returned as floating-point numbers (e.g., 25.5 means 25.5%)
4. All endpoints use GET method for data retrieval
5. Consider implementing caching for frequently accessed data
6. Add authentication middleware before deploying to production