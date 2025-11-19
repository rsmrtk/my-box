# Analytics Module - Quick Start Guide

## Overview
The Analytics module provides comprehensive financial insights through 5 main endpoints:

1. **Dashboard** - Overall financial health snapshot
2. **Top Expenses** - Highest spending categories
3. **Expense Trends** - Monthly spending patterns
4. **Income Growth** - Month-over-month income analysis
5. **Anomaly Detection** - Unusual expense identification

## Quick Test Setup

### 1. Load Test Data
```bash
# Connect to your PostgreSQL database and run:
psql -U your_user -d your_database -f analytics_test_data.sql
```

### 2. Start the Server
```bash
go run cmd/main.go
# Server starts on http://localhost:9595
```

### 3. Test the Endpoints

#### Option A: Use the Test Script (Recommended)
```bash
# Make sure the script is executable
chmod +x test_analytics_api.sh

# Run all tests
./test_analytics_api.sh
```

#### Option B: Manual Testing with curl

```bash
# Dashboard
curl http://localhost:9595/analytics/dashboard | jq

# Top 3 Expense Categories
curl http://localhost:9595/analytics/expenses/top | jq

# Expense Trends (6 months)
curl http://localhost:9595/analytics/expenses/trends | jq

# Income Growth
curl http://localhost:9595/analytics/income/growth | jq

# Anomaly Detection
curl http://localhost:9595/analytics/expenses/anomalies | jq
```

#### Option C: Import Postman Collection
1. Open Postman
2. Import `Analytics_API_Postman.json`
3. Set the `baseUrl` variable to `http://localhost:9595`
4. Run the collection

## API Endpoints Summary

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/analytics/dashboard` | GET | Complete financial overview |
| `/analytics/expenses/top?limit=3` | GET | Top expense categories |
| `/analytics/expenses/trends?months=6` | GET | Monthly expense trends |
| `/analytics/income/growth` | GET | Income growth analysis |
| `/analytics/expenses/anomalies?threshold=1.5` | GET | Detect unusual expenses |

## Key Features

### Dashboard Metrics
- **Income**: Total, monthly, daily average
- **Expenses**: Total, monthly, daily average
- **Cash Flow**: Net amount, savings rate
- **Stability**: Ratio and status (Excellent/Good/Stable/Warning/Critical)
- **Top Categories**: Top 3 spending categories with percentages

### Anomaly Detection
Identifies expenses that deviate significantly from category averages:
- **Critical**: >3x average
- **High**: >2x average
- **Medium**: >threshold (default 1.5x)

### Trend Analysis
- Month-over-month expense changes
- Percentage changes between periods
- Transaction counts per month

## Response Format
All endpoints return JSON with consistent structure:
```json
{
  "success": true,
  "data": {
    // Endpoint-specific data
  }
}
```

Amounts use array format for multi-currency support:
```json
"amount": [
  {
    "amount": 1000.00,
    "currency_code": "USD",
    "currency_symbol": "$"
  }
]
```

## Troubleshooting

### No Data Returned
- Check if test data was loaded: `SELECT COUNT(*) FROM expense;`
- Verify date ranges in queries match your data

### Server Not Responding
- Ensure server is running on port 9595
- Check logs for database connection issues

### Anomaly Detection Not Working
- Need at least 3 months of historical data
- Expenses must have `expense_type` field populated

## Files Overview

| File | Purpose |
|------|---------|
| `ANALYTICS_API_DOCS.md` | Complete API documentation |
| `analytics_test_data.sql` | Sample data for testing |
| `test_analytics_api.sh` | Automated testing script |
| `Analytics_API_Postman.json` | Postman collection |
| `ANALYTICS_README.md` | This quick start guide |

## Clean Up Test Data
```sql
-- Remove test data when done
DELETE FROM expense WHERE expense_id LIKE 'test_%';
DELETE FROM income WHERE income_id LIKE 'test_%';
```

## Production Considerations
1. Add authentication middleware
2. Implement caching for dashboard endpoint
3. Add rate limiting
4. Create database indexes on date columns
5. Consider pagination for large datasets
6. Add logging and monitoring

## Support
For issues or questions about the analytics module, check:
- API Documentation: `ANALYTICS_API_DOCS.md`
- Service implementations: `/internal/rest/services/analytics/`
- Domain models: `/internal/rest/domain/analytics/models.go`