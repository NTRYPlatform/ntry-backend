package notary

import (
	"net/smtp"
)

type emailConf struct {
	from, pass, server string
}

func newEmail() *emailConf {
	return &emailConf{}
}

func (email *emailConf) sendEmail(recipient, msg string) (err error) {
	return smtp.SendMail(email.server+":25",
		smtp.PlainAuth("", email.from, email.pass, email.server),
		email.from, []string{recipient}, []byte(msg))
}

func (email *emailConf) ok() bool {
	if len(email.from) < 16 {
		return false
	}
	if len(email.pass) < 8 {
		return false
	}
	if len(email.server) < 16 {
		return false
	}
	return true
}

func verificationAccountMessage(sender, recipient, uuid, address string) string {
	return "From: " + sender + "\n" +
		"To: " + recipient + "\n" +
		"Subject: Confirm your email to join Notary Platform\n\n" +
		"Hello! Once you've verified your email address, " +
		"you'll be the newest member of the Notary Platform!\n\n" +
		"Please execute the mapper function of the contract @ " + address + " with the uid:\n\n\t" + uuid +
		`\n\nABI Interface: [{"constant":false,"inputs":[{"name":"secondary","type":"bytes16"}],"name":"mapAddress","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
}
