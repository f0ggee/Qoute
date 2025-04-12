package cmd

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Handle struct {
	DB *sql.DB
}

type Person struct {
	Id       int64
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handle) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}

	var p Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}

	var email string

	err := h.DB.QueryRow("SELECT email FROM person WHERE email=$1", p.Email).Scan(&email)
	if err != sql.ErrNoRows && err != nil {
		log.Println("Func Rgister1:failed to connetct", err)
		return
	}

	if err != nil {
		log.Println("Func Rgister2:failed to connetct", err)

	}

	f, err := h.DB.Exec("INSERT INTO person (name,email,password)  VALUES ($1, $2, $3)", p.Name, p.Email, p.Password)
	if err != nil {
		log.Println("Func Rgister4:failed to connetct", err)
		return
	}
	log.Println(f)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
