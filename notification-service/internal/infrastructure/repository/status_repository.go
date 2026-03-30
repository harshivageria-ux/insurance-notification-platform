package repository

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/status"
)

type StatusRepository struct{ db *pgxpool.Pool }

func NewStatusRepository(db *pgxpool.Pool) *StatusRepository { return &StatusRepository{db: db} }

func (r *StatusRepository) GetAll(ctx context.Context) ([]status.Status, error) {
	rows, err := r.db.Query(ctx, `SELECT status_id, COALESCE(status_code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_final, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_statuses_master ORDER BY status_id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []status.Status
	for rows.Next() {
		var item status.Status
		if err := rows.Scan(&item.StatusID, &item.StatusCode, &item.Name, &item.Description, &item.IsFinal, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
			return nil, err
		}
		item.ID = item.StatusID
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *StatusRepository) GetByID(ctx context.Context, id int) (*status.Status, error) {
	item := &status.Status{}
	if err := r.db.QueryRow(ctx, `SELECT status_id, COALESCE(status_code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_final, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_statuses_master WHERE status_id = $1`, id).Scan(&item.StatusID, &item.StatusCode, &item.Name, &item.Description, &item.IsFinal, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	item.ID = item.StatusID
	return item, nil
}
func (r *StatusRepository) Create(ctx context.Context, req status.CreateRequest) (*status.Status, error) {
	item := &status.Status{}
	statusCode := strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(req.Name)), " ", "_")
	isActive := strings.EqualFold(req.Status, "Active")
	if req.Status == "" {
		isActive = true
	}
	if req.CreatedBy == "" {
		req.CreatedBy = "system"
	}
	if err := r.db.QueryRow(ctx, `INSERT INTO notification_statuses_master (status_id, status_code, name, description, is_final, created_at, created_by, is_active, version) VALUES ((SELECT COALESCE(MAX(status_id),0)+1 FROM notification_statuses_master), $1, $2, $3, $4, NOW(), $5, $6, 1) RETURNING status_id, COALESCE(status_code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_final, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, statusCode, req.Name, req.Description, req.IsFinal, req.CreatedBy, isActive).Scan(&item.StatusID, &item.StatusCode, &item.Name, &item.Description, &item.IsFinal, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	item.ID = item.StatusID
	return item, nil
}
func (r *StatusRepository) Update(ctx context.Context, req status.UpdateRequest) (*status.Status, error) {
	item := &status.Status{}
	statusCode := strings.ReplaceAll(strings.ToUpper(strings.TrimSpace(req.Name)), " ", "_")
	isActive := strings.EqualFold(req.Status, "Active")
	if req.Status == "" {
		isActive = true
	}
	if err := r.db.QueryRow(ctx, `UPDATE notification_statuses_master SET status_code = $2, name = $3, description = $4, is_final = $5, is_active = $6 WHERE status_id = $1 RETURNING status_id, COALESCE(status_code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_final, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.StatusID, statusCode, req.Name, req.Description, req.IsFinal, isActive).Scan(&item.StatusID, &item.StatusCode, &item.Name, &item.Description, &item.IsFinal, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	item.ID = item.StatusID
	return item, nil
}
func (r *StatusRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE notification_statuses_master SET is_active = false WHERE status_id = $1`, id)
	return err
}
