package server

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Representa os dados a serem enviados no email, como o slice de emails de destino, o assunto e o corpo da mensagem em si.
type Input struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

// Reprsenta os dados do usuário ou remetente do email, email e senha que serão usados no processo de autenticação perante o google
// e envio do email.
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Embed é resposável por inserir os arquivos da pasta static dentro do arquivo binário gerado ao se compilar o programa.
//go:embed static/*
var F embed.FS

// Instancia do tipo *template.Template que guardará os arquivos de template.
var tpl *template.Template
var user User

// Inicia a veriavel de template passando pra ela a variavel F do tipo embed.FS e o caminho para o diretório que possui os arquivos de template.
func InitTemplate() {
	tpl = template.Must(template.ParseFS(F, "static/*.gohtml"))
}

// Função responsável pelo login, chamando a função InitTemplate e em seguida retornando como resposta a página de login
func Login(res http.ResponseWriter, req *http.Request) {
	InitTemplate()
	tpl.ExecuteTemplate(res, "login.gohtml", nil)

}

// Responsávle por criar o corpo do email, recebendo os dados de login e verificando se o método da requisição é do tipo POST, caso
// Caso contrário ela redireciona o usuário de volta para a página de login.
// Verifica também se algum dos dados de login são vazios, mantendo o usuário logado, mesmo ao mudar de rota, sevindo de uma alternativa
// a serviços de perssitência e autenticação, que não são o foco dessa aplicação. Por fim ela retorna ao usuário a página de email
// aguardando o preenchimenmto do mesmo.
func Mail(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	res.Header().Set("Content-Type", "html")

	if req.FormValue("username") != "" && req.FormValue("password") != "" {
		user.Email = req.FormValue("username")
		user.Password = req.FormValue("password")

	}

	tpl.ExecuteTemplate(res, "sendemail.gohtml", user)

}

// Responsável por de fato enviar o email, recebendo os dados do forumlário da página de email, alimentando as structs com eles e
// chamando o servidor SMTP de modo a de fato enviar o email, retornando ao usuário uma página condizente com o resultado da operação
func process(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	formIn := Input{
		Subject: req.FormValue("subject"),
		Message: req.FormValue("message"),
		To:      strings.Split(req.FormValue("to"), ","),
	}
	log.Println(formIn)

	Isdone := SendEmail(formIn, user)

	tpl.ExecuteTemplate(res, "processor.gohtml", Isdone)
}

// Gerencia as rotas e dá início de fato ao servidor http, ouvindo na porta 3000.
func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Login)
	router.HandleFunc("/sendmail", Mail)
	router.HandleFunc("/processing", process)

	log.Default().Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
