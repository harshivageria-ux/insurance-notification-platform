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
	query := `
		SELECT id, name, code, status, created_at, updated_at, deleted_at
		FROM languages
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []language.Language
	for rows.Next() {
		var lang language.Language
		err := rows.Scan(&lang.ID, &lang.Name, &lang.Code, &lang.Status, &lang.CreatedAt, &lang.UpdatedAt, &lang.DeletedAt)
		if err != nil {
			return nil, err
		}
		languages = append(languages, lang)
	}

	return languages, rows.Err()
}

func (r *LanguageRepository) GetByID(ctx context.Context, id int) (*language.Language, error) {
	query := `
		SELECT id, name, code, status, created_at, updated_at, deleted_at
		FROM languages
		WHERE id = $1 AND deleted_at IS NULL
	`
	lang := &language.Language{}
	err := r.db.QueryRow(ctx, query, id).Scan(&lang.ID, &lang.Name, &lang.Code, &lang.Status, &lang.CreatedAt, &lang.UpdatedAt, &lang.DeletedAt)
	if err != nil {
		return nil, err
	}
	return lang, nil
}

func (r *LanguageRepository) Create(ctx context.Context, req language.CreateRequest) (*language.Language, error) {
	query := `
		INSERT INTO languages (name, code, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, code, status, created_at, updated_at, deleted_at
	`
	lang := &language.Language{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Code, req.Status).
		Scan(&lang.ID, &lang.Name, &lang.Code, &lang.Status, &lang.CreatedAt, &lang.UpdatedAt, &lang.DeletedAt)
	if err != nil {
		return nil, err
	}
	return lang, nil
}

func (r *LanguageRepository) Update(ctx context.Context, req language.UpdateRequest) (*language.Language, error) {
	query := `
		UPDATE languages
		SET name = $2, code = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, code, status, created_at, updated_at, deleted_at
	`
	lang := &language.Language{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Code, req.Status).
		Scan(&lang.ID, &lang.Name, &lang.Code, &lang.Status, &lang.CreatedAt, &lang.UpdatedAt, &lang.DeletedAt)
	if err != nil {
		return nil, err
	}
	return lang, nil
}

func (r *LanguageRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE languages
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
