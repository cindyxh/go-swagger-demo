package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/godemo/common"
	"github.com/godemo/model"
)

// UserHandler handles user operations
type UserHandler struct{}

// GetUser returns user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user := &model.User{ID: id, Name: "Mary Hather", Email: "mary@test.com"}
	common.RespondWithJSON(w, http.StatusOK, user)
}
