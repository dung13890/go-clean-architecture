package mail

import (
	"bytes"
	"context"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"

	"go-app/internal/domain/service"
	"go-app/internal/infrastructure/config"
	"go-app/pkg/errors"
)

// Email is the type used for email messages github.com/jordan-wright/email
type Email struct {
	From    string
	To      []string
	Subject string
	Body    string // Plaintext message or Html message (optional)
	auth    smtp.Auth
	addr    string
}

// NewEmail creates an Email, and returns the pointer to it.
func NewEmail() service.MailService {
	mailConf := config.GetEmailConfig()

	return &Email{
		From: mailConf.From,
		auth: smtp.PlainAuth("", mailConf.Username, mailConf.Password, mailConf.Host),
		addr: mailConf.Host + ":" + strconv.Itoa(mailConf.Port),
	}
}

// msg converts the Email object to a []byte representation, including all needed MIMEHeaders, boundaries, etc.
func (e *Email) msg() []byte {
	buff := &bytes.Buffer{}
	buff.WriteString("To: " + strings.Join(e.To, ", ") + "\r\n")
	buff.WriteString("From: " + e.From + "\r\n")
	buff.WriteString("Subject: " + e.Subject + "\r\n")
	buff.WriteString("MIME-Version: 1.0\r\n")
	buff.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buff.WriteString("Content-Transfer-Encoding: quoted-printable\r\n")
	buff.WriteString("\r\n" + e.Body)

	return buff.Bytes()
}

// Send an email using the given host and SMTP auth (optional), returns any error thrown by smtp.SendMail
func (e *Email) Send(_ context.Context, subj, body string, to []string) error {
	e.Subject = subj
	e.Body = body
	e.To = to

	for i := range e.To {
		addr, err := mail.ParseAddress(e.To[i])
		if err != nil {
			return errors.ErrBadRequest.Wrap(err)
		}
		to = append(to, addr.Address)
	}

	from, err := mail.ParseAddress(e.From)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}
	msg := e.msg()

	// Check to make sure there is at least one recipient and one "From" address
	if e.From == "" || len(to) == 0 {
		return errors.ErrSendEmailFromToInvalid.Trace()
	}

	return smtp.SendMail(e.addr, e.auth, from.Address, to, msg)
}
