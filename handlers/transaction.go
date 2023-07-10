package handlers

import (
	dto "landtick-be/dto/result"
	transactionsdto "landtick-be/dto/transactions"
	"landtick-be/models"
	"landtick-be/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlertransaction struct {
	TransactionRepository repository.TransactionRepository
}

func HandlerTransaction(TransactionRepository repository.TransactionRepository) *handlertransaction {
	return &handlertransaction{TransactionRepository}
}

func (h *handlertransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	for i, p := range transactions {
		transactions[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: transactions})
}

func (h *handlertransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: transactionResponse(transaction)})
}

func (h *handlertransaction) CreateTransaction(c echo.Context) error {
	request := new(transactionsdto.CreateTransactionRequest)

	UserID, _ := strconv.Atoi(c.FormValue("user_id"))
	TicketID, _ := strconv.Atoi(c.FormValue("ticket_id"))
	Image := c.Get("dataFile").(string)
	request.UserID = UserID
	request.TicketID = TicketID
	request.Image = Image

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Status:  "Failed",
			Message: err.Error()})
	}

	data := models.Transaction{
		UserID:   request.UserID,
		TicketID: request.TicketID,
		Image:    request.Image,
	}

	response, err := h.TransactionRepository.CreateTransaction(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Status: "Success",
		Data:   transactionResponse(response)})
}

func transactionResponse(u models.Transaction) transactionsdto.TransactionResponse {
	return transactionsdto.TransactionResponse{
		UserID:   u.UserID,
		TicketID: u.TicketID,
	}
}
