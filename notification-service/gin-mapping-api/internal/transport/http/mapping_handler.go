package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"log"
	"probus-notification-system/gin-mapping-api/internal/domain"
	"probus-notification-system/gin-mapping-api/internal/service"
)

type MappingHandler struct {
	service *service.MappingService
}

func NewMappingHandler(service *service.MappingService) *MappingHandler {
	return &MappingHandler{service: service}
}

func (h *MappingHandler) AddCategoryChannel(c *gin.Context) {
	var req domain.CreateNotificationCategoryChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("AddCategoryChannel invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	rec, err := h.service.AddCategoryChannel(c.Request.Context(), req)
	if err != nil {
		log.Printf("AddCategoryChannel error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rec)
}

func (h *MappingHandler) GetCategoryChannels(c *gin.Context) {
	limit, offset := parsePagination(c)
	list, err := h.service.GetCategoryChannels(c.Request.Context(), limit, offset)
	if err != nil {
		log.Printf("GetCategoryChannels error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch mappings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": list, "limit": limit, "offset": offset})
}

func (h *MappingHandler) DeleteCategoryChannel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.DeleteCategoryChannel(c.Request.Context(), id); err != nil {
		log.Printf("DeleteCategoryChannel error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MappingHandler) AddChannelProvider(c *gin.Context) {
	var req domain.CreateChannelProviderMasterMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	rec, err := h.service.AddChannelProvider(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rec)
}

func (h *MappingHandler) GetChannelProviders(c *gin.Context) {
	limit, offset := parsePagination(c)
	items, err := h.service.GetChannelProviders(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch mappings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items, "limit": limit, "offset": offset})
}

func (h *MappingHandler) DeleteChannelProvider(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.DeleteChannelProvider(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MappingHandler) AddTemplateChannelLanguage(c *gin.Context) {
	var req domain.CreateTemplateChannelLanguageMasterMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	rec, err := h.service.AddTemplateChannelLanguage(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rec)
}

func (h *MappingHandler) GetTemplateChannelLanguages(c *gin.Context) {
	limit, offset := parsePagination(c)
	items, err := h.service.GetTemplateChannelLanguages(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch mappings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items, "limit": limit, "offset": offset})
}

func (h *MappingHandler) DeleteTemplateChannelLanguage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.DeleteTemplateChannelLanguage(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func parsePagination(c *gin.Context) (int, int) {
	limit := 25
	offset := 0
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 && v <= 100 {
			limit = v
		}
	}
	if o := c.Query("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil && v >= 0 {
			offset = v
		}
	}
	return limit, offset
}
