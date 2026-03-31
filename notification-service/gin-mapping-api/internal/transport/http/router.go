package http

import (
	"github.com/gin-gonic/gin"
	"probus-notification-system/gin-mapping-api/internal/service"
)

func SetupRouter(mappingService *service.MappingService) *gin.Engine {
	r := gin.Default()
	h := NewMappingHandler(mappingService)

	api := r.Group("/api")
	{
		api.POST("/notification-category-channel", h.AddCategoryChannel)
		api.GET("/notification-category-channel", h.GetCategoryChannels)
		api.DELETE("/notification-category-channel/:id", h.DeleteCategoryChannel)

		api.POST("/channel-provider-map", h.AddChannelProvider)
		api.GET("/channel-provider-map", h.GetChannelProviders)
		api.DELETE("/channel-provider-map/:id", h.DeleteChannelProvider)

		api.POST("/template-channel-language-map", h.AddTemplateChannelLanguage)
		api.GET("/template-channel-language-map", h.GetTemplateChannelLanguages)
		api.DELETE("/template-channel-language-map/:id", h.DeleteTemplateChannelLanguage)
	}

	return r
}
