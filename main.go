package main

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

// Functions
func sendCustomMail(sender, password, smtpHost, smtpPort string, receivers []string) {

	customMail := template.Must(template.ParseFiles("./templates/confirmation.html"))

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

	// Parse templates
	tmpl = template.Must(template.ParseGlob("./templates/*"))

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

	fmt.Println("Server listening on 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
