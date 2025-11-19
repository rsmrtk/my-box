package expense

// DeleteRequest represents the request structure for deleting an expense
type DeleteRequest struct {
	ExpenseID string `json:"expense_id" binding:"required"`
}

// DeleteResponse represents the response structure for deleting an expense
type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
