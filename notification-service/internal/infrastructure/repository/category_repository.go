package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/category"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]category.Category, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM categories
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []category.Category
	for rows.Next() {
		var c category.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int) (*category.Category, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM categories
		WHERE id = $1 AND deleted_at IS NULL
	`
	c := &category.Category{}
	err := r.db.QueryRow(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Create(ctx context.Context, req category.CreateRequest) (*category.Category, error) {
	query := `
		INSERT INTO categories (name, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	c := &category.Category{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Description, req.Status).
		Scan(&c.ID, &c.Name, &c.Description, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Update(ctx context.Context, req category.UpdateRequest) (*category.Category, error) {
	query := `
		UPDATE categories
		SET name = $2, description = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	c := &category.Category{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Description, req.Status).
		Scan(&c.ID, &c.Name, &c.Description, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE categories
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
