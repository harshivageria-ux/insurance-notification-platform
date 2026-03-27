package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	rr "probus-notification-system/internal/domain/routing_rule"
)

type RoutingRuleRepository struct {
	db *pgxpool.Pool
}

func NewRoutingRuleRepository(db *pgxpool.Pool) *RoutingRuleRepository {
	return &RoutingRuleRepository{db: db}
}

func (r *RoutingRuleRepository) GetAll(ctx context.Context) ([]rr.RoutingRule, error) {
	query := `
		SELECT id, name, condition, target_channel, priority, is_active, status, created_at, updated_at, deleted_at
		FROM routing_rules
		WHERE deleted_at IS NULL
		ORDER BY priority DESC, created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []rr.RoutingRule
	for rows.Next() {
		var rule rr.RoutingRule
		err := rows.Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.TargetChannel, &rule.Priority, &rule.IsActive, &rule.Status, &rule.CreatedAt, &rule.UpdatedAt, &rule.DeletedAt)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	return rules, rows.Err()
}

func (r *RoutingRuleRepository) GetByID(ctx context.Context, id int) (*rr.RoutingRule, error) {
	query := `
		SELECT id, name, condition, target_channel, priority, is_active, status, created_at, updated_at, deleted_at
		FROM routing_rules
		WHERE id = $1 AND deleted_at IS NULL
	`
	rule := &rr.RoutingRule{}
	err := r.db.QueryRow(ctx, query, id).Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.TargetChannel, &rule.Priority, &rule.IsActive, &rule.Status, &rule.CreatedAt, &rule.UpdatedAt, &rule.DeletedAt)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *RoutingRuleRepository) Create(ctx context.Context, req rr.CreateRequest) (*rr.RoutingRule, error) {
	query := `
		INSERT INTO routing_rules (name, condition, target_channel, priority, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, condition, target_channel, priority, is_active, status, created_at, updated_at, deleted_at
	`
	rule := &rr.RoutingRule{}
	err := r.db.QueryRow(ctx, query, req.Name, req.Condition, req.TargetChannel, req.Priority, req.Status).
		Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.TargetChannel, &rule.Priority, &rule.IsActive, &rule.Status, &rule.CreatedAt, &rule.UpdatedAt, &rule.DeletedAt)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *RoutingRuleRepository) Update(ctx context.Context, req rr.UpdateRequest) (*rr.RoutingRule, error) {
	query := `
		UPDATE routing_rules
		SET name = $2, condition = $3, target_channel = $4, priority = $5, status = $6, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, condition, target_channel, priority, is_active, status, created_at, updated_at, deleted_at
	`
	rule := &rr.RoutingRule{}
	err := r.db.QueryRow(ctx, query, req.ID, req.Name, req.Condition, req.TargetChannel, req.Priority, req.Status).
		Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.TargetChannel, &rule.Priority, &rule.IsActive, &rule.Status, &rule.CreatedAt, &rule.UpdatedAt, &rule.DeletedAt)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *RoutingRuleRepository) Toggle(ctx context.Context, id int, isActive bool) (*rr.RoutingRule, error) {
	query := `
		UPDATE routing_rules
		SET is_active = $2, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, name, condition, target_channel, priority, is_active, status, created_at, updated_at, deleted_at
	`
	rule := &rr.RoutingRule{}
	err := r.db.QueryRow(ctx, query, id, isActive).
		Scan(&rule.ID, &rule.Name, &rule.Condition, &rule.TargetChannel, &rule.Priority, &rule.IsActive, &rule.Status, &rule.CreatedAt, &rule.UpdatedAt, &rule.DeletedAt)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *RoutingRuleRepository) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE routing_rules
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
