package auth

import (
	"log"
	"net/smtp"

	"github.com/ntryapp/auth/config"
)

func SendVerificationEmail(email, code string) {
	from, pass, server := config.GetEmailInfo()
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Confirm your email to join Notary Platform\n\n" +
		"Hello! Once you've verified your email address, " +
		"you'll be the newest member of the Notary Platform!\n\n" +
		"Please sign the following code with your ethereum key for validation:\n\n\t" + code

	err := smtp.SendMail(server+":25", smtp.PlainAuth("", from, pass, server), from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
