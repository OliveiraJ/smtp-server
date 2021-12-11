package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Input struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

func Mail(res http.ResponseWriter, req *http.Request) {
	formIn := &Input{}
	req.ParseForm()
	formIn.From = req.Form.Get("from")
	formIn.To = append(formIn.To, req.Form.Get("To"))
	formIn.Subject = req.Form.Get("subject")
	formIn.Message = req.Form.Get("message")

	res.Header().Set("Content-Type", "html")
	tmp, _ := template.ParseFiles("static/email.html")
	tmp.Execute(res, nil)

}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Mail)
	log.Default().Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
