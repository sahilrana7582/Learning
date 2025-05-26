package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"time"

	"github.com/go-mail/mail"
)

//go:embed templates/*
var templateFS embed.FS

type Mailer struct {
	dialer *mail.Dialer
	sender string
}

func New(host string, port int, username, password, sender string) Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	return Mailer{
		dialer: dialer,
		sender: sender,
	}
}

func (m Mailer) Send(recipient, templateFile string, data interface{}) error {

	tmpl, err := template.New("email").ParseFS(templateFS, "templates/user_welcome.tmpl")
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	var subject bytes.Buffer
	if err := tmpl.ExecuteTemplate(&subject, "subject", data); err != nil {
		return fmt.Errorf("executing subject template: %w", err)
	}

	var plainBody bytes.Buffer
	if err := tmpl.ExecuteTemplate(&plainBody, "plainBody", data); err != nil {
		return fmt.Errorf("executing plainBody template: %w", err)
	}

	var htmlBody bytes.Buffer
	if err := tmpl.ExecuteTemplate(&htmlBody, "htmlBody", data); err != nil {
		return fmt.Errorf("executing htmlBody template: %w", err)
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	for i := 0; i < 3; i++ {

		err = m.dialer.DialAndSend(msg)
		if err == nil {
			fmt.Printf("Email sent successfully to %s\n", recipient)
			return nil
		}

		time.Sleep(500 * time.Millisecond)
	}

	return err
}
