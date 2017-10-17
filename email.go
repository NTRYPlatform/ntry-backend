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
	return "From: " + sender + "\r\n" +
		"To: " + recipient + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"Subject: Confirm your email to join Notary Platform\n\n" +
		"Hello!<br><br>Once you've verified your email address, " +
		"you'll be the newest member of the Notary Platform!<br><br>\n\n" +
		"Please execute the mapper function of the contract @ " + address + " with the uid:\n\n\t" + uuid + "<br><br>\n\n" +
		`ABI Interface: [{"constant":false,"inputs":[{"name":"secondary","type":"bytes16"}],"name":"signup","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]` +
		"\n\n<br><br>For more elaborate instructions, click <a href=\"https://medium.com/@sash87/4433f6c1bf7d\">here</a>." +
		"<br><br>Cheers!<br>Notary Team"
}

func changePasswordMessage(sender, recipient, time, password string) string {
	return "From: " + sender + "\r\n" +
		"To: " + recipient + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"Subject: Notary Platform: Forgot your password?\n\n" +
		"Hello!<br><br>We received a request to change your password at " + time + ". " +
		"<br><br>If that wasn't you, we're sorry for the trouble. Please ignore this email and keep using Notary Platform!\n\n" +
		"\n<br><br>If it was you... um, ever heard of password managers?" +
		"\n\n<br><br>But you're in luck. You can now use the following temporary password to change your current one:\n\n\t" + password +
		"\n\n<br><br>Cheers!<br>\nNotary Team"
}
