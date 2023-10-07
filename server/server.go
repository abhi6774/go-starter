package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func RunSimpleHelloWorldServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":8080", nil)
}

func RunWithTemplates() {
	type Page struct {
		Title string
		Body  string
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define the data to be rendered in the template
		data := Page{
			Title: "Welcome to my website",
			Body:  "This is the body of my website",
		}

		// Parse the template file
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Render the template with the data
		err = tmpl.ExecuteTemplate(w, "content", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)

}

func RunStaticServer() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
