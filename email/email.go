package email

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func Email(nameTo, emailTo string) error {
	// Set up authentication information.

	smtpServer := "tzk601.nic.ua"
	auth := smtp.PlainAuth(
		"",
		"admin@micro.pp.ua",
		"atebHoc4",
		smtpServer,
	)

	from := mail.Address{"GGSMG-site", "admin@micro.pp.ua"}
	to := mail.Address{nameTo, emailTo}
	title := "GGSMG"

	body := "Your sitemap"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := "Here is your sitemap."
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":465",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	if err != nil {
		return err
	}
	return nil
}
