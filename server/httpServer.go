package server

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Input struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var tpl *template.Template
var user User

func init() {
	tpl = template.Must(template.ParseGlob("static/*.gohtml"))
}

func Login(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "login.gohtml", nil)

}

func Mail(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	res.Header().Set("Content-Type", "html")

	user.Email = req.FormValue("username")
	user.Password = req.FormValue("password")

	tpl.ExecuteTemplate(res, "email.gohtml", user)

}

func process(res http.ResponseWriter, req *http.Request) {
	formIn := Input{
		Subject: req.FormValue("subject"),
		Message: req.FormValue("message"),
		To:      strings.Split(req.FormValue("to"), ","),
	}
	log.Println(user)
	log.Println(formIn)

	Isdone := SendEmail(formIn, user)
	//data := []interface{}{Isdone, user}
	tpl.ExecuteTemplate(res, "processor.gohtml", Isdone)
}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Login)
	router.HandleFunc("/sendmail", Mail)
	router.HandleFunc("/processing", process)

	log.Default().Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
