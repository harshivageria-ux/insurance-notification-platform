package service

import (
	"context"
	"fmt"

	"probus-notification-system/gin-mapping-api/internal/domain"
	"probus-notification-system/gin-mapping-api/internal/infrastructure/repository"
)

type MappingService struct {
	repo *repository.MappingRepository
}

func NewMappingService(repo *repository.MappingRepository) *MappingService {
	return &MappingService{repo: repo}
}

func (s *MappingService) AddCategoryChannel(ctx context.Context, req domain.CreateNotificationCategoryChannelRequest) (domain.NotificationCategoryChannel, error) {
	if req.CategoryID <= 0 {
		return domain.NotificationCategoryChannel{}, fmt.Errorf("category_id is required and must be positive")
	}
	if req.ChannelID <= 0 {
		return domain.NotificationCategoryChannel{}, fmt.Errorf("channel_id is required and must be positive")
	}
	return s.repo.AddCategoryChannel(ctx, req)
}

func (s *MappingService) GetCategoryChannels(ctx context.Context, limit, offset int) ([]domain.NotificationCategoryChannel, error) {
	if limit <= 0 || limit > 100 {
		limit = 25
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetCategoryChannels(ctx, limit, offset)
}

func (s *MappingService) DeleteCategoryChannel(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id is required")
	}
	return s.repo.SoftDeleteCategoryChannel(ctx, id)
}

func (s *MappingService) AddChannelProvider(ctx context.Context, req domain.CreateChannelProviderMasterMapRequest) (domain.ChannelProviderMasterMap, error) {
	if req.ChannelID <= 0 {
		return domain.ChannelProviderMasterMap{}, fmt.Errorf("channel_id is required and must be positive")
	}
	if req.ProviderID <= 0 {
		return domain.ChannelProviderMasterMap{}, fmt.Errorf("provider_id is required and must be positive")
	}
	if req.Priority <= 0 {
		return domain.ChannelProviderMasterMap{}, fmt.Errorf("priority is required and must be positive")
	}
	return s.repo.AddChannelProvider(ctx, req)
}

func (s *MappingService) GetChannelProviders(ctx context.Context, limit, offset int) ([]domain.ChannelProviderMasterMap, error) {
	if limit <= 0 || limit > 100 {
		limit = 25
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetChannelProviders(ctx, limit, offset)
}

func (s *MappingService) DeleteChannelProvider(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id is required")
	}
	return s.repo.SoftDeleteChannelProvider(ctx, id)
}

func (s *MappingService) AddTemplateChannelLanguage(ctx context.Context, req domain.CreateTemplateChannelLanguageMasterMapRequest) (domain.TemplateChannelLanguageMasterMap, error) {
	if req.TemplateGroupID <= 0 {
		return domain.TemplateChannelLanguageMasterMap{}, fmt.Errorf("template_group_id is required and must be positive")
	}
	if req.TemplateID <= 0 {
		return domain.TemplateChannelLanguageMasterMap{}, fmt.Errorf("template_id is required and must be positive")
	}
	if req.ChannelID <= 0 {
		return domain.TemplateChannelLanguageMasterMap{}, fmt.Errorf("channel_id is required and must be positive")
	}
	if req.LanguageID <= 0 {
		return domain.TemplateChannelLanguageMasterMap{}, fmt.Errorf("language_id is required and must be positive")
	}
	return s.repo.AddTemplateChannelLanguage(ctx, req)
}

func (s *MappingService) GetTemplateChannelLanguages(ctx context.Context, limit, offset int) ([]domain.TemplateChannelLanguageMasterMap, error) {
	if limit <= 0 || limit > 100 {
		limit = 25
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetTemplateChannelLanguages(ctx, limit, offset)
}

func (s *MappingService) DeleteTemplateChannelLanguage(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id is required")
	}
	return s.repo.SoftDeleteTemplateChannelLanguage(ctx, id)
}
