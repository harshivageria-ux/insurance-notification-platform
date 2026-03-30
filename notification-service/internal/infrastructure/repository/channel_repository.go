package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/channel"
)

type ChannelRepository struct{ db *pgxpool.Pool }

func NewChannelRepository(db *pgxpool.Pool) *ChannelRepository { return &ChannelRepository{db: db} }

func (r *ChannelRepository) GetAll(ctx context.Context) ([]channel.Channel, error) {
	rows, err := r.db.Query(ctx, `SELECT id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at FROM notification_channels_master ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []channel.Channel
	for rows.Next() {
		var item channel.Channel
		if err := rows.Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *ChannelRepository) GetByID(ctx context.Context, id int) (*channel.Channel, error) {
	item := &channel.Channel{}
	if err := r.db.QueryRow(ctx, `SELECT id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at FROM notification_channels_master WHERE id = $1`, id).Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelRepository) Create(ctx context.Context, req channel.CreateRequest) (*channel.Channel, error) {
	item := &channel.Channel{}
	if err := r.db.QueryRow(ctx, `INSERT INTO notification_channels_master (code, name, description, created_at, created_by, is_active, version) VALUES ($1, $2, $3, NOW(), $4, true, 1) RETURNING id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at`, req.Code, req.Name, req.Description, req.CreatedBy).Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelRepository) Update(ctx context.Context, req channel.UpdateRequest) (*channel.Channel, error) {
	item := &channel.Channel{}
	if err := r.db.QueryRow(ctx, `UPDATE notification_channels_master SET code = $2, name = $3, description = $4 WHERE id = $1 RETURNING id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at`, req.ID, req.Code, req.Name, req.Description).Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelRepository) Toggle(ctx context.Context, id int, isActive bool) (*channel.Channel, error) {
	item := &channel.Channel{}
	if err := r.db.QueryRow(ctx, `UPDATE notification_channels_master SET is_active = $2 WHERE id = $1 RETURNING id, COALESCE(code, ''), COALESCE(name, ''), COALESCE(description, ''), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at`, id, isActive).Scan(&item.ID, &item.Code, &item.Name, &item.Description, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE notification_channels_master SET is_active = false WHERE id = $1`, id)
	return err
}
