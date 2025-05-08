package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
		log.Println("Func Login1:failed to connetct", err)
	}
	defer r.Body.Close()

	var email string
	var password string
	var ide int

	login := h.DB.QueryRow(`SELECT id, email, password FROM person WHERE email = $1 AND password = $2`, p.Email, p.Password).Scan(&ide, &email, &password)
	if login != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("Func login2:failed to connetct", err)
		return
	}
	if login == sql.ErrNoRows {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		log.Println("Func login3:failed to connetct", err)
		return
	}

	if password != p.Password {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		log.Println("Func login4:failed to connetct", err)
		return
	}

	log.Printf("id user", ide)

	cookie, err := generateid()
	if err != nil {
		log.Println("Func login5:failed to connetct", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	expressionist := time.Now().Add(time.Hour * 24)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Path:     "/profile/api",
		Value:    cookie,
		HttpOnly: true,
		Secure:   false,
		Expires:  expressionist,
	})

	log.Println("Func login:куки установлено")
	_, err = h.DB.Exec("UPDATE person SET session_id = $1 WHERE id = $2", cookie, ide)
	if err != nil {
		log.Println("Func login6:failed to connetct", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id": "%s"}`, cookie)
	log.Println("Func Rgister5:failed to connetct", err)
}
