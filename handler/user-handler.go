package handler

import (
	"net/http"
	"strconv"

	"github.com/godemo/common"
	"github.com/godemo/model"
	"github.com/gorilla/mux"
)

// UserHandler handles user operations
type UserHandler struct{}

// GetUser returns user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user := &model.User{ID: id, Name: "Mary Hather", Email: "mary@test.com"}
	common.RespondWithJSON(w, http.StatusOK, user)
}
