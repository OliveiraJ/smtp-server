package server

import (
	"fmt"
	"net/smtp"
)

func SendEmail() {
	from := "jordansilva102@gmail.com"
	password := "j81073880"

	to := []string{
		"jordansvaoliveira@gmail.com",
		"JordanOli@protonmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("To: jordansvaoliveira@gmail.com\r\n" + "Subject: Email de teste\r\n" +
		"\r\n" + "this is a test email message")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent succesfully!")
}
