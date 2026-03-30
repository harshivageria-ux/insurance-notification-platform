package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"probus-notification-system/internal/crypto"
	"probus-notification-system/internal/infrastructure/repository"
)

type Server struct {
	db        *pgxpool.Pool
	encryptor *crypto.Encryptor

	// Repositories
	languageRepo        *repository.LanguageRepository
	priorityRepo        *repository.PriorityRepository
	statusRepo          *repository.StatusRepository
	scheduleTypeRepo    *repository.ScheduleTypeRepository
	categoryRepo        *repository.CategoryRepository
	channelRepo         *repository.ChannelRepository
	channelProviderRepo *repository.ChannelProviderRepository
	providerSettingRepo *repository.ProviderSettingRepository
	templateGroupRepo   *repository.TemplateGroupRepository
	templateRepo        *repository.TemplateRepository
	routingRuleRepo     *repository.RoutingRuleRepository
}

func NewServer(db *pgxpool.Pool, encryptor *crypto.Encryptor) *Server {
	return &Server{
		db:                  db,
		encryptor:           encryptor,
		languageRepo:        repository.NewLanguageRepository(db),
		priorityRepo:        repository.NewPriorityRepository(db),
		statusRepo:          repository.NewStatusRepository(db),
		scheduleTypeRepo:    repository.NewScheduleTypeRepository(db),
		categoryRepo:        repository.NewCategoryRepository(db),
		channelRepo:         repository.NewChannelRepository(db),
		channelProviderRepo: repository.NewChannelProviderRepository(db),
		providerSettingRepo: repository.NewProviderSettingRepository(db),
		templateGroupRepo:   repository.NewTemplateGroupRepository(db),
		templateRepo:        repository.NewTemplateRepository(db),
		routingRuleRepo:     repository.NewRoutingRuleRepository(db),
	}
}

func (s *Server) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Language routes
	r.Route("/languages", func(r chi.Router) {
		r.Get("/", s.listLanguages)
		r.Post("/", s.createLanguage)
		r.Put("/{id}", s.updateLanguage)
		r.Delete("/{id}", s.deactivateLanguage)
	})

	// Priority routes
	r.Route("/priorities", func(r chi.Router) {
		r.Get("/", s.listPriorities)
		r.Post("/", s.createPriority)
		r.Put("/{id}", s.updatePriority)
		r.Delete("/{id}", s.deactivatePriority)
	})

	// Status routes
	r.Route("/statuses", func(r chi.Router) {
		r.Get("/", s.listStatuses)
		r.Post("/", s.createStatus)
		r.Put("/{id}", s.updateStatus)
		r.Delete("/{id}", s.deactivateStatus)
	})

	// Schedule Type routes
	r.Route("/schedule-types", func(r chi.Router) {
		r.Get("/", s.listScheduleTypes)
		r.Post("/", s.createScheduleType)
		r.Put("/{id}", s.updateScheduleType)
		r.Delete("/{id}", s.deactivateScheduleType)
	})

	// Category routes
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", s.listCategories)
		r.Post("/", s.createCategory)
		r.Put("/{id}", s.updateCategory)
		r.Delete("/{id}", s.deactivateCategory)
	})

	// Channel routes
	r.Route("/channels", func(r chi.Router) {
		r.Get("/", s.listChannels)
		r.Post("/", s.createChannel)
		r.Put("/{id}", s.updateChannel)
		r.Patch("/{id}/toggle", s.toggleChannel)
		r.Delete("/{id}", s.deactivateChannel)
	})

	// Channel Provider routes
	r.Route("/channel-providers", func(r chi.Router) {
		r.Get("/", s.listChannelProviders)
		r.Post("/", s.createChannelProvider)
		r.Put("/{id}", s.updateChannelProvider)
		r.Patch("/{id}/toggle", s.toggleChannelProvider)
		r.Delete("/{id}", s.deactivateChannelProvider)
	})

	// Provider Settings routes
	r.Route("/provider-settings", func(r chi.Router) {
		r.Get("/{provider_id}", s.listProviderSettings)
		r.Post("/", s.createProviderSetting)
		r.Put("/{id}", s.updateProviderSetting)
		r.Delete("/{id}", s.deactivateProviderSetting)
	})

	// Template Group routes
	r.Route("/template-groups", func(r chi.Router) {
		r.Get("/", s.listTemplateGroups)
		r.Post("/", s.createTemplateGroup)
		r.Put("/{id}", s.updateTemplateGroup)
		r.Delete("/{id}", s.deactivateTemplateGroup)
	})

	// Template routes
	r.Route("/templates", func(r chi.Router) {
		r.Get("/", s.listTemplates)
		r.Post("/", s.createTemplate)
		r.Put("/{id}", s.updateTemplate)
		r.Delete("/{id}", s.deactivateTemplate)
		r.Post("/{id}/preview", s.previewTemplate)
	})

	// Routing Rule routes
	r.Route("/routing-rules", func(r chi.Router) {
		r.Get("/", s.listRoutingRules)
		r.Post("/", s.createRoutingRule)
		r.Put("/{id}", s.updateRoutingRule)
		r.Patch("/{id}/toggle", s.toggleRoutingRule)
		r.Delete("/{id}", s.deactivateRoutingRule)
	})

	return r
}

// Helper function to respond with JSON
func respondJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// Helper function to respond with error
func respondError(w http.ResponseWriter, statusCode int, message string) {
	respondJSON(w, statusCode, map[string]string{"error": message})
}

// Helper function to get ID from URL parameters
func getIDParam(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.Atoi(idStr)
}

func getStringIDParam(r *http.Request) string {
	return chi.URLParam(r, "id")
}

// Helper function to get providerID from URL parameters
func getProviderIDParam(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "provider_id")
	return strconv.Atoi(idStr)
}

func getStringProviderIDParam(r *http.Request) string {
	return chi.URLParam(r, "provider_id")
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
