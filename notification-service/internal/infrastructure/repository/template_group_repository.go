package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	tg "probus-notification-system/internal/domain/template_group"
)

type TemplateGroupRepository struct {
	db *pgxpool.Pool
}

func NewTemplateGroupRepository(db *pgxpool.Pool) *TemplateGroupRepository {
	return &TemplateGroupRepository{db: db}
}

func (r *TemplateGroupRepository) GetAll(ctx context.Context) ([]tg.TemplateGroup, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM template_groups
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []tg.TemplateGroup
	for rows.Next() {
		var g tg.TemplateGroup
		err := rows.Scan(&g.ID, &g.Name, &g.Description, &g.Status, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	return groups, rows.Err()
}

func (r *TemplateGroupRepository) GetByID(ctx context.Context, id int) (*tg.TemplateGroup, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM template_groups
		WHERE id = $1 AND deleted_at IS NULL
	`
	g := &tg.TemplateGroup{}
	err := r.db.QueryRow(ctx, query, id).Scan(&g.ID, &g.Name, &g.Description, &g.Status, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *TemplateGroupRepository) Create(ctx context.Context, req tg.CreateRequest) (*tg.TemplateGroup, error) {
	query := `
		INSERT INTO template_groups (name, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	g := &tg.TemplateGroup{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Description, req.Status).
		Scan(&g.ID, &g.Name, &g.Description, &g.Status, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *TemplateGroupRepository) Update(ctx context.Context, req tg.UpdateRequest) (*tg.TemplateGroup, error) {
	query := `
		UPDATE template_groups
		SET name = $2, description = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	g := &tg.TemplateGroup{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Description, req.Status).
		Scan(&g.ID, &g.Name, &g.Description, &g.Status, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *TemplateGroupRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE template_groups
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
