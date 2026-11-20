package controller

import (
	"database/sql"
	"encoding/json"
	"kaizen/usecase"
	"log"
	"net/http"
)

func UserSearchController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name")
	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, users := usecase.UserSearchUsecase(name, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
