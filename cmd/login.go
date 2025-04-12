package cmd

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type LoginCmd struct {
	DB *sql.DB
}

type Persone struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *LoginCmd) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var p *Persone

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("Func Rgister1:failed to connetct", err)
	}

	var email string
	var passworde string
	err = h.DB.QueryRow("SELECT email,password FROM person WHERE email=$1 AND  password=$2", p.Email, p.Password).Scan(&email, &passworde)
	if err == sql.ErrNoRows {
		log.Println("Func Rgister2:email not found")
		return
	} else if err != nil {
		log.Println("Func Rgister3:failed to query", err)
		return
	}

	if p.Password != passworde {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		log.Println("Func Rgister3:passwords do not match")
		return
	} else {
		log.Println("Func Rgister3:success")
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("Func Rgister3:email:", email)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})

}
