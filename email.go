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
		"Hello!<br><br>Thank you for downloading the Notary Platform Beta application for testing purposes.<br><br>" +
		"Please follow the guide <a href=\"https://medium.com/@sash87/4433f6c1bf7d\">here</a> and use the data below for a successful sign up!<br><br>" +
		"<b>Contract Address:</b> " + address +
		"<br><br><b>Your User ID (uid):</b> " + uuid + "<br><br>\n\n" +
		`<b>ABI Interface:</b> [{"constant":false,"inputs":[{"name":"secondary","type":"bytes16"}],"name":"signup","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]` +
		"<br><br>If you have any trouble with the app, please contact our team on <a href=\"https://t.me/joinchat/F2wLskHLoA8UUhb3Iomt_g\">telegram</a>." +
		"<br><br>Happy testing!" +
		"<br><br>Cheers!<br>Notary Platform Team<br>"
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
		"\n\n<br><br>But you're in luck. You can now use the following code to change your current one:\n\n\t" + password +
		"<br><br>If you have any trouble with the app, please contact our team on <a href=\"https://t.me/joinchat/F2wLskHLoA8UUhb3Iomt_g\">telegram</a>." +
		"\n\n<br><br>Cheers!<br>\nNotary Team"
}
