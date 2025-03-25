package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hokamsingh/go-backend-template/internal/repository"
)

type SpeakerController struct {
	repo *repository.SpeakerRepository
}

func NewSpeakerController(repo *repository.SpeakerRepository) *SpeakerController {
	return &SpeakerController{repo}
}

func (sc *SpeakerController) GetAllSpeakers(c *gin.Context) {
	// Проверяем, есть ли параметр id в запросе
	if idStr := c.Query("id"); idStr != "" {
		// Если id указан, вызываем GetSpeakerByID
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		speaker, err := sc.repo.GetByID(c, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Speaker not found"})
			return
		}
		c.JSON(http.StatusOK, speaker)
		return
	}

	// Если id не указан, возвращаем всех спикеров
	speakers, err := sc.repo.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, speakers)
}

func (sc *SpeakerController) GetSpeakerByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	speaker, err := sc.repo.GetByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Speaker not found"})
		return
	}
	c.JSON(http.StatusOK, speaker)
}
