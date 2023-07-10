package transactionsdto

type TransactionResponse struct {
	UserID   int    `json:"user_id"`
	TicketID int    `json:"ticket_id"`
	Image    string `json:"image"`
}
