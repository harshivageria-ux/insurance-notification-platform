package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	st "probus-notification-system/internal/domain/schedule_type"
)

type ScheduleTypeRepository struct {
	db *pgxpool.Pool
}

func NewScheduleTypeRepository(db *pgxpool.Pool) *ScheduleTypeRepository {
	return &ScheduleTypeRepository{db: db}
}

func (r *ScheduleTypeRepository) GetAll(ctx context.Context) ([]st.ScheduleType, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM schedule_types
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []st.ScheduleType
	for rows.Next() {
		var t st.ScheduleType
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, rows.Err()
}

func (r *ScheduleTypeRepository) GetByID(ctx context.Context, id int) (*st.ScheduleType, error) {
	query := `
		SELECT id, name, description, status, created_at, updated_at, deleted_at
		FROM schedule_types
		WHERE id = $1 AND deleted_at IS NULL
	`
	t := &st.ScheduleType{}
	err := r.db.QueryRow(ctx, query, id).Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *ScheduleTypeRepository) Create(ctx context.Context, req st.CreateRequest) (*st.ScheduleType, error) {
	query := `
		INSERT INTO schedule_types (name, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	t := &st.ScheduleType{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Description, req.Status).
		Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *ScheduleTypeRepository) Update(ctx context.Context, req st.UpdateRequest) (*st.ScheduleType, error) {
	query := `
		UPDATE schedule_types
		SET name = $2, description = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, description, status, created_at, updated_at, deleted_at
	`
	t := &st.ScheduleType{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Description, req.Status).
		Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *ScheduleTypeRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE schedule_types
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
