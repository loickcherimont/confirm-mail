package handlers

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
	"os"
)

// *** Structs ***
type Form struct {
	Email, Password string
	Error, Success  bool
}

// *** Variables ***

// Declare multiple variables?
var tmpl *template.Template
var port string
var templateLocation string = "./templates/*"

// *** Functions ***
func sendCustomMail(sender, password, smtpHost, smtpPort string, receivers []string) {

	customMail := template.Must(template.ParseGlob(templateLocation))

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Mail confirmation! \n%s\n\n", mimeHeaders)))

	// Create a /auth ???
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	customMail.ExecuteTemplate(&body, "confirmation.html", "")

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, receivers, body.Bytes()); err != nil {
		log.Fatal(err)
		return
	}

}

func ServeStaticFiles() { // to serve CSS files
	fs := http.FileServer(http.Dir("../../styles"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

func Signup(w http.ResponseWriter, r *http.Request) {

	ServeStaticFiles()
	tmpl = template.Must(template.ParseGlob(templateLocation))

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
		sendCustomMail(os.Getenv("SENDER_ADDRESS"), os.Getenv("SENDER_PASSWORD"), "smtp.gmail.com", "587", []string{r.FormValue("email")})
		tmpl.ExecuteTemplate(w, "index.html", form)

	case "GET":
		tmpl.ExecuteTemplate(w, "index.html", form)
		return
	}

}
