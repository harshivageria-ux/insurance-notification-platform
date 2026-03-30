package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/domain/template"
)

type TemplateRepository struct{ db *pgxpool.Pool }

func NewTemplateRepository(db *pgxpool.Pool) *TemplateRepository { return &TemplateRepository{db: db} }

func (r *TemplateRepository) GetAll(ctx context.Context) ([]template.Template, error) {
	rows, err := r.db.Query(ctx, `SELECT id, template_group_id, channel_id, language_id, COALESCE(title_template, ''), COALESCE(message_template, ''), COALESCE(has_variables, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_templates_master ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []template.Template
	for rows.Next() {
		var item template.Template
		if err := rows.Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.LanguageID, &item.TitleTemplate, &item.MessageTemplate, &item.HasVariables, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *TemplateRepository) GetByID(ctx context.Context, id int) (*template.Template, error) {
	item := &template.Template{}
	if err := r.db.QueryRow(ctx, `SELECT id, template_group_id, channel_id, language_id, COALESCE(title_template, ''), COALESCE(message_template, ''), COALESCE(has_variables, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0) FROM notification_templates_master WHERE id = $1`, id).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.LanguageID, &item.TitleTemplate, &item.MessageTemplate, &item.HasVariables, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *TemplateRepository) Create(ctx context.Context, req template.CreateRequest) (*template.Template, error) {
	item := &template.Template{}
	if err := r.db.QueryRow(ctx, `INSERT INTO notification_templates_master (template_group_id, channel_id, language_id, title_template, message_template, has_variables, created_at, created_by, is_active, version) VALUES ($1, $2, $3, $4, $5, $6, NOW(), $7, true, 1) RETURNING id, template_group_id, channel_id, language_id, COALESCE(title_template, ''), COALESCE(message_template, ''), COALESCE(has_variables, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.TemplateGroupID, req.ChannelID, req.LanguageID, req.TitleTemplate, req.MessageTemplate, req.HasVariables, req.CreatedBy).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.LanguageID, &item.TitleTemplate, &item.MessageTemplate, &item.HasVariables, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *TemplateRepository) Update(ctx context.Context, req template.UpdateRequest) (*template.Template, error) {
	item := &template.Template{}
	if err := r.db.QueryRow(ctx, `UPDATE notification_templates_master SET template_group_id = $2, channel_id = $3, language_id = $4, title_template = $5, message_template = $6, has_variables = $7 WHERE id = $1 RETURNING id, template_group_id, channel_id, language_id, COALESCE(title_template, ''), COALESCE(message_template, ''), COALESCE(has_variables, false), created_at, COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)`, req.ID, req.TemplateGroupID, req.ChannelID, req.LanguageID, req.TitleTemplate, req.MessageTemplate, req.HasVariables).Scan(&item.ID, &item.TemplateGroupID, &item.ChannelID, &item.LanguageID, &item.TitleTemplate, &item.MessageTemplate, &item.HasVariables, &item.CreatedAt, &item.CreatedBy, &item.IsActive, &item.Version); err != nil {
		return nil, err
	}
	return item, nil
}
func (r *TemplateRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `UPDATE notification_templates_master SET is_active = false WHERE id = $1`, id)
	return err
}
