package handler

import (
	"net/http"

	"github.com/XiaoleC05/AIHelper/internal/db"
	"github.com/XiaoleC05/AIHelper/internal/model"
	"github.com/gin-gonic/gin"
)

type SettingsHandler struct {
	repo *db.SettingsRepository
}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{repo: db.NewSettingsRepository()}
}

func (h *SettingsHandler) Get(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	settings, err := h.repo.GetByUser(c.Request.Context(), userID)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	if settings == nil {
		settings = &model.UserSettings{
			UserID: userID,
			APIKey: "",
			APIBase: "",
			Model:  "gpt-4o-mini",
		}
	}

	c.JSON(http.StatusOK, gin.H{"settings": settings})
}

func (h *SettingsHandler) Update(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	var req model.UpdateSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := h.repo.Upsert(c.Request.Context(), userID, req)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"settings": settings})
}
