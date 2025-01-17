package pocketworks

import (
	"net/mail"

	"github.com/pocketbase/pocketbase/tools/mailer"
)

func (p *PocketWorksApp) SendEmail(to, subject, htmlBody string) error {
	message := &mailer.Message{
		From: mail.Address{
			Address: p.pb.Settings().Meta.SenderAddress,
			Name:    p.pb.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: to}},
		Subject: subject,
		HTML:    htmlBody,
	}

	return p.pb.NewMailClient().Send(message)
}

// Usage example:
// app.SendEmail("user@example.com", "Welcome!", "<h1>Hello</h1>")
