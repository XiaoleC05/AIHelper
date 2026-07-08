package db

import (
	"context"

	"github.com/XiaoleC05/AIHelper/internal/model"
)

type TemplateRepository struct{}

func NewTemplateRepository() *TemplateRepository {
	return &TemplateRepository{}
}

func (r *TemplateRepository) List(ctx context.Context, category string) ([]*model.Template, error) {
	query := `
		SELECT id, name, category, content, created_at
		FROM aihelper.templates
	`
	args := []interface{}{}

	if category != "" {
		query += ` WHERE category = $1`
		args = append(args, category)
	}

	query += ` ORDER BY category, id`

	rows, err := Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*model.Template
	for rows.Next() {
		var t model.Template
		if err := rows.Scan(&t.ID, &t.Name, &t.Category, &t.Content, &t.CreatedAt); err != nil {
			return nil, err
		}
		templates = append(templates, &t)
	}

	return templates, nil
}

func (r *TemplateRepository) GetByID(ctx context.Context, id int64) (*model.Template, error) {
	query := `
		SELECT id, name, category, content, created_at
		FROM aihelper.templates
		WHERE id = $1
	`

	var t model.Template
	err := Pool.QueryRow(ctx, query, id).Scan(&t.ID, &t.Name, &t.Category, &t.Content, &t.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
