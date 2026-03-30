package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/language"
)

type LanguageRepository struct {
	db *pgxpool.Pool
}

func NewLanguageRepository(db *pgxpool.Pool) *LanguageRepository {
	return &LanguageRepository{db: db}
}

func (r *LanguageRepository) GetAll(ctx context.Context) ([]language.Language, error) {
	const query = `
		SELECT id, name, code, COALESCE(created_by, ''), COALESCE(updated_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at, COALESCE(updated_at, created_at)
		FROM languages_master
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []language.Language
	for rows.Next() {
		var item language.Language
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Code,
			&item.CreatedBy,
			&item.UpdatedBy,
			&item.IsActive,
			&item.Version,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *LanguageRepository) GetByID(ctx context.Context, id int64) (*language.Language, error) {
	const query = `
		SELECT id, name, code, COALESCE(created_by, ''), COALESCE(updated_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at, COALESCE(updated_at, created_at)
		FROM languages_master
		WHERE id = $1
	`
	item := &language.Language{}
	if err := r.db.QueryRow(ctx, query, id).Scan(&item.ID, &item.Name, &item.Code, &item.CreatedBy, &item.UpdatedBy, &item.IsActive, &item.Version, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *LanguageRepository) Create(ctx context.Context, req language.CreateRequest) (*language.Language, error) {
	const query = `
		INSERT INTO languages_master (code, name, created_by, updated_by, is_active, version, created_at, updated_at)
		VALUES ($1, $2, $3, $3, $4, 1, NOW(), NOW())
		RETURNING id, name, code, COALESCE(created_by, ''), COALESCE(updated_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at, COALESCE(updated_at, created_at)
	`
	item := &language.Language{}
	if err := r.db.QueryRow(ctx, query, req.Code, req.Name, req.CreatedBy, req.Status == "Active").Scan(&item.ID, &item.Name, &item.Code, &item.CreatedBy, &item.UpdatedBy, &item.IsActive, &item.Version, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *LanguageRepository) Update(ctx context.Context, req language.UpdateRequest) (*language.Language, error) {
	const query = `
		UPDATE languages_master
		SET code = $2, name = $3, updated_by = $4, is_active = $5, updated_at = NOW()
		WHERE id = $1
		RETURNING id, name, code, COALESCE(created_by, ''), COALESCE(updated_by, ''), COALESCE(is_active, false), COALESCE(version, 0), created_at, COALESCE(updated_at, created_at)
	`
	item := &language.Language{}
	if err := r.db.QueryRow(ctx, query, req.ID, req.Code, req.Name, req.UpdatedBy, req.Status == "Active").Scan(&item.ID, &item.Name, &item.Code, &item.CreatedBy, &item.UpdatedBy, &item.IsActive, &item.Version, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *LanguageRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `UPDATE languages_master SET is_active = false, updated_at = NOW() WHERE id = $1`, id)
	return err
}
