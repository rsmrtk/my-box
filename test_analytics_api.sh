#!/bin/bash

# Analytics API Testing Script
# This script tests all analytics endpoints with various parameters

# Configuration
BASE_URL="http://localhost:9595"
CONTENT_TYPE="Content-Type: application/json"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}Analytics API Testing Script${NC}"
echo -e "${BLUE}================================${NC}\n"

# Function to test endpoint
test_endpoint() {
    local name="$1"
    local endpoint="$2"

    echo -e "${YELLOW}Testing: $name${NC}"
    echo -e "Endpoint: ${GREEN}$endpoint${NC}"
    echo "----------------------------------------"

    response=$(curl -s -X GET "$BASE_URL$endpoint" -H "$CONTENT_TYPE")

    # Pretty print JSON response
    echo "$response" | python3 -m json.tool 2>/dev/null || echo "$response"

    echo -e "\n========================================\n"
    sleep 1  # Small delay between requests
}

# Test Dashboard
test_endpoint "Dashboard Summary" "/analytics/dashboard"

# Test Top Expenses
test_endpoint "Top 3 Expense Categories (Default)" "/analytics/expenses/top"
test_endpoint "Top 5 Expense Categories" "/analytics/expenses/top?limit=5"
test_endpoint "Top 10 Expense Categories" "/analytics/expenses/top?limit=10"

# Test Expense Trends
test_endpoint "Expense Trends - Last 6 Months (Default)" "/analytics/expenses/trends"
test_endpoint "Expense Trends - Last 3 Months" "/analytics/expenses/trends?months=3"
test_endpoint "Expense Trends - Last 12 Months" "/analytics/expenses/trends?months=12"

# Test Income Growth
test_endpoint "Income Growth Analysis" "/analytics/income/growth"

# Test Anomaly Detection
test_endpoint "Expense Anomalies - Default Threshold (1.5)" "/analytics/expenses/anomalies"
test_endpoint "Expense Anomalies - Sensitive Threshold (1.2)" "/analytics/expenses/anomalies?threshold=1.2"
test_endpoint "Expense Anomalies - Strict Threshold (2.0)" "/analytics/expenses/anomalies?threshold=2.0"
test_endpoint "Expense Anomalies - Very Strict Threshold (3.0)" "/analytics/expenses/anomalies?threshold=3.0"

echo -e "${GREEN}✓ All tests completed!${NC}\n"

# Optional: Save all responses to files
read -p "Do you want to save all responses to files? (y/n): " save_responses

if [[ $save_responses == "y" || $save_responses == "Y" ]]; then
    mkdir -p analytics_test_results

    echo -e "\n${BLUE}Saving responses to analytics_test_results/ directory...${NC}"

    # Save each endpoint response
    curl -s "$BASE_URL/analytics/dashboard" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/dashboard.json
    curl -s "$BASE_URL/analytics/expenses/top" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/top_expenses_3.json
    curl -s "$BASE_URL/analytics/expenses/top?limit=5" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/top_expenses_5.json
    curl -s "$BASE_URL/analytics/expenses/trends" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/trends_6m.json
    curl -s "$BASE_URL/analytics/expenses/trends?months=3" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/trends_3m.json
    curl -s "$BASE_URL/analytics/income/growth" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/income_growth.json
    curl -s "$BASE_URL/analytics/expenses/anomalies" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/anomalies_default.json
    curl -s "$BASE_URL/analytics/expenses/anomalies?threshold=2.0" -H "$CONTENT_TYPE" | python3 -m json.tool > analytics_test_results/anomalies_strict.json

    echo -e "${GREEN}✓ Responses saved to analytics_test_results/ directory${NC}"
    ls -la analytics_test_results/
fi

echo -e "\n${BLUE}Testing complete!${NC}"