# API Test Requests

## Server Information
- **Base URL**: `http://localhost:9595`
- **Date Format**: `YYYY-MM-DD` (example: `2025-11-19`)

---

## EXPENSE ENDPOINTS

### 1. Create Expense
**Endpoint**: `POST /expense`

```bash
curl -X POST http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d '{
    "expense_name": "Grocery Shopping",
    "expense_amount": [
      {
        "amount": 125.50,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "expense_type": "Food",
    "expense_date": "2025-11-19"
  }'
```

### 2. Get Single Expense
**Endpoint**: `GET /expense`

```bash
curl -X GET http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d '{
    "expense_id": "YOUR_EXPENSE_ID_HERE"
  }'
```

### 3. List All Expenses
**Endpoint**: `GET /expense/list`

```bash
curl -X GET http://localhost:9595/expense/list \
  -H "Content-Type: application/json" \
  -d '{
    "limit": 10,
    "offset": 0,
    "sort_by": "expense_date",
    "order": "desc"
  }'
```

Or without body (will use default parameters):
```bash
curl -X GET http://localhost:9595/expense/list
```

### 4. Update Expense
**Endpoint**: `PUT /expense`

```bash
curl -X PUT http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d '{
    "expense_id": "YOUR_EXPENSE_ID_HERE",
    "expense_name": "Updated Grocery Shopping",
    "expense_amount": [
      {
        "amount": 150.75,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "expense_type": "Food & Groceries",
    "expense_date": "2025-11-20"
  }'
```

### 5. Delete Expense
**Endpoint**: `DELETE /expense`

```bash
curl -X DELETE http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d '{
    "expense_id": "YOUR_EXPENSE_ID_HERE"
  }'
```

---

## INCOME ENDPOINTS

### 1. Create Income
**Endpoint**: `POST /income`

```bash
curl -X POST http://localhost:9595/income \
  -H "Content-Type: application/json" \
  -d '{
    "income_name": "Monthly Salary",
    "income_amount": [
      {
        "amount": 5000.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "income_type": "Salary",
    "income_date": "2025-11-15"
  }'
```

### 2. Get Single Income
**Endpoint**: `GET /income`

```bash
curl -X GET http://localhost:9595/income \
  -H "Content-Type: application/json" \
  -d '{
    "income_id": "YOUR_INCOME_ID_HERE"
  }'
```

### 3. List All Incomes
**Endpoint**: `GET /income/list`

```bash
curl -X GET http://localhost:9595/income/list \
  -H "Content-Type: application/json" \
  -d '{
    "limit": 10,
    "offset": 0,
    "sort_by": "income_date",
    "order": "desc"
  }'
```

Or without body (will use default parameters):
```bash
curl -X GET http://localhost:9595/income/list
```

### 4. Update Income
**Endpoint**: `PUT /income`

```bash
curl -X PUT http://localhost:9595/income \
  -H "Content-Type: application/json" \
  -d '{
    "income_id": "YOUR_INCOME_ID_HERE",
    "income_name": "Updated Monthly Salary",
    "income_amount": [
      {
        "amount": 5500.00,
        "currency_code": "USD",
        "currency_symbol": "$"
      }
    ],
    "income_type": "Salary Increase",
    "income_date": "2025-11-16"
  }'
```

### 5. Delete Income
**Endpoint**: `DELETE /income`

```bash
curl -X DELETE http://localhost:9595/income \
  -H "Content-Type: application/json" \
  -d '{
    "income_id": "YOUR_INCOME_ID_HERE"
  }'
```

---

## TEST SCENARIOS

### Testing Date Format Support

#### ✅ Valid Date Formats (Should Work):
```json
{
  "expense_date": "2025-11-19",
  "income_date": "2025-01-15"
}
```

#### ❌ Old Format (Will Now Be Rejected):
```json
{
  "expense_date": "2025-11-19T00:00:00Z",
  "income_date": "2025-01-15T14:30:00+02:00"
}
```

### Complete Test Flow Example:

```bash
# 1. Create an expense
EXPENSE_RESPONSE=$(curl -s -X POST http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d '{
    "expense_name": "Test Expense",
    "expense_amount": [{"amount": 100.00, "currency_code": "USD", "currency_symbol": "$"}],
    "expense_type": "Test",
    "expense_date": "2025-11-19"
  }')

echo "Created Expense: $EXPENSE_RESPONSE"

# 2. Extract expense_id from response (using jq if installed)
EXPENSE_ID=$(echo $EXPENSE_RESPONSE | jq -r '.expense_id')

# 3. Get the expense
curl -X GET http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d "{\"expense_id\": \"$EXPENSE_ID\"}"

# 4. Update the expense
curl -X PUT http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d "{
    \"expense_id\": \"$EXPENSE_ID\",
    \"expense_name\": \"Updated Test Expense\",
    \"expense_date\": \"2025-11-20\"
  }"

# 5. List all expenses
curl -X GET http://localhost:9595/expense/list

# 6. Delete the expense
curl -X DELETE http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d "{\"expense_id\": \"$EXPENSE_ID\"}"
```

---

## NOTES

1. **Date Format**: All date fields now accept `YYYY-MM-DD` format (e.g., `2025-11-19`)
2. **Response Dates**: All date fields in responses will also be in `YYYY-MM-DD` format
3. **IDs**: Replace `YOUR_EXPENSE_ID_HERE` and `YOUR_INCOME_ID_HERE` with actual UUIDs returned from create operations
4. **Amount Structure**: The `amount` field is an array to support multiple currencies in the future
5. **Optional Fields**: In update requests, you can omit fields you don't want to change (except for the ID which is required)

---

## Quick Test Commands

### Health Check
```bash
curl http://localhost:9595/health
# Expected: {"status":"ok"}
```

### Test Date Parsing (Create with simple date)
```bash
curl -X POST http://localhost:9595/expense \
  -H "Content-Type: application/json" \
  -d '{
    "expense_name": "Date Test",
    "expense_amount": [{"amount": 50, "currency_code": "USD", "currency_symbol": "$"}],
    "expense_type": "Test",
    "expense_date": "2025-11-19"
  }'
```

If the response is successful and returns the expense with dates in `YYYY-MM-DD` format, the fix is working correctly!