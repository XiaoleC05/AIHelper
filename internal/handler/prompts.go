package handler

import (
	"net/http"
	"strconv"

	"github.com/XiaoleC05/AIHelper/internal/db"
	"github.com/XiaoleC05/AIHelper/internal/model"
	"github.com/gin-gonic/gin"
)

type PromptHandler struct {
	repo *db.PromptRepository
}

func NewPromptHandler() *PromptHandler {
	return &PromptHandler{repo: db.NewPromptRepository()}
}

func (h *PromptHandler) List(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	search := c.Query("q")
	category := c.Query("category")
	tag := c.Query("tag")

	var favorite *bool
	if favStr := c.Query("favorite"); favStr != "" {
		b, err := strconv.ParseBool(favStr)
		if err == nil {
			favorite = &b
		}
	}

	prompts, err := h.repo.List(c.Request.Context(), userID, search, category, tag, favorite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if prompts == nil {
		prompts = []*model.Prompt{}
	}

	c.JSON(http.StatusOK, gin.H{"prompts": prompts})
}

func (h *PromptHandler) Create(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	var req model.CreatePromptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prompt, err := h.repo.Create(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"prompt": prompt})
}

func (h *PromptHandler) Get(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	prompt, err := h.repo.GetByID(c.Request.Context(), id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prompt == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "prompt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prompt": prompt})
}

func (h *PromptHandler) Update(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req model.UpdatePromptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prompt, err := h.repo.Update(c.Request.Context(), id, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prompt == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "prompt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prompt": prompt})
}

func (h *PromptHandler) Delete(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.Delete(c.Request.Context(), id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *PromptHandler) ToggleFavorite(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	prompt, err := h.repo.ToggleFavorite(c.Request.Context(), id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prompt == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "prompt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prompt": prompt})
}
