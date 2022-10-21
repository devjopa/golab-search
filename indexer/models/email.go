package models

import (
	"fmt"
	"io"
	"time"

	"github.com/veqryn/go-email/email"
)

type EmailMessage struct {
	From    string
	To      string
	Subject string
	Date    string
	Body    string
}

func ParseMessage(file io.Reader) (*EmailMessage, error) {

	to := "N/A"
	emailDateValue := ""
	message, err := email.ParseMessage(file)

	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return nil, err
	}

	if len(message.Header.To()) > 0 {
		to = message.Header.To()[0]
	} else {
		to = message.Header.Get("X-To")
	}

	emailDate, err := message.Header.Date()
	if err != nil {
		emailDateValue = ""
	} else {
		emailDateValue = emailDate.Format(time.RFC822)
	}

	return &EmailMessage{
		From:    message.Header.From(),
		To:      to,
		Subject: message.Header.Subject(),
		Body:    fmt.Sprintf("%s", message.Body),
		Date:    emailDateValue,
	}, nil

}
