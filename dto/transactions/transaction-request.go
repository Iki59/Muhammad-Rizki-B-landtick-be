package transactionsdto

type CreateTransactionRequest struct {
	UserID   int    `json:"user_id" form:"user_id" validation:"required"`
	TicketID int    `json:"ticket_id" form:"ticket_id" validation:"required"`
	Image    string `json:"image" form:"image"`
}
