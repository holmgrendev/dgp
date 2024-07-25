package main

import (
	"database/sql"
	"fmt"
	"holmgrendev/dgp/internal/router"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

//var db *sql.DB

var dbc = initDatabase()

type DatabaseCredentials struct {
	Host     string
	Database string
	User     string
	Password string
}

func initDatabase() DatabaseCredentials {
	// Set default credentials
	dbc := DatabaseCredentials{Host: "localhost", Database: os.Getenv("POSTGRES_DB"), User: "postgres", Password: os.Getenv("POSTGRES_PASSWORD")}

	// Check user
	if os.Getenv("POSTGRES_USER") != "" {
		dbc.User = os.Getenv("POSTGRES_USER")
	}

	// Check database
	if dbc.Database == "" {
		dbc.Database = dbc.User
	}

	// Check host
	if os.Getenv("POSTGRES_HOST") != "" {
		dbc.Host = os.Getenv("POSTGRES_HOST")
	}

	return dbc
}

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

	db, _ := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbc.Host, dbc.Database, dbc.User, dbc.Password))

	err := db.Ping()

	if err != nil {
		log.Printf("Error conecting to database:\n%s\n", err)
		fmt.Fprintf(w, "Error conecting to database:<br>%s", err)
	} else {

		// Define example page
		p := Page{Title: "DOGOW Template", Body: "Hello DOGOW"}
		t, err := template.ParseFiles("./templates/pages/default.html")

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
