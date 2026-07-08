package handler

import (
	"net/http"

	"github.com/XiaoleC05/AIHelper/internal/db"
	"github.com/gin-gonic/gin"
)

type TemplateHandler struct {
	repo *db.TemplateRepository
}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{repo: db.NewTemplateRepository()}
}

func (h *TemplateHandler) List(c *gin.Context) {
	category := c.Query("category")

	templates, err := h.repo.List(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"templates": templates})
}
