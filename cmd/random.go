package cmd

import (
	"context"
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

	// Создание контекста
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//

	var value string
	var val string
	var err error

	//Достать куки

	///

	err = h.DB.QueryRowContext(ctx, `SELECT qoutes,auth FROM qoute  ORDER BY RANDOM() LIMIT 1`).Scan(&value, &val)
	if err != nil {
		log.Printf("func ServeHttp:failed to make request %w", err)
	}
	time.Sleep(1 * time.Millisecond)

	///resp := fmt.Sprintf("Ваша цитата дня ", value)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp := map[string]string{
		"quote":  value,
		"author": val,
	}
	if err23 := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		log.Printf("func ServeHttp:failed to make response %w", err23)

	}

}
