package repository

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/category"
)

type CategoryRepository struct{ db *pgxpool.Pool }

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository { return &CategoryRepository{db: db} }

func (r *CategoryRepository) GetAll(ctx context.Context) ([]category.Category, error) {
	rows, err := r.db.Query(ctx, `SELECT id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at FROM notification_categories_master ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []category.Category
	for rows.Next() {
		var item category.Category
		if err := rows.Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int) (*category.Category, error) {
	item := &category.Category{}
	if err := r.db.QueryRow(ctx, `SELECT id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at FROM notification_categories_master WHERE id = $1`, id).Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *CategoryRepository) Create(ctx context.Context, req category.CreateRequest) (*category.Category, error) {
	item := &category.Category{}
	code := strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(req.Name), " ", "_"))
	if err := r.db.QueryRow(ctx, `INSERT INTO notification_categories_master (code, name, description, created_at, created_by, is_active, version) VALUES ($1, $2, $3, NOW(), $4, $5, 1) RETURNING id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at`, code, req.Name, req.Description, req.CreatedBy, req.Status == "Active").Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *CategoryRepository) Update(ctx context.Context, req category.UpdateRequest) (*category.Category, error) {
	item := &category.Category{}
	code := strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(req.Name), " ", "_"))
	if err := r.db.QueryRow(ctx, `UPDATE notification_categories_master SET code = $2, name = $3, description = $4, is_active = $5 WHERE id = $1 RETURNING id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at`, req.ID, code, req.Name, req.Description, req.Status == "Active").Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.CreatedBy, &item.IsActive, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE notification_categories_master SET is_active = false WHERE id = $1`, id)
	return err
}
