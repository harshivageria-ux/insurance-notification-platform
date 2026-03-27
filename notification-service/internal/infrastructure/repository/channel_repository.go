package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/channel"
)

type ChannelRepository struct {
	db *pgxpool.Pool
}

func NewChannelRepository(db *pgxpool.Pool) *ChannelRepository {
	return &ChannelRepository{db: db}
}

func (r *ChannelRepository) GetAll(ctx context.Context) ([]channel.Channel, error) {
	query := `
		SELECT id, name, channel_type, is_active, status, created_at, updated_at, deleted_at
		FROM channels
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []channel.Channel
	for rows.Next() {
		var c channel.Channel
		err := rows.Scan(&c.ID, &c.Name, &c.Type, &c.IsActive, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
		if err != nil {
			return nil, err
		}
		channels = append(channels, c)
	}

	return channels, rows.Err()
}

func (r *ChannelRepository) GetByID(ctx context.Context, id int) (*channel.Channel, error) {
	query := `
		SELECT id, name, channel_type, is_active, status, created_at, updated_at, deleted_at
		FROM channels
		WHERE id = $1 AND deleted_at IS NULL
	`
	c := &channel.Channel{}
	err := r.db.QueryRow(ctx, query, id).Scan(&c.ID, &c.Name, &c.Type, &c.IsActive, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ChannelRepository) Create(ctx context.Context, req channel.CreateRequest) (*channel.Channel, error) {
	query := `
		INSERT INTO channels (name, channel_type, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, channel_type, is_active, status, created_at, updated_at, deleted_at
	`
	c := &channel.Channel{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Type, req.Status).
		Scan(&c.ID, &c.Name, &c.Type, &c.IsActive, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ChannelRepository) Update(ctx context.Context, req channel.UpdateRequest) (*channel.Channel, error) {
	query := `
		UPDATE channels
		SET name = $2, channel_type = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, channel_type, is_active, status, created_at, updated_at, deleted_at
	`
	c := &channel.Channel{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Type, req.Status).
		Scan(&c.ID, &c.Name, &c.Type, &c.IsActive, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ChannelRepository) Toggle(ctx context.Context, id int, isActive bool) (*channel.Channel, error) {
	query := `
		UPDATE channels
		SET is_active = $2, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, channel_type, is_active, status, created_at, updated_at, deleted_at
	`
	c := &channel.Channel{}
	err := r.db.QueryRow(ctx, query, id, isActive).
		Scan(&c.ID, &c.Name, &c.Type, &c.IsActive, &c.Status, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ChannelRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE channels
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
