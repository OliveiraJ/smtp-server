package server

import (
	"fmt"
	"net/smtp"
)

func SendEmail(formIn Input) {
	//from := "jordansilva102@gmail.com"
	password := "j81073880"

	// to := []string{
	// 	"jordansvaoliveira@gmail.com",
	// 	"JordanOli@protonmail.com",
	// }

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: " + formIn.Subject + "\r\n" +
		"\r\n" + formIn.Message)

	auth := smtp.PlainAuth("", formIn.From, password, smtpHost)
	if formIn.From != "" {
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, formIn.From, formIn.To, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Email sent succesfully!")
}
