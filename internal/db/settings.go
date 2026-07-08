package db

import (
	"context"

	"github.com/XiaoleC05/AIHelper/internal/model"
	"github.com/jackc/pgx/v5"
)

type SettingsRepository struct{}

func NewSettingsRepository() *SettingsRepository {
	return &SettingsRepository{}
}

func (r *SettingsRepository) GetByUser(ctx context.Context, userID int64) (*model.UserSettings, error) {
	query := `
		SELECT user_id, api_key, api_base, model, updated_at
		FROM aihelper.user_settings
		WHERE user_id = $1
	`

	var s model.UserSettings
	err := Pool.QueryRow(ctx, query, userID).Scan(&s.UserID, &s.APIKey, &s.APIBase, &s.Model, &s.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &s, err
}

func (r *SettingsRepository) Upsert(ctx context.Context, userID int64, req model.UpdateSettingsRequest) (*model.UserSettings, error) {
	query := `
		INSERT INTO aihelper.user_settings (user_id, api_key, api_base, model)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id) DO UPDATE
		SET api_key = COALESCE(NULLIF($2, ''), aihelper.user_settings.api_key),
		    api_base = COALESCE(NULLIF($3, ''), aihelper.user_settings.api_base),
		    model = COALESCE(NULLIF($4, ''), aihelper.user_settings.model),
		    updated_at = NOW()
		RETURNING user_id, api_key, api_base, model, updated_at
	`

	apiKey := req.APIKey
	apiBase := req.APIBase
	modelName := req.Model
	if modelName == "" {
		modelName = "gpt-4o-mini"
	}

	var s model.UserSettings
	err := Pool.QueryRow(ctx, query, userID, apiKey, apiBase, modelName).Scan(
		&s.UserID, &s.APIKey, &s.APIBase, &s.Model, &s.UpdatedAt,
	)

	return &s, err
}
