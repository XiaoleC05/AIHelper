package db

import (
	"context"
	"fmt"

	"github.com/XiaoleC05/AIHelper/internal/model"
	"github.com/jackc/pgx/v5"
)

type PromptRepository struct{}

func NewPromptRepository() *PromptRepository {
	return &PromptRepository{}
}

func (r *PromptRepository) Create(ctx context.Context, userID int64, req model.CreatePromptRequest) (*model.Prompt, error) {
	query := `
		INSERT INTO aihelper.prompts (user_id, title, content, category, tags, variables)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, title, content, category, tags, variables, is_favorite, created_at, updated_at
	`

	category := req.Category
	if category == "" {
		category = "通用"
	}

	var p model.Prompt
	err := Pool.QueryRow(ctx, query,
		userID, req.Title, req.Content, category, req.Tags, req.Variables,
	).Scan(
		&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category,
		&p.Tags, &p.Variables, &p.IsFavorite, &p.CreatedAt, &p.UpdatedAt,
	)

	return &p, err
}

func (r *PromptRepository) GetByID(ctx context.Context, id, userID int64) (*model.Prompt, error) {
	query := `
		SELECT id, user_id, title, content, category, tags, variables, is_favorite, created_at, updated_at
		FROM aihelper.prompts
		WHERE id = $1 AND user_id = $2
	`

	var p model.Prompt
	err := Pool.QueryRow(ctx, query, id, userID).Scan(
		&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category,
		&p.Tags, &p.Variables, &p.IsFavorite, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &p, err
}

func (r *PromptRepository) List(ctx context.Context, userID int64, search, category, tag string, favorite *bool) ([]*model.Prompt, error) {
	query := `
		SELECT id, user_id, title, content, category, tags, variables, is_favorite, created_at, updated_at
		FROM aihelper.prompts
		WHERE user_id = $1
	`
	args := []interface{}{userID}
	argIdx := 2

	if search != "" {
		query += fmt.Sprintf(` AND (title ILIKE $%[1]d OR content ILIKE $%[1]d OR $%[1]d = ANY(tags))`, argIdx)
		args = append(args, "%"+search+"%")
		argIdx++
	}

	if category != "" {
		query += fmt.Sprintf(` AND category = $%d`, argIdx)
		args = append(args, category)
		argIdx++
	}

	if tag != "" {
		query += fmt.Sprintf(` AND $%d = ANY(tags)`, argIdx)
		args = append(args, tag)
		argIdx++
	}

	if favorite != nil {
		query += fmt.Sprintf(` AND is_favorite = $%d`, argIdx)
		args = append(args, *favorite)
		argIdx++
	}

	query += ` ORDER BY updated_at DESC`

	rows, err := Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prompts []*model.Prompt
	for rows.Next() {
		var p model.Prompt
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category,
			&p.Tags, &p.Variables, &p.IsFavorite, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		prompts = append(prompts, &p)
	}

	return prompts, nil
}

func (r *PromptRepository) Update(ctx context.Context, id, userID int64, req model.UpdatePromptRequest) (*model.Prompt, error) {
	query := `
		UPDATE aihelper.prompts
		SET title = COALESCE(NULLIF($3, ''), title),
		    content = COALESCE(NULLIF($4, ''), content),
		    category = COALESCE(NULLIF($5, ''), category),
		    tags = COALESCE($6, tags),
		    variables = COALESCE($7, variables),
		    updated_at = NOW()
		WHERE id = $1 AND user_id = $2
		RETURNING id, user_id, title, content, category, tags, variables, is_favorite, created_at, updated_at
	`

	var p model.Prompt
	err := Pool.QueryRow(ctx, query,
		id, userID, req.Title, req.Content, req.Category, req.Tags, req.Variables,
	).Scan(
		&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category,
		&p.Tags, &p.Variables, &p.IsFavorite, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &p, err
}

func (r *PromptRepository) Delete(ctx context.Context, id, userID int64) error {
	query := `DELETE FROM aihelper.prompts WHERE id = $1 AND user_id = $2`
	_, err := Pool.Exec(ctx, query, id, userID)
	return err
}

func (r *PromptRepository) ToggleFavorite(ctx context.Context, id, userID int64) (*model.Prompt, error) {
	query := `
		UPDATE aihelper.prompts
		SET is_favorite = NOT is_favorite, updated_at = NOW()
		WHERE id = $1 AND user_id = $2
		RETURNING id, user_id, title, content, category, tags, variables, is_favorite, created_at, updated_at
	`

	var p model.Prompt
	err := Pool.QueryRow(ctx, query, id, userID).Scan(
		&p.ID, &p.UserID, &p.Title, &p.Content, &p.Category,
		&p.Tags, &p.Variables, &p.IsFavorite, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &p, err
}
