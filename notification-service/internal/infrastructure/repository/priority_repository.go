package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/priority"
)

type PriorityRepository struct {
	db *pgxpool.Pool
}

func NewPriorityRepository(db *pgxpool.Pool) *PriorityRepository {
	return &PriorityRepository{db: db}
}

func (r *PriorityRepository) GetAll(ctx context.Context) ([]priority.Priority, error) {
	query := `
		SELECT id, name, level, status, created_at, updated_at, deleted_at
		FROM priorities
		WHERE deleted_at IS NULL
		ORDER BY level DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priorities []priority.Priority
	for rows.Next() {
		var p priority.Priority
		err := rows.Scan(&p.ID, &p.Name, &p.Level, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
		if err != nil {
			return nil, err
		}
		priorities = append(priorities, p)
	}

	return priorities, rows.Err()
}

func (r *PriorityRepository) GetByID(ctx context.Context, id int) (*priority.Priority, error) {
	query := `
		SELECT id, name, level, status, created_at, updated_at, deleted_at
		FROM priorities
		WHERE id = $1 AND deleted_at IS NULL
	`
	p := &priority.Priority{}
	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.Level, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PriorityRepository) Create(ctx context.Context, req priority.CreateRequest) (*priority.Priority, error) {
	query := `
		INSERT INTO priorities (name, level, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, level, status, created_at, updated_at, deleted_at
	`
	p := &priority.Priority{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Level, req.Status).
		Scan(&p.ID, &p.Name, &p.Level, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PriorityRepository) Update(ctx context.Context, req priority.UpdateRequest) (*priority.Priority, error) {
	query := `
		UPDATE priorities
		SET name = $2, level = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, level, status, created_at, updated_at, deleted_at
	`
	p := &priority.Priority{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Level, req.Status).
		Scan(&p.ID, &p.Name, &p.Level, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PriorityRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE priorities
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
