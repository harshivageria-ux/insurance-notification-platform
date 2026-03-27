package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	cp "probus-notification-system/internal/domain/channel_provider"
)

type ChannelProviderRepository struct {
	db *pgxpool.Pool
}

func NewChannelProviderRepository(db *pgxpool.Pool) *ChannelProviderRepository {
	return &ChannelProviderRepository{db: db}
}

func (r *ChannelProviderRepository) GetAll(ctx context.Context) ([]cp.ChannelProvider, error) {
	query := `
		SELECT id, name, provider_type, is_active, status, created_at, updated_at, deleted_at
		FROM channel_providers
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var providers []cp.ChannelProvider
	for rows.Next() {
		var p cp.ChannelProvider
		err := rows.Scan(&p.ID, &p.Name, &p.ProviderType, &p.IsActive, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
		if err != nil {
			return nil, err
		}
		providers = append(providers, p)
	}

	return providers, rows.Err()
}

func (r *ChannelProviderRepository) GetByID(ctx context.Context, id int) (*cp.ChannelProvider, error) {
	query := `
		SELECT id, name, provider_type, is_active, status, created_at, updated_at, deleted_at
		FROM channel_providers
		WHERE id = $1 AND deleted_at IS NULL
	`
	p := &cp.ChannelProvider{}
	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.ProviderType, &p.IsActive, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ChannelProviderRepository) Create(ctx context.Context, req cp.CreateRequest) (*cp.ChannelProvider, error) {
	query := `
		INSERT INTO channel_providers (name, provider_type, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, provider_type, is_active, status, created_at, updated_at, deleted_at
	`
	p := &cp.ChannelProvider{}
	err := r.db.QueryRow(ctx, query, req.Name, req.ProviderType, req.Status).
		Scan(&p.ID, &p.Name, &p.ProviderType, &p.IsActive, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ChannelProviderRepository) Update(ctx context.Context, req cp.UpdateRequest) (*cp.ChannelProvider, error) {
	query := `
		UPDATE channel_providers
		SET name = $2, provider_type = $3, status = $4, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, provider_type, is_active, status, created_at, updated_at, deleted_at
	`
	p := &cp.ChannelProvider{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.ProviderType, req.Status).
		Scan(&p.ID, &p.Name, &p.ProviderType, &p.IsActive, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ChannelProviderRepository) Toggle(ctx context.Context, id int, isActive bool) (*cp.ChannelProvider, error) {
	query := `
		UPDATE channel_providers
		SET is_active = $2, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, provider_type, is_active, status, created_at, updated_at, deleted_at
	`
	p := &cp.ChannelProvider{}
	err := r.db.QueryRow(ctx, query, id, isActive).
		Scan(&p.ID, &p.Name, &p.ProviderType, &p.IsActive, &p.Status, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ChannelProviderRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE channel_providers
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
