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
	const query = `
		SELECT priority_id, COALESCE(priority_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)
		FROM notification_priority_master
		ORDER BY priority_id ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []priority.Priority
	for rows.Next() {
		var item priority.Priority
		if err := rows.Scan(
			&item.PriorityID,
			&item.PriorityCode,
			&item.Description,
			&item.CreatedAt,
			&item.CreatedBy,
			&item.IsActive,
			&item.Version,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *PriorityRepository) GetByID(ctx context.Context, id int) (*priority.Priority, error) {
	item := &priority.Priority{}
	if err := r.db.QueryRow(ctx, `SELECT priority_id, COALESCE(priority_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_priority_master WHERE priority_id = $1`, id).Scan(&item.PriorityID, &item.PriorityCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *PriorityRepository) Create(ctx context.Context, req priority.CreateRequest) (*priority.Priority, error) {
	item := &priority.Priority{}
	if err := r.db.QueryRow(ctx, `
		-- Some environments have a broken/absent SMALLSERIAL default; explicitly set priority_id.
		INSERT INTO notification_priority_master (priority_id, priority_code, description, created_at, created_by, is_active, version)
		VALUES (
			(SELECT COALESCE(MAX(priority_id), 0) + 1 FROM notification_priority_master),
			$1, $2, NOW(), $3, true, 1
		)
		RETURNING priority_id, COALESCE(priority_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)
	`, req.PriorityCode, req.Description, req.CreatedBy).Scan(&item.PriorityID, &item.PriorityCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *PriorityRepository) Update(ctx context.Context, req priority.UpdateRequest) (*priority.Priority, error) {
	item := &priority.Priority{}
	if err := r.db.QueryRow(ctx, `
		UPDATE notification_priority_master
		SET priority_code = $2, description = $3, created_by = COALESCE(NULLIF($4, ''), created_by)
		WHERE priority_id = $1
		RETURNING priority_id, COALESCE(priority_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)
	`, req.PriorityID, req.PriorityCode, req.Description, req.CreatedBy).Scan(&item.PriorityID, &item.PriorityCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *PriorityRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE notification_priority_master SET is_active = false WHERE priority_id = $1`, id)
	return err
}
