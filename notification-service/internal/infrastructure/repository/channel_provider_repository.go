package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	cp "probus-notification-system/internal/domain/channel_provider"
)

type ChannelProviderRepository struct{ db *pgxpool.Pool }

func NewChannelProviderRepository(db *pgxpool.Pool) *ChannelProviderRepository {
	return &ChannelProviderRepository{db: db}
}

func (r *ChannelProviderRepository) GetAll(ctx context.Context) ([]cp.ChannelProvider, error) {
	rows, err := r.db.Query(ctx, `SELECT id, channel_id, COALESCE(name, ''), COALESCE(code, ''), COALESCE(priority, 0), COALESCE(is_active, false), created_at, COALESCE(created_by, ''), COALESCE(version, 0) FROM channel_providers_master ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []cp.ChannelProvider
	for rows.Next() {
		var item cp.ChannelProvider
		if err := rows.Scan(&item.ID, &item.ChannelID, &item.Name, &item.Code, &item.Priority, &item.IsActive, &item.CreatedAt, &item.CreatedBy, &item.Version); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *ChannelProviderRepository) GetByID(ctx context.Context, id int) (*cp.ChannelProvider, error) {
	item := &cp.ChannelProvider{}
	if err := r.db.QueryRow(ctx, `SELECT id, channel_id, COALESCE(name, ''), COALESCE(code, ''), COALESCE(priority, 0), COALESCE(is_active, false), created_at, COALESCE(created_by, ''), COALESCE(version, 0) FROM channel_providers_master WHERE id = $1`, id).Scan(&item.ID, &item.ChannelID, &item.Name, &item.Code, &item.Priority, &item.IsActive, &item.CreatedAt, &item.CreatedBy, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelProviderRepository) Create(ctx context.Context, req cp.CreateRequest) (*cp.ChannelProvider, error) {
	item := &cp.ChannelProvider{}
	if err := r.db.QueryRow(ctx, `INSERT INTO channel_providers_master (channel_id, name, code, priority, is_active, created_at, created_by, version) VALUES ($1, $2, $3, $4, true, NOW(), $5, 1) RETURNING id, channel_id, COALESCE(name, ''), COALESCE(code, ''), COALESCE(priority, 0), COALESCE(is_active, false), created_at, COALESCE(created_by, ''), COALESCE(version, 0)`, req.ChannelID, req.Name, req.Code, req.Priority, req.CreatedBy).Scan(&item.ID, &item.ChannelID, &item.Name, &item.Code, &item.Priority, &item.IsActive, &item.CreatedAt, &item.CreatedBy, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelProviderRepository) Update(ctx context.Context, req cp.UpdateRequest) (*cp.ChannelProvider, error) {
	item := &cp.ChannelProvider{}
	if err := r.db.QueryRow(ctx, `UPDATE channel_providers_master SET name = $2, priority = $3 WHERE id = $1 RETURNING id, channel_id, COALESCE(name, ''), COALESCE(code, ''), COALESCE(priority, 0), COALESCE(is_active, false), created_at, COALESCE(created_by, ''), COALESCE(version, 0)`, req.ID, req.Name, req.Priority).Scan(&item.ID, &item.ChannelID, &item.Name, &item.Code, &item.Priority, &item.IsActive, &item.CreatedAt, &item.CreatedBy, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelProviderRepository) Toggle(ctx context.Context, id int, isActive bool) (*cp.ChannelProvider, error) {
	item := &cp.ChannelProvider{}
	if err := r.db.QueryRow(ctx, `UPDATE channel_providers_master SET is_active = $2 WHERE id = $1 RETURNING id, channel_id, COALESCE(name, ''), COALESCE(code, ''), COALESCE(priority, 0), COALESCE(is_active, false), created_at, COALESCE(created_by, ''), COALESCE(version, 0)`, id, isActive).Scan(&item.ID, &item.ChannelID, &item.Name, &item.Code, &item.Priority, &item.IsActive, &item.CreatedAt, &item.CreatedBy, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ChannelProviderRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE channel_providers_master SET is_active = false WHERE id = $1`, id)
	return err
}
