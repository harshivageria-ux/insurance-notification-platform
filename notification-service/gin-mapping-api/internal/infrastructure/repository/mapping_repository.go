package repository

import (
	"context"
	"fmt"

	"probus-notification-system/gin-mapping-api/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

// MappingRepository handles mapping table DB operations
// It uses parameterized queries and soft-delete logic.

type MappingRepository struct {
	pool *pgxpool.Pool
}

func NewMappingRepository(pool *pgxpool.Pool) *MappingRepository {
	return &MappingRepository{pool: pool}
}

func (r *MappingRepository) AddCategoryChannel(ctx context.Context, req domain.CreateNotificationCategoryChannelRequest) (domain.NotificationCategoryChannel, error) {
	var rec domain.NotificationCategoryChannel
	query := `INSERT INTO notification_category_channel_master (category_id, channel_id, is_active, created_at) VALUES ($1, $2, TRUE, NOW()) RETURNING id, category_id, channel_id, is_active, created_at`
	row := r.pool.QueryRow(ctx, query, req.CategoryID, req.ChannelID)
	if err := row.Scan(&rec.ID, &rec.CategoryID, &rec.ChannelID, &rec.IsActive, &rec.CreatedAt); err != nil {
		return rec, fmt.Errorf("insert category channel mapping: %w", err)
	}
	return rec, nil
}

func (r *MappingRepository) GetCategoryChannels(ctx context.Context, limit int, offset int) ([]domain.NotificationCategoryChannel, error) {
	query := `SELECT id, category_id, channel_id, is_active, created_at FROM notification_category_channel_master WHERE is_active = TRUE ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []domain.NotificationCategoryChannel
	for rows.Next() {
		var m domain.NotificationCategoryChannel
		if err := rows.Scan(&m.ID, &m.CategoryID, &m.ChannelID, &m.IsActive, &m.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, rows.Err()
}

func (r *MappingRepository) SoftDeleteCategoryChannel(ctx context.Context, id int) error {
	query := `UPDATE notification_category_channel_master SET is_active = FALSE WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}

func (r *MappingRepository) AddChannelProvider(ctx context.Context, req domain.CreateChannelProviderMasterMapRequest) (domain.ChannelProviderMasterMap, error) {
	var rec domain.ChannelProviderMasterMap
	query := `INSERT INTO channel_provider_master_map (channel_id, provider_id, priority, is_active) VALUES ($1, $2, $3, TRUE) RETURNING id, channel_id, provider_id, priority, is_active`
	row := r.pool.QueryRow(ctx, query, req.ChannelID, req.ProviderID, req.Priority)
	if err := row.Scan(&rec.ID, &rec.ChannelID, &rec.ProviderID, &rec.Priority, &rec.IsActive); err != nil {
		return rec, fmt.Errorf("insert channel provider map: %w", err)
	}
	return rec, nil
}

func (r *MappingRepository) GetChannelProviders(ctx context.Context, limit int, offset int) ([]domain.ChannelProviderMasterMap, error) {
	query := `SELECT id, channel_id, provider_id, priority, is_active FROM channel_provider_master_map WHERE is_active = TRUE ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []domain.ChannelProviderMasterMap
	for rows.Next() {
		var m domain.ChannelProviderMasterMap
		if err := rows.Scan(&m.ID, &m.ChannelID, &m.ProviderID, &m.Priority, &m.IsActive); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, rows.Err()
}

func (r *MappingRepository) SoftDeleteChannelProvider(ctx context.Context, id int) error {
	query := `UPDATE channel_provider_master_map SET is_active = FALSE WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}

func (r *MappingRepository) AddTemplateChannelLanguage(ctx context.Context, req domain.CreateTemplateChannelLanguageMasterMapRequest) (domain.TemplateChannelLanguageMasterMap, error) {
	var rec domain.TemplateChannelLanguageMasterMap
	query := `INSERT INTO template_channel_language_master_map (template_group_id, template_id, channel_id, language_id, is_active, created_at) VALUES ($1, $2, $3, $4, TRUE, NOW()) RETURNING id, template_group_id, template_id, channel_id, language_id, is_active, created_at`
	row := r.pool.QueryRow(ctx, query, req.TemplateGroupID, req.TemplateID, req.ChannelID, req.LanguageID)
	if err := row.Scan(&rec.ID, &rec.TemplateGroupID, &rec.TemplateID, &rec.ChannelID, &rec.LanguageID, &rec.IsActive, &rec.CreatedAt); err != nil {
		return rec, fmt.Errorf("insert template channel language: %w", err)
	}
	return rec, nil
}

func (r *MappingRepository) GetTemplateChannelLanguages(ctx context.Context, limit int, offset int) ([]domain.TemplateChannelLanguageMasterMap, error) {
	query := `SELECT id, template_group_id, template_id, channel_id, language_id, is_active, created_at FROM template_channel_language_master_map WHERE is_active = TRUE ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []domain.TemplateChannelLanguageMasterMap
	for rows.Next() {
		var m domain.TemplateChannelLanguageMasterMap
		if err := rows.Scan(&m.ID, &m.TemplateGroupID, &m.TemplateID, &m.ChannelID, &m.LanguageID, &m.IsActive, &m.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, rows.Err()
}

func (r *MappingRepository) SoftDeleteTemplateChannelLanguage(ctx context.Context, id int) error {
	query := `UPDATE template_channel_language_master_map SET is_active = FALSE WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
