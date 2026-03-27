package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	ps "probus-notification-system/internal/domain/provider_setting"
)

type ProviderSettingRepository struct {
	db *pgxpool.Pool
}

func NewProviderSettingRepository(db *pgxpool.Pool) *ProviderSettingRepository {
	return &ProviderSettingRepository{db: db}
}

func (r *ProviderSettingRepository) GetByProviderID(ctx context.Context, providerID int) ([]ps.ProviderSetting, error) {
	query := `
		SELECT id, provider_id, setting_key, setting_value, description, status, created_at, updated_at, deleted_at
		FROM provider_settings
		WHERE provider_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settings []ps.ProviderSetting
	for rows.Next() {
		var s ps.ProviderSetting
		err := rows.Scan(&s.ID, &s.ProviderID, &s.SettingKey, &s.SettingValue, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, err
		}
		settings = append(settings, s)
	}

	return settings, rows.Err()
}

func (r *ProviderSettingRepository) GetByID(ctx context.Context, id int) (*ps.ProviderSetting, error) {
	query := `
		SELECT id, provider_id, setting_key, setting_value, description, status, created_at, updated_at, deleted_at
		FROM provider_settings
		WHERE id = $1 AND deleted_at IS NULL
	`
	s := &ps.ProviderSetting{}
	err := r.db.QueryRow(ctx, query, id).Scan(&s.ID, &s.ProviderID, &s.SettingKey, &s.SettingValue, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *ProviderSettingRepository) Create(ctx context.Context, req ps.CreateRequest) (*ps.ProviderSetting, error) {
	query := `
		INSERT INTO provider_settings (provider_id, setting_key, setting_value, description, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, provider_id, setting_key, setting_value, description, status, created_at, updated_at, deleted_at
	`
	s := &ps.ProviderSetting{}
	err := r.db.QueryRow(ctx, query, req.ProviderID, req.SettingKey, req.SettingValue, req.Description, req.Status).
		Scan(&s.ID, &s.ProviderID, &s.SettingKey, &s.SettingValue, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *ProviderSettingRepository) Update(ctx context.Context, req ps.UpdateRequest) (*ps.ProviderSetting, error) {
	query := `
		UPDATE provider_settings
		SET provider_id = $2, setting_key = $3, setting_value = $4, description = $5, status = $6, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, provider_id, setting_key, setting_value, description, status, created_at, updated_at, deleted_at
	`
	s := &ps.ProviderSetting{}
	err := r.db.QueryRow(ctx, query, req.ID, req.ProviderID, req.SettingKey, req.SettingValue, req.Description, req.Status).
		Scan(&s.ID, &s.ProviderID, &s.SettingKey, &s.SettingValue, &s.Description, &s.Status, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *ProviderSettingRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE provider_settings
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
