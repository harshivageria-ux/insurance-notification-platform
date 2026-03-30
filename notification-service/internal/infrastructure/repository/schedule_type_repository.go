package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	st "probus-notification-system/internal/domain/schedule_type"
)

type ScheduleTypeRepository struct{ db *pgxpool.Pool }

func NewScheduleTypeRepository(db *pgxpool.Pool) *ScheduleTypeRepository {
	return &ScheduleTypeRepository{db: db}
}

func (r *ScheduleTypeRepository) GetAll(ctx context.Context) ([]st.ScheduleType, error) {
	rows, err := r.db.Query(ctx, `SELECT schedule_type_id, COALESCE(schedule_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_schedule_type_master ORDER BY schedule_type_id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []st.ScheduleType
	for rows.Next() {
		var item st.ScheduleType
		if err := rows.Scan(&item.ScheduleTypeID, &item.ScheduleCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *ScheduleTypeRepository) GetByID(ctx context.Context, id int) (*st.ScheduleType, error) {
	item := &st.ScheduleType{}
	if err := r.db.QueryRow(ctx, `SELECT schedule_type_id, COALESCE(schedule_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_schedule_type_master WHERE schedule_type_id = $1`, id).Scan(&item.ScheduleTypeID, &item.ScheduleCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ScheduleTypeRepository) Create(ctx context.Context, req st.CreateRequest) (*st.ScheduleType, error) {
	item := &st.ScheduleType{}
	if err := r.db.QueryRow(ctx, `INSERT INTO notification_schedule_type_master (schedule_code, description, created_at, created_by, is_active, version) VALUES ($1, $2, NOW(), $3, true, 1) RETURNING schedule_type_id, COALESCE(schedule_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.ScheduleCode, req.Description, req.CreatedBy).Scan(&item.ScheduleTypeID, &item.ScheduleCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ScheduleTypeRepository) Update(ctx context.Context, req st.UpdateRequest) (*st.ScheduleType, error) {
	item := &st.ScheduleType{}
	if err := r.db.QueryRow(ctx, `UPDATE notification_schedule_type_master SET schedule_code = $2, description = $3 WHERE schedule_type_id = $1 RETURNING schedule_type_id, COALESCE(schedule_code, ''), COALESCE(description, ''), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.ScheduleTypeID, req.ScheduleCode, req.Description).Scan(&item.ScheduleTypeID, &item.ScheduleCode, &item.Description, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ScheduleTypeRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE notification_schedule_type_master SET is_active = false WHERE schedule_type_id = $1`, id)
	return err
}
