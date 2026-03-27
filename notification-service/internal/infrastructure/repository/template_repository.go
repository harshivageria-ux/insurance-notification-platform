package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/template"
)

type TemplateRepository struct {
	db *pgxpool.Pool
}

func NewTemplateRepository(db *pgxpool.Pool) *TemplateRepository {
	return &TemplateRepository{db: db}
}

func (r *TemplateRepository) GetAll(ctx context.Context) ([]template.Template, error) {
	query := `
		SELECT id, template_group_id, name, content, variables, status, created_at, updated_at, deleted_at
		FROM templates
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []template.Template
	for rows.Next() {
		var t template.Template
		err := rows.Scan(&t.ID, &t.TemplateGroupID, &t.Name, &t.Content, &t.Variables, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return nil, err
		}
		templates = append(templates, t)
	}

	return templates, rows.Err()
}

func (r *TemplateRepository) GetByID(ctx context.Context, id int) (*template.Template, error) {
	query := `
		SELECT id, template_group_id, name, content, variables, status, created_at, updated_at, deleted_at
		FROM templates
		WHERE id = $1 AND deleted_at IS NULL
	`
	t := &template.Template{}
	err := r.db.QueryRow(ctx, query, id).Scan(&t.ID, &t.TemplateGroupID, &t.Name, &t.Content, &t.Variables, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TemplateRepository) Create(ctx context.Context, req template.CreateRequest) (*template.Template, error) {
	query := `
		INSERT INTO templates (template_group_id, name, content, variables, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, template_group_id, name, content, variables, status, created_at, updated_at, deleted_at
	`
	t := &template.Template{}
	err := r.db.QueryRow(ctx, query, req.TemplateGroupID, req.Name, req.Content, req.Variables, req.Status).
		Scan(&t.ID, &t.TemplateGroupID, &t.Name, &t.Content, &t.Variables, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TemplateRepository) Update(ctx context.Context, req template.UpdateRequest) (*template.Template, error) {
	query := `
		UPDATE templates
		SET template_group_id = $2, name = $3, content = $4, variables = $5, status = $6, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, template_group_id, name, content, variables, status, created_at, updated_at, deleted_at
	`
	t := &template.Template{}
	err := r.db.QueryRow(ctx, query, req.ID, req.TemplateGroupID, req.Name, req.Content, req.Variables, req.Status).
		Scan(&t.ID, &t.TemplateGroupID, &t.Name, &t.Content, &t.Variables, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TemplateRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE templates
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
