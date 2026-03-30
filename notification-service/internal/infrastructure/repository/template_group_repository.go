package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	tg "probus-notification-system/internal/domain/template_group"
)

type TemplateGroupRepository struct{ db *pgxpool.Pool }

func NewTemplateGroupRepository(db *pgxpool.Pool) *TemplateGroupRepository {
	return &TemplateGroupRepository{db: db}
}

func (r *TemplateGroupRepository) GetAll(ctx context.Context) ([]tg.TemplateGroup, error) {
	rows, err := r.db.Query(ctx, `SELECT id, COALESCE(name, ''), category_id, COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at FROM template_groups_master ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []tg.TemplateGroup
	for rows.Next() {
		var item tg.TemplateGroup
		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *TemplateGroupRepository) GetByID(ctx context.Context, id int) (*tg.TemplateGroup, error) {
	item := &tg.TemplateGroup{}
	if err := r.db.QueryRow(ctx, `SELECT id, COALESCE(name, ''), category_id, COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at FROM template_groups_master WHERE id = $1`, id).Scan(&item.ID, &item.Name, &item.CategoryID, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *TemplateGroupRepository) Create(ctx context.Context, req tg.CreateRequest) (*tg.TemplateGroup, error) {
	item := &tg.TemplateGroup{}
	if err := r.db.QueryRow(ctx, `INSERT INTO template_groups_master (name, category_id, description, created_at, created_by, is_active, version) VALUES ($1, $2, $3, NOW(), $4, true, 1) RETURNING id, COALESCE(name, ''), category_id, COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at`, req.Name, req.CategoryID, req.Description, req.CreatedBy).Scan(&item.ID, &item.Name, &item.CategoryID, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *TemplateGroupRepository) Update(ctx context.Context, req tg.UpdateRequest) (*tg.TemplateGroup, error) {
	item := &tg.TemplateGroup{}
	if err := r.db.QueryRow(ctx, `UPDATE template_groups_master SET name = $2, category_id = $3, description = $4 WHERE id = $1 RETURNING id, COALESCE(name, ''), category_id, COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at`, req.ID, req.Name, req.CategoryID, req.Description).Scan(&item.ID, &item.Name, &item.CategoryID, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *TemplateGroupRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE template_groups_master SET is_active = false WHERE id = $1`, id)
	return err
}
