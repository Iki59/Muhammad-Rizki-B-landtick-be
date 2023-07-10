package handlers

import (
	dto "landtick-be/dto/result"
	ticketsdto "landtick-be/dto/tickets"
	"landtick-be/models"
	"landtick-be/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerticket struct {
	TicketRepository repository.TicketRepository
}

func HandlerTicket(TicketRepository repository.TicketRepository) *handlerticket {
	return &handlerticket{TicketRepository}
}

func (h *handlerticket) FindTickets(c echo.Context) error {
	tickets, err := h.TicketRepository.FindTickets()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: tickets})
}

func (h *handlerticket) GetTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ticket, err := h.TicketRepository.GetTicket(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: ticketResponse(ticket)})
}

func (h *handlerticket) CreateTicket(c echo.Context) error {
	request := new(ticketsdto.CreateTicketRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Status:  "Failed",
			Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Status:  "Failed",
			Message: err.Error()})
	}

	data := models.Ticket{
		NameTrain:            request.NameTrain,
		TypeTrain:            request.TypeTrain,
		StartDate:            request.StartDate,
		StartStationID:       request.StartStationID,
		StartTime:            request.StartTime,
		DestinationStationID: request.DestinationStationID,
		ArivalTime:           request.ArivalTime,
		Price:                request.Price,
		Quantity:             request.Quantity,
	}

	response, err := h.TicketRepository.CreateTicket(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Status: "Success",
		Data:   ticketResponse(response)})
}

func ticketResponse(u models.Ticket) ticketsdto.TicketResponse {
	return ticketsdto.TicketResponse{
		NameTrain:            u.NameTrain,
		TypeTrain:            u.TypeTrain,
		StartDate:            u.StartDate,
		StartStationID:       u.StartStationID,
		StartTime:            u.StartTime,
		DestinationStationID: u.DestinationStationID,
		ArivalTime:           u.ArivalTime,
		Price:                u.Price,
		Quantity:             u.Quantity,
	}
}
