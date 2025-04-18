package cmd

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var value string
	var val string

	err := h.DB.QueryRow(`SELECT qoutes,auth FROM qoute  ORDER BY RANDOM() LIMIT 1`).Scan(&value, &val)
	if err != nil {
		log.Printf("func ServeHttp:failed to make request %w", err)
	}
	time.Sleep(1 * time.Millisecond)

	///resp := fmt.Sprintf("Ваша цитата дня ", value)

	w.Header().Set("Content-Type", "aplication/json; charset=utf-8")

	resp := map[string]string{
		"qoutes": value,
		"auth":   val,
	}
	json.NewEncoder(w).Encode(resp)
}
