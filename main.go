package main

import (
	"log"
	"net/http"

	"qoutes/cmd"
	"qoutes/iternal/database"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//front

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./Fronted/index.html")
			return
		}
		http.ServeFile(w, r, "./Fronted"+r.URL.Path)
	})

	db := database.Connect()
	defer db.Close()

	h := &cmd.Handler{DB: db}
	p := &cmd.Handle{DB: db}
	e := &cmd.LoginCmd{DB: db}

	http.HandleFunc("/login/api", e.Execute)
	http.HandleFunc("/register/api", p.Register)
	http.HandleFunc("/generate/api", h.ServeHTTP)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
