package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"probus-notification-system/internal/domain/category"
	"probus-notification-system/internal/domain/language"
	"probus-notification-system/internal/domain/priority"
	st "probus-notification-system/internal/domain/schedule_type"
	"probus-notification-system/internal/domain/status"
)

// ============= LANGUAGE HANDLERS =============

func (s *Server) listLanguages(w http.ResponseWriter, r *http.Request) {
	languages, err := s.languageRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch languages")
		return
	}
	respondJSON(w, http.StatusOK, languages)
}

func (s *Server) createLanguage(w http.ResponseWriter, r *http.Request) {
	var req language.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		respondLanguageValidationError(w, err)
		return
	}

	lang, err := s.languageRepo.Create(r.Context(), req)
	if err != nil {
		respondLanguageRepositoryError(w, err, "Failed to create language")
		return
	}
	respondJSON(w, http.StatusCreated, lang)
}

func (s *Server) updateLanguage(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil || id <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req language.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req = req.Normalize()
	req.ID = int64(id)
	if err := req.Validate(); err != nil {
		respondLanguageValidationError(w, err)
		return
	}

	lang, err := s.languageRepo.Update(r.Context(), req)
	if err != nil {
		respondLanguageRepositoryError(w, err, "Failed to update language")
		return
	}
	respondJSON(w, http.StatusOK, lang)
}

func (s *Server) deactivateLanguage(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil || id <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.languageRepo.Delete(r.Context(), int64(id)); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete language")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Language deleted successfully"})
}

// ============= PRIORITY HANDLERS =============

func (s *Server) listPriorities(w http.ResponseWriter, r *http.Request) {
	priorities, err := s.priorityRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch priorities")
		return
	}
	respondJSON(w, http.StatusOK, priorities)
}

func (s *Server) createPriority(w http.ResponseWriter, r *http.Request) {
	var req priority.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// Normalize (trim/uppercase) to keep code comparisons consistent.
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	p, err := s.priorityRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create priority")
		return
	}
	respondJSON(w, http.StatusCreated, p)
}

func (s *Server) updatePriority(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req priority.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.PriorityID = int16(id)
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	p, err := s.priorityRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update priority")
		return
	}
	respondJSON(w, http.StatusOK, p)
}

func (s *Server) deactivatePriority(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.priorityRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete priority")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Priority deleted successfully"})
}

// ============= STATUS HANDLERS =============

func (s *Server) listStatuses(w http.ResponseWriter, r *http.Request) {
	statuses, err := s.statusRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch statuses")
		return
	}
	respondJSON(w, http.StatusOK, statuses)
}

func (s *Server) createStatus(w http.ResponseWriter, r *http.Request) {
	var req status.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	st, err := s.statusRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create status")
		return
	}
	respondJSON(w, http.StatusCreated, st)
}

func (s *Server) updateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req status.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req = req.Normalize()
	req.StatusID = int16(id)
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	st, err := s.statusRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update status")
		return
	}
	respondJSON(w, http.StatusOK, st)
}

func (s *Server) deactivateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.statusRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete status")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Status deleted successfully"})
}

// ============= SCHEDULE TYPE HANDLERS =============

func (s *Server) listScheduleTypes(w http.ResponseWriter, r *http.Request) {
	types, err := s.scheduleTypeRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch schedule types")
		return
	}
	respondJSON(w, http.StatusOK, types)
}

func (s *Server) createScheduleType(w http.ResponseWriter, r *http.Request) {
	var req st.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	t, err := s.scheduleTypeRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create schedule type")
		return
	}
	respondJSON(w, http.StatusCreated, t)
}

func (s *Server) updateScheduleType(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req st.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.ScheduleTypeID = int16(id)
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	t, err := s.scheduleTypeRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update schedule type")
		return
	}
	respondJSON(w, http.StatusOK, t)
}

func (s *Server) deactivateScheduleType(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.scheduleTypeRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete schedule type")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Schedule type deleted successfully"})
}

// ============= CATEGORY HANDLERS =============

func (s *Server) listCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.categoryRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}
	respondJSON(w, http.StatusOK, categories)
}

func (s *Server) createCategory(w http.ResponseWriter, r *http.Request) {
	var req category.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req = req.Normalize()
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	c, err := s.categoryRepo.Create(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create category")
		return
	}
	respondJSON(w, http.StatusCreated, c)
}

func (s *Server) updateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil || id <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req category.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req = req.Normalize()
	req.ID = id
	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	c, err := s.categoryRepo.Update(r.Context(), req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update category")
		return
	}
	respondJSON(w, http.StatusOK, c)
}

func (s *Server) deactivateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getIDParam(r)
	if err != nil || id <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := s.categoryRepo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete category")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}

func respondLanguageValidationError(w http.ResponseWriter, err error) {
	respondError(w, http.StatusBadRequest, err.Error())
}

func respondValidationError(w http.ResponseWriter, err error) {
	respondError(w, http.StatusBadRequest, err.Error())
}

func respondLanguageRepositoryError(w http.ResponseWriter, err error, fallback string) {
	switch {
	case errors.Is(err, language.ErrDuplicateName), errors.Is(err, language.ErrDuplicateCode):
		respondError(w, http.StatusConflict, err.Error())
	case errors.Is(err, language.ErrLanguageNotFound):
		respondError(w, http.StatusNotFound, err.Error())
	default:
		respondError(w, http.StatusInternalServerError, fmt.Sprintf("%s: %v", fallback, err))
	}
}
