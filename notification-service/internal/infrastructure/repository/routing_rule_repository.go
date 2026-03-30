package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	rr "probus-notification-system/internal/domain/routing_rule"
)

type RoutingRuleRepository struct{ db *pgxpool.Pool }

func NewRoutingRuleRepository(db *pgxpool.Pool) *RoutingRuleRepository {
	return &RoutingRuleRepository{db: db}
}

func (r *RoutingRuleRepository) GetAll(ctx context.Context) ([]rr.RoutingRule, error) {
	rows, err := r.db.Query(ctx, `SELECT id, template_group_id, channel_id, preferred_provider_id, COALESCE(fallback_provider_id, 0), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM provider_routing_rules_master ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []rr.RoutingRule
	for rows.Next() {
		var item rr.RoutingRule
		if err := rows.Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.PreferredProviderID, &item.FallbackProviderID, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *RoutingRuleRepository) GetByID(ctx context.Context, id int) (*rr.RoutingRule, error) {
	item := &rr.RoutingRule{}
	if err := r.db.QueryRow(ctx, `SELECT id, template_group_id, channel_id, preferred_provider_id, COALESCE(fallback_provider_id, 0), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM provider_routing_rules_master WHERE id = $1`, id).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.PreferredProviderID, &item.FallbackProviderID, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *RoutingRuleRepository) Create(ctx context.Context, req rr.CreateRequest) (*rr.RoutingRule, error) {
	item := &rr.RoutingRule{}
	if err := r.db.QueryRow(ctx, `INSERT INTO provider_routing_rules_master (template_group_id, channel_id, preferred_provider_id, fallback_provider_id, created_at, created_by, is_active, version) VALUES ($1, $2, $3, NULLIF($4, 0), NOW(), $5, true, 1) RETURNING id, template_group_id, channel_id, preferred_provider_id, COALESCE(fallback_provider_id, 0), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.TemplateGroupID, req.ChannelID, req.PreferredProviderID, req.FallbackProviderID, req.CreatedBy).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.PreferredProviderID, &item.FallbackProviderID, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *RoutingRuleRepository) Update(ctx context.Context, req rr.UpdateRequest) (*rr.RoutingRule, error) {
	item := &rr.RoutingRule{}
	if err := r.db.QueryRow(ctx, `UPDATE provider_routing_rules_master SET preferred_provider_id = $2, fallback_provider_id = NULLIF($3, 0) WHERE id = $1 RETURNING id, template_group_id, channel_id, preferred_provider_id, COALESCE(fallback_provider_id, 0), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.ID, req.PreferredProviderID, req.FallbackProviderID).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.PreferredProviderID, &item.FallbackProviderID, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *RoutingRuleRepository) Toggle(ctx context.Context, id int, isActive bool) (*rr.RoutingRule, error) {
	item := &rr.RoutingRule{}
	if err := r.db.QueryRow(ctx, `UPDATE provider_routing_rules_master SET is_active = $2 WHERE id = $1 RETURNING id, template_group_id, channel_id, preferred_provider_id, COALESCE(fallback_provider_id, 0), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, id, isActive).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.PreferredProviderID, &item.FallbackProviderID, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *RoutingRuleRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE provider_routing_rules_master SET is_active = false WHERE id = $1`, id)
	return err
}
