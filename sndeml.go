package sndeml

import (
	"context"
	"strings"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v3"
)

// SendEmail sends an html message using mailgun
func SendEmail(apiKey, recipient, sender, senderName, subject, body string) error {
	domain := strings.Split(sender, "@")[1]
	mg := mailgun.NewMailgun(domain, apiKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(senderName+"<"+sender+">", subject, "", recipient)
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
