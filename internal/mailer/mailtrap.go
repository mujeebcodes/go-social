package mailer

import (
	"bytes"
	"errors"
	gomail "gopkg.in/mail.v2"
	"text/template"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}
func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) error {
	// Template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}
	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return err
	}
	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())
	message.AddAlternative("text/html", body.String())

	var dialer *gomail.Dialer
	if isSandbox {
		// Use Mailtrap Sandbox SMTP credentials
		dialer = gomail.NewDialer("smtp.mailtrap.io", 587, "ef3c89bf1767f2", "271faa657718e6")
	} else {
		// Use Mailtrap Production SMTP credentials
		dialer = gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", m.apiKey)
	}

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
