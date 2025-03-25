package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hokamsingh/go-backend-template/internal/repository"
)

type EventController struct {
	repo *repository.EventRepository
}

func NewEventController(repo *repository.EventRepository) *EventController {
	return &EventController{repo}
}

func (ec *EventController) GetAllEvents(c *gin.Context) {
	events, err := ec.repo.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (ec *EventController) GetEventByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	event, err := ec.repo.GetByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}
