package transport

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"uuid"

	"github.com/NeoFSociety/cleanArch/internal/domain"
	"github.com/NeoFSociety/cleanArch/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		http.Error(w, "username is empty", http.StatusBadRequest)
		return
	}

	user, err := h.svc.Register(r.Context(), req.Username)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserResponse{
		UUID:     user.UUID,
		Username: user.Username,
	})

	log.Println("user created")
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("uuid")
	if uid == "" {
		http.Error(w, "uuid is empty", http.StatusBadRequest)
		return
	}

	if _, err := uuid.Parse(uid); err != nil {
		http.Error(w, "invalid uuid format", http.StatusBadRequest)
		return
	}

	user, err := h.svc.GetByUID(r.Context(), uid)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserResponse{
		UUID:     user.UUID,
		Username: user.Username,
	})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("uuid")
	if uid == "" {
		http.Error(w, "uuid is empty", http.StatusBadRequest)
		return
	}
	if _, err := uuid.Parse(uid); err != nil {
		http.Error(w, "invalid uuid format", http.StatusBadRequest)
		return
	}

	if err := h.svc.Delete(r.Context(), uid); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
