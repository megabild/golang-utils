package mbmail

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
)

type mailer struct {
	Host     string
	Username string
	Password string
	From     string
	FromName string
	Port     int
	Charset  string
	IsHTML   bool
}

// Create
func Create(host string, username string, password string, fromName string) mailer {
	return mailer{
		Host:     host,
		Username: username,
		Password: password,
		From:     username,
		FromName: fromName,
		Port:     587,
		Charset:  "utf-8",
		IsHTML:   true,
	}
}

// Send to one recipient
func (m *mailer) Send(to string, subject string, msg string) error {
	return m.SendMultiple([]string{to}, subject, msg)
}

// Send to multiple recipients
func (m *mailer) SendMultiple(to []string, subject string, msg string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	contentType := "text/plain"
	if m.IsHTML {
		contentType = "text/html"
	}

	header := make(map[string]string)
	header["From"] = m.FromName + " <" + m.From + ">"
	header["To"] = strings.Join(to, ";")
	header["Subject"] = subject
	header["Content-Type"] = contentType + "; charset=\"" + m.Charset + "\""
	header["Content-Transfer-Encoding"] = "base64"

	body := ""
	for k, v := range header {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	body += "\r\n" + base64.StdEncoding.EncodeToString([]byte(msg))

	return smtp.SendMail(m.Host+":"+strconv.Itoa(m.Port), auth, m.Username, to, []byte(body))
}
