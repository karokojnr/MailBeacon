package mailer

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sendgrid/sendgrid-go"
)

var (
	sendgridApiKey = os.Getenv("SENDGRID_API_KEY")
)

type sendgridMailer struct {
	sendgridClient *sendgrid.Client
}

func NewSendGrid() *sendgridMailer {
	sendgridClient := sendgrid.NewSendClient(sendgridApiKey)
	return &sendgridMailer{
		sendgridClient: sendgridClient,
	}
}

func (s *sendgridMailer) SendConfirmationEmail(email string, token string) error {
	return nil
}

func (s *sendgridMailer) SendWelcomeEmail(email string) error {
	return nil
}
