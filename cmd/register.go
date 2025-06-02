package cmd

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type Handle struct {
	DB *sql.DB
}

func generateid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err

	}
	return hex.EncodeToString(b), nil
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
		log.Printf("func register1: json decode err: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}

	var exists bool

	err := h.DB.QueryRow("SELECT EXISTS (select 1 FROM person WHERE email=$1)", p.Email).Scan(&exists)
	if err != nil {
		log.Printf("func register1: query err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if exists {
		log.Printf("func register2: person already exists")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var userID int64
	var NameFromBD string

	err = h.DB.QueryRow("INSERT INTO person(name,email,password) VALUES ($1,$2,$3) RETURNING (id)", p.Name, p.Email, p.Password).Scan(&userID)
	if err != nil {
		log.Printf("func register3: insert err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.Printf("func register4: person not found")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusUnauthorized)
	}
	cookie, err := generateid()
	if err != nil {
		log.Printf("func register5: generate sessionid err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	expires := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    cookie,
		Path:     "/profile/api",
		HttpOnly: true,
		Secure:   false,
		Expires:  expires,
	})

	_, err = h.DB.Exec("UPDATE person SET session_id = $1 WHERE  id = $2 ", cookie, userID)
	if err != nil {
		log.Printf("func register6: update err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"id":   userID,
		"name": NameFromBD,
	}
	slog.Info("func register6: response ", response)

	json.NewEncoder(w).Encode(response)

}
