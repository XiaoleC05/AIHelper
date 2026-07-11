package db

import (
	"context"
	"fmt"

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
	existing, err := r.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	apiKey := ""
	apiBase := ""
	modelName := "gpt-4o-mini"
	if existing != nil {
		apiKey = existing.APIKey
		apiBase = existing.APIBase
		modelName = existing.Model
	}

	if req.APIKey != nil {
		apiKey = *req.APIKey
	}
	if req.APIBase != "" {
		apiBase = req.APIBase
	}
	if req.Model != "" {
		modelName = req.Model
	}

	query := `
		INSERT INTO aihelper.user_settings (user_id, api_key, api_base, model)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id) DO UPDATE
		SET api_key = EXCLUDED.api_key,
		    api_base = EXCLUDED.api_base,
		    model = EXCLUDED.model,
		    updated_at = NOW()
		RETURNING user_id, api_key, api_base, model, updated_at
	`

	var s model.UserSettings
	err = Pool.QueryRow(ctx, query, userID, apiKey, apiBase, modelName).Scan(
		&s.UserID, &s.APIKey, &s.APIBase, &s.Model, &s.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("upsert settings: %w", err)
	}

	return &s, err
}
