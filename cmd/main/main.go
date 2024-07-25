package main

import (
	"database/sql"
	"fmt"
	"holmgrendev/dgp/internal/database"
	"holmgrendev/dgp/internal/router"
	"html/template"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

	// Create a new HTTP request Multiplexer
	mux := router.NewRouter()

	// Handle requests
	mux.HandleFunc("/", exampleHandler) // Example

	http.ListenAndServe(":3080", mux)

}

/* ------------ EXAMPLE CODE ------------ */
type Page struct {
	Title string
	Body  string
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {

	dbc := database.InitDatabase()

	db, _ = sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbc.Host, dbc.Database, dbc.User, dbc.Password))

	err := db.Ping()

	if err != nil {
		log.Printf("Error conecting to database:\n%s\n", err)
		fmt.Fprintf(w, "Error conecting to database:<br>%s", err)
	} else {

		// Define example page
		p := Page{Title: "DGP Template", Body: "Hello DGP"}
		t, err := template.ParseFiles("./templates/pages/default.html", "./templates/elements/div.template.html")

		if err != nil {
			log.Printf("Error parsing template:\n%s\n", err)
			fmt.Fprintf(w, "Error parsing template:<br>%s", err)
		}

		// Execute template
		err = t.Execute(w, p)

		if err != nil {
			log.Printf("Error Executing:\n%s\n", err)
			fmt.Fprintf(w, "Error Executing:<br>%s", err)
		}
	}

}

/* ---------- END EXAMPLE CODE ---------- */
