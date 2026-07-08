package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/XiaoleC05/AIHelper/internal/db"
	"github.com/XiaoleC05/AIHelper/internal/model"
	"github.com/gin-gonic/gin"
)

const enhanceSystemPrompt = `你是一个专业的提示词工程师。你的任务是将用户的模糊需求转化为结构化的、高质量的提示词。

优化原则：
1. 明确角色和身份
2. 提供清晰的上下文
3. 指定输出格式
4. 添加约束条件
5. 包含示例（如适用）

请输出优化后的提示词，保持简洁专业。`

type EnhanceHandler struct {
	settingsRepo *db.SettingsRepository
}

func NewEnhanceHandler() *EnhanceHandler {
	return &EnhanceHandler{settingsRepo: db.NewSettingsRepository()}
}

func (h *EnhanceHandler) Enhance(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	var req model.EnhanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := h.settingsRepo.GetByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if settings == nil || settings.APIKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "API key not configured"})
		return
	}

	apiBase := settings.APIBase
	if apiBase == "" {
		apiBase = "https://api.openai.com/v1"
	}

	enhanced, err := callLLM(apiBase, settings.APIKey, settings.Model, req.Content)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("LLM call failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, model.EnhanceResponse{
		Original:  req.Content,
		Enhanced: enhanced,
	})
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func callLLM(apiBase, apiKey, modelName, userContent string) (string, error) {
	reqBody := chatRequest{
		Model: modelName,
		Messages: []chatMessage{
			{Role: "system", Content: enhanceSystemPrompt},
			{Role: "user", Content: "请优化以下提示词：\n\n" + userContent},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	client := &http.Client{Timeout: 60 * time.Second}
	httpReq, err := http.NewRequest("POST", apiBase+"/chat/completions", bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var chatResp chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", err
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response from LLM")
	}

	return chatResp.Choices[0].Message.Content, nil
}
