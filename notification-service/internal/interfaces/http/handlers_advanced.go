package http

import (
	"encoding/json"
	"net/http"

	"probus-notification-system/internal/domain/channel"
	cp "probus-notification-system/internal/domain/channel_provider"
	ps "probus-notification-system/internal/domain/provider_setting"
	rr "probus-notification-system/internal/domain/routing_rule"
	"probus-notification-system/internal/domain/template"
	tg "probus-notification-system/internal/domain/template_group"
)

// ============= CHANNEL HANDLERS =============

func (s *Server) listChannels(w http.ResponseWriter, r *http.Request) {
	channels, err := s.channelRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch channels")
		return
	}
	respondJSON(w, http.StatusOK, channels)
}

func (s *Server) createChannel(w http.ResponseWriter, r *http.Request) {
	var req channel.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	c, err := s.channelRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create channel")
		return
	}
	respondJSON(w, http.StatusCreated, c)
}

func (s *Server) updateChannel(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req channel.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ID = id

	c, err := s.channelRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update channel")
		return
	}
	respondJSON(w, http.StatusOK, c)
}

func (s *Server) toggleChannel(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req channel.ToggleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	c, err := s.channelRepo.Toggle(r.Context(), id, req.IsActive)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to toggle channel")
		return
	}
	respondJSON(w, http.StatusOK, c)
}

func (s *Server) deactivateChannel(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.channelRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete channel")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Channel deleted successfully"})
}

// ============= CHANNEL PROVIDER HANDLERS =============

func (s *Server) listChannelProviders(w http.ResponseWriter, r *http.Request) {
	providers, err := s.channelProviderRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch channel providers")
		return
	}
	respondJSON(w, http.StatusOK, providers)
}

func (s *Server) createChannelProvider(w http.ResponseWriter, r *http.Request) {
	var req cp.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p, err := s.channelProviderRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create channel provider")
		return
	}
	respondJSON(w, http.StatusCreated, p)
}

func (s *Server) updateChannelProvider(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req cp.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ID = id

	p, err := s.channelProviderRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update channel provider")
		return
	}
	respondJSON(w, http.StatusOK, p)
}

func (s *Server) toggleChannelProvider(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req cp.ToggleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p, err := s.channelProviderRepo.Toggle(r.Context(), id, req.IsActive)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to toggle channel provider")
		return
	}
	respondJSON(w, http.StatusOK, p)
}

func (s *Server) deactivateChannelProvider(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.channelProviderRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete channel provider")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Channel provider deleted successfully"})
}

// ============= PROVIDER SETTINGS HANDLERS =============

func (s *Server) listProviderSettings(w http.ResponseWriter, r *http.Request) {
	providerID, err := getProviderIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid provider ID")
		return
	}

	settings, err := s.providerSettingRepo.GetByProviderID(r.Context(), providerID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch provider settings")
		return
	}
	respondJSON(w, http.StatusOK, settings)
}

func (s *Server) createProviderSetting(w http.ResponseWriter, r *http.Request) {
	var req ps.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	setting, err := s.providerSettingRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create provider setting")
		return
	}
	respondJSON(w, http.StatusCreated, setting)
}

func (s *Server) updateProviderSetting(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req ps.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ID = id

	setting, err := s.providerSettingRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update provider setting")
		return
	}
	respondJSON(w, http.StatusOK, setting)
}

func (s *Server) deactivateProviderSetting(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.providerSettingRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete provider setting")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Provider setting deleted successfully"})
}

// ============= TEMPLATE GROUP HANDLERS =============

func (s *Server) listTemplateGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := s.templateGroupRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch template groups")
		return
	}
	respondJSON(w, http.StatusOK, groups)
}

func (s *Server) createTemplateGroup(w http.ResponseWriter, r *http.Request) {
	var req tg.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	g, err := s.templateGroupRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create template group")
		return
	}
	respondJSON(w, http.StatusCreated, g)
}

func (s *Server) updateTemplateGroup(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req tg.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ID = id

	g, err := s.templateGroupRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update template group")
		return
	}
	respondJSON(w, http.StatusOK, g)
}

func (s *Server) deactivateTemplateGroup(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.templateGroupRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete template group")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Template group deleted successfully"})
}

// ============= TEMPLATE HANDLERS =============

func (s *Server) listTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := s.templateRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch templates")
		return
	}
	respondJSON(w, http.StatusOK, templates)
}

func (s *Server) createTemplate(w http.ResponseWriter, r *http.Request) {
	var req template.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	t, err := s.templateRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create template")
		return
	}
	respondJSON(w, http.StatusCreated, t)
}

func (s *Server) updateTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req template.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ID = id

	t, err := s.templateRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update template")
		return
	}
	respondJSON(w, http.StatusOK, t)
}

func (s *Server) previewTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req template.PreviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	t, err := s.templateRepo.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Template not found")
		return
	}

	// Simple template rendering (can be extended with more sophisticated templating)
	respondJSON(w, http.StatusOK, template.PreviewResponse{
		RenderedContent: t.Content,
	})
}

func (s *Server) deactivateTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.templateRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete template")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Template deleted successfully"})
}

// ============= ROUTING RULE HANDLERS =============

func (s *Server) listRoutingRules(w http.ResponseWriter, r *http.Request) {
	rules, err := s.routingRuleRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch routing rules")
		return
	}
	respondJSON(w, http.StatusOK, rules)
}

func (s *Server) createRoutingRule(w http.ResponseWriter, r *http.Request) {
	var req rr.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	rule, err := s.routingRuleRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create routing rule")
		return
	}
	respondJSON(w, http.StatusCreated, rule)
}

func (s *Server) updateRoutingRule(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req rr.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ID = id

	rule, err := s.routingRuleRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update routing rule")
		return
	}
	respondJSON(w, http.StatusOK, rule)
}

func (s *Server) toggleRoutingRule(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req rr.ToggleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	rule, err := s.routingRuleRepo.Toggle(r.Context(), id, req.IsActive)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to toggle routing rule")
		return
	}
	respondJSON(w, http.StatusOK, rule)
}

func (s *Server) deactivateRoutingRule(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.routingRuleRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete routing rule")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Routing rule deleted successfully"})
}
