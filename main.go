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
	fs := http.FileServer(http.Dir("./Fronted"))
	http.Handle("/", fs)
	http.Handle("/generate", fs)

	db := database.Connect()
	defer db.Close()

	h := &cmd.Handler{DB: db}

	http.HandleFunc("/generate/api", h.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
