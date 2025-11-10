//go:generate mockgen -source=$GOFILE -destination=mock/mail_svc_mock.go
package service

import (
	"context"
)

// MailService is interface for mail service
type MailService interface {
	Send(ctx context.Context, subject, body string, to []string) error
}
