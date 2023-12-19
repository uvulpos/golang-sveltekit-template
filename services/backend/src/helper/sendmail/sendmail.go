package sendmail

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Dialer *gomail.Dialer

	Host     string
	Username string
	Email    string
}

func NewMailer(host string, port int, emailaddr, username, password string) *Mailer {
	dialer := gomail.NewDialer(host, port, username, password)
	return &Mailer{
		Dialer:   dialer,
		Host:     fmt.Sprintf("%s:%d", host, port),
		Username: username,
		Email:    emailaddr,
	}
}

type Mail struct {
	Mailer *Mailer

	Sender         string
	Receipiants    []string
	CcReceipiants  []string
	BccReceipiants []string

	Subject string
	Body    string
}

func (m *Mailer) NewMail() *Mail {
	return &Mail{
		Mailer: m,
		Sender: m.Email,
	}
}

func (m *Mail) SetSender(sender string) {
	m.Sender = sender
}

func (m *Mail) SetReceipiants(receipiants []string) {
	m.Receipiants = receipiants
}

func (m *Mail) SetCcReceipiants(receipiants []string) {
	m.CcReceipiants = receipiants
}

func (m *Mail) SetBccReceipiants(receipiants []string) {
	m.BccReceipiants = receipiants
}

func (m *Mail) SetSubject(subject string) {
	m.Subject = subject
}

func (m *Mail) SetHTMLMessage(message string) {
	m.Body = message
}

func (m *Mail) SendMail() {
	mail := gomail.NewMessage()

	if m.Sender != "" {
		mail.SetHeader("From", m.Sender)
	}
	if len(m.Receipiants) > 0 {
		mail.SetHeader("To", m.Receipiants...)
	}
	if len(m.CcReceipiants) > 0 {
		mail.SetHeader("Cc", m.CcReceipiants...)
	}
	if len(m.BccReceipiants) > 0 {
		mail.SetHeader("Bcc", m.BccReceipiants...)
	}

	mail.SetHeader("Subject", m.Subject)
	mail.SetBody("text/html", m.Body)
	// m.Attach("/home/Alex/lolcat.jpg")

	if m.Mailer != nil && m.Mailer.Dialer != nil {
		if err := m.Mailer.Dialer.DialAndSend(mail); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("undefined dailer, could not send email")
	}
}
