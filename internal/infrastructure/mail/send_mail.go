package mail

import (
	"emailn/internal/domain/campaign"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {
	d := gomail.NewDialer(os.Getenv("GOMAIL_SMTP"), 587, os.Getenv("GOMAIL_USER"), os.Getenv("GOMAIL_PASS"))

	var emails []string

	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GOMAIL_USER"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	err := d.DialAndSend(m)

	return err
}
