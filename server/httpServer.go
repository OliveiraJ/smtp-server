package server

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Input struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Mail(res http.ResponseWriter, req *http.Request) {
	tmp := template.Must(template.ParseFiles("static/email.html"))

	formIn := Input{
		From:    req.FormValue("from"),
		Subject: req.FormValue("subject"),
		Message: req.FormValue("message"),
		To:      strings.Split(req.FormValue("to"), ","),
	}

	log.Println(formIn)
	SendEmail(formIn)
	res.Header().Set("Content-Type", "html")
	tmp.Execute(res, struct{ Success bool }{false})

}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Mail)
	log.Default().Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
