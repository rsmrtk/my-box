package income

// DeleteRequest represents the request structure for deleting an income
type DeleteRequest struct {
	IncomeID string `json:"income_id" binding:"required"`
}

// DeleteResponse represents the response structure for deleting an income
type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
