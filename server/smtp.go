package server

import (
	"fmt"
	"net/smtp"
)

func SendEmail(formIn Input, user User) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: " + formIn.Subject + "\r\n" +
		"\r\n" + formIn.Message)

	auth := smtp.PlainAuth("", user.Email, user.Password, smtpHost)
	if user.Email != "" {
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, user.Email, formIn.To, message)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	fmt.Println("Email sent succesfully!")
	return true
}
