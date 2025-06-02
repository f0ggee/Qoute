package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
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
	var err error

	var p *Persone

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("Func Login1:failed to connetct", err)
	}
	defer r.Body.Close()

	///Создание контекса
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var email string
	var password string
	var ide int

	err = h.DB.QueryRowContext(ctx, `SELECT id, email, password FROM person WHERE email = $1 AND password = $2`, p.Email, p.Password).Scan(&ide, &email, &password)
	switch {
	case err == context.DeadlineExceeded:
		slog.Error("Func Login1:deadline exceeded")
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return

	case err == sql.ErrNoRows:
		slog.Error("Func Login1:no rows")
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return

	case err != nil:
		slog.Error("Func Login1:failed to make request %w", err)
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
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
		Value:    cookie,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  expressionist,
	})

	log.Println("Func login:куки установлено")
	slog.Info(cookie)
	_, err = h.DB.ExecContext(ctx, "UPDATE person SET session_id = $1  WHERE id = $2", cookie, ide)
	switch {
	case err == context.DeadlineExceeded:
		slog.Error("Func login5:deadline exceeded")
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	case err == sql.ErrNoRows:
		slog.Error("Func login5:no rows")
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	case err != nil:
		slog.Error("Func login5:failed to make request %w", err)
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"cookie": cookie,
	}

	json.NewEncoder(w).Encode(resp)
}
