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
		"Please execute the mapper function of the contract @ " + address + " with the uid:\n\n\t" + uuid + "\n\n" +
		`ABI Interface: Latest ABI to be sent in email:[{"constant":false,"inputs":[{"name":"secondary","type":"bytes16"}],"name":"signup","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]` +
		"\n\n For more elaborate instructions, click <a href=\"https://medium.com/@sash87/4433f6c1bf7d\">here</a>."
}

func changePasswordMessage(sender, recipient, time, password string) string {
	return "From: " + sender + "\n" +
		"To: " + recipient + "\n" +
		"Subject: Notary Platform: Forgot your password?\n\n" +
		"Hello! We received a request to change your password at " + time + ". " +
		"If that wasn't you, we're sorry for the trouble. Please ignore this email and keep using Notary Platform!\n\n" +
		"\nIf it was you... um, ever heard of password managers?" +
		"\n\nBut you're in luck. You can now use the following temporary password to change your current one:\n\n\t" + password +
		"\n\nCheers!\nNotary Team"
}
