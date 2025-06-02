package cmd

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
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
	log.Printf("дошел1")

	err = h.DB.QueryRow("INSERT INTO person(name,email,password) VALUES ($1,$2,$3) RETURNING id", p.Name, p.Email, p.Password).Scan(&userID)
	if err != nil {
		log.Printf("func register3: insert err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.Printf("func register4: person not found")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusUnauthorized)
	}
	log.Printf("доешел2")
	sessionid, err := generateid()
	if err != nil {
		log.Printf("func register5: generate sessionid err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	expires := time.Now().Add(24 * time.Hour)

	log.Printf("дошел3")
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Path:     "/",
		MaxAge:   int(expires.Sub(time.Now()).Seconds()),
		Value:    sessionid,
		Secure:   false,
		HttpOnly: true,
	})
	_, err = h.DB.Exec("UPDATE person SET session_id = $1 WHERE  id = $2 ", sessionid, userID)
	if err != nil {
		log.Printf("func register6: update err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fogge := map[string]interface{}{
		"user_id": userID,
		"name":    p.Name,
		"email":   p.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(fogge)

}
