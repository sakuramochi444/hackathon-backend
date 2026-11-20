package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"kaizen/model"
	"kaizen/usecase"
	"log"
	"net/http"
)

func UserRegisterController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, id := usecase.UserRegisterUsecase(r.Context(), req, db)
	if errors.Is(err, usecase.ErrInvalidUser) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(map[string]string{"id": id.String()})
	if err != nil {
		log.Printf("fail: json.Marshal, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
