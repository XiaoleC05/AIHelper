package model

import "time"

type Prompt struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	Variables []string  `json:"variables"`
	IsFavorite bool    `json:"is_favorite"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePromptRequest struct {
	Title     string   `json:"title" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	Variables []string `json:"variables"`
}

type UpdatePromptRequest struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	Variables []string `json:"variables"`
}

type Template struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type UserSettings struct {
	UserID    int64     `json:"user_id"`
	APIKey    string    `json:"api_key"`
	APIBase   string    `json:"api_base"`
	Model     string    `json:"model"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateSettingsRequest struct {
	APIKey  *string `json:"api_key"`
	APIBase string  `json:"api_base"`
	Model   string  `json:"model"`
}

type EnhanceRequest struct {
	Content string `json:"content" binding:"required"`
}

type EnhanceResponse struct {
	Original  string `json:"original"`
	Enhanced string `json:"enhanced"`
}
