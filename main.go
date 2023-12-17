package main

/*

**************************** IMPORTANT ********************
Before commit, remove all sensitive data!!!
************************************** ********************

*/

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
)

// Structs
type Form struct {
	Email, Password string
	Error, Success  bool
}

// Variables
var tmpl *template.Template
var port string
var templateLocation string = "./templates/*"
var fs http.Handler

// Functions

// Serve CSS, JS files
func serveStaticFiles() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func sendCustomMail(sender, password, smtpHost, smtpPort string, receivers []string) {

	customMail := template.Must(template.ParseGlob(templateLocation))

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Mail confirmation! \n%s\n\n", mimeHeaders)))

	// Authentication
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	customMail.ExecuteTemplate(&body, "confirmation.html", "")

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, receivers, body.Bytes()); err != nil {
		log.Fatal(err)
		return
	}

}

func main() {

	serveStaticFiles()

	port = "3000"

	// Parse templates
	tmpl = template.Must(template.ParseGlob(templateLocation))

	// TESTS ***
	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.ExecuteTemplate(w, "index.html", "")
	// 	return
	// })

	// http.HandleFunc("/confirmation", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("\n confirmation!")
	// 	tmpl.ExecuteTemplate(w, "confirmation.html", "")
	// 	return
	// })
	// END TESTS ***

	// Handlers
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		form := Form{}

		// Check if user send the data using form
		switch r.Method {
		case "POST":

			form.Email = r.FormValue("email")
			form.Password = r.FormValue("password")

			// Check if user send empty or filled fields
			if form.Email == "" || form.Password == "" {
				form.Error = true
				tmpl.ExecuteTemplate(w, "index.html", form)
				return
			}

			// User sends correct fields
			form.Error = false
			form.Success = true

			// Send email confirmation
			sendCustomMail("sender_email", "sender_password", "smtp.gmail.com", "587", []string{r.FormValue("email")})
			tmpl.ExecuteTemplate(w, "index.html", form)

		case "GET":
			tmpl.ExecuteTemplate(w, "index.html", form)
			return
		}

	})

	fmt.Printf("Server listening on %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
