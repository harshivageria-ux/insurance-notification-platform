package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/status"
)

type StatusRepository struct {
	db *pgxpool.Pool
}

func NewStatusRepository(db *pgxpool.Pool) *StatusRepository {
	return &StatusRepository{db: db}
}

func (r *StatusRepository) GetAll(ctx context.Context) ([]status.Status, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM statuses
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []status.Status
	for rows.Next() {
		var s status.Status
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}

	return statuses, rows.Err()
}

func (r *StatusRepository) GetByID(ctx context.Context, id int) (*status.Status, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM statuses
		WHERE id = $1 AND deleted_at IS NULL
	`
	s := &status.Status{}
	err := r.db.QueryRow(ctx, query, id).Scan(&s.ID, &s.Name, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StatusRepository) Create(ctx context.Context, req status.CreateRequest) (*status.Status, error) {
	query := `
		INSERT INTO statuses (name, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	s := &status.Status{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Description, req.Status).
		Scan(&s.ID, &s.Name, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StatusRepository) Update(ctx context.Context, req status.UpdateRequest) (*status.Status, error) {
	query := `
		UPDATE statuses
		SET name = $2, description = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	s := &status.Status{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Description, req.Status).
		Scan(&s.ID, &s.Name, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StatusRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE statuses
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
