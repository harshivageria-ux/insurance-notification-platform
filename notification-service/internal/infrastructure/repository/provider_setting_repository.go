package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	ps "probus-notification-system/internal/domain/provider_setting"
)

type ProviderSettingRepository struct{ db *pgxpool.Pool }

func NewProviderSettingRepository(db *pgxpool.Pool) *ProviderSettingRepository {
	return &ProviderSettingRepository{db: db}
}

func (r *ProviderSettingRepository) GetByProviderID(ctx context.Context, providerID int) ([]ps.ProviderSetting, error) {
	rows, err := r.db.Query(ctx, `SELECT id, provider_id, COALESCE(setting_key, ''), COALESCE(setting_value, ''), COALESCE(is_sensitive, false), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at FROM channel_provider_settings_master WHERE provider_id = $1 ORDER BY created_at DESC`, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ps.ProviderSetting
	for rows.Next() {
		var item ps.ProviderSetting
		if err := rows.Scan(&item.ID, &item.ProviderID, &item.SettingKey, &item.SettingValue, &item.IsSensitive, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *ProviderSettingRepository) GetByID(ctx context.Context, id int) (*ps.ProviderSetting, error) {
	item := &ps.ProviderSetting{}
	if err := r.db.QueryRow(ctx, `SELECT id, provider_id, COALESCE(setting_key, ''), COALESCE(setting_value, ''), COALESCE(is_sensitive, false), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at FROM channel_provider_settings_master WHERE id = $1`, id).Scan(&item.ID, &item.ProviderID, &item.SettingKey, &item.SettingValue, &item.IsSensitive, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ProviderSettingRepository) Create(ctx context.Context, req ps.CreateRequest) (*ps.ProviderSetting, error) {
	item := &ps.ProviderSetting{}
	if err := r.db.QueryRow(ctx, `INSERT INTO channel_provider_settings_master (provider_id, setting_key, setting_value, is_sensitive, is_active, created_at, created_by, version) VALUES ($1, $2, $3, $4, true, NOW(), $5, 1) RETURNING id, provider_id, COALESCE(setting_key, ''), COALESCE(setting_value, ''), COALESCE(is_sensitive, false), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at`, req.ProviderID, req.SettingKey, req.SettingValue, req.IsSensitive, req.CreatedBy).Scan(&item.ID, &item.ProviderID, &item.SettingKey, &item.SettingValue, &item.IsSensitive, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ProviderSettingRepository) Update(ctx context.Context, req ps.UpdateRequest) (*ps.ProviderSetting, error) {
	item := &ps.ProviderSetting{}
	if err := r.db.QueryRow(ctx, `UPDATE channel_provider_settings_master SET setting_key = $2, setting_value = $3, is_sensitive = $4 WHERE id = $1 RETURNING id, provider_id, COALESCE(setting_key, ''), COALESCE(setting_value, ''), COALESCE(is_sensitive, false), COALESCE(is_active, false), COALESCE(created_by, ''), COALESCE(version, 0), created_at`, req.ID, req.SettingKey, req.SettingValue, req.IsSensitive).Scan(&item.ID, &item.ProviderID, &item.SettingKey, &item.SettingValue, &item.IsSensitive, &item.IsActive, &item.CreatedBy, &item.Version, &item.CreatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *ProviderSettingRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE channel_provider_settings_master SET is_active = false WHERE id = $1`, id)
	return err
}
