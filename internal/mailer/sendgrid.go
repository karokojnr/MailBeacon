package mailer

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	sendgridApiKey      = os.Getenv("SENDGRID_API_KEY")
	sendgridSenderEmail = os.Getenv("SENDGRID_SENDER_EMAIL")
	appUrl              = os.Getenv("APP_URL")
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
	link := fmt.Sprintf("%s/confirm-email?token=%s&email=%s", appUrl, token, email)
	from := mail.NewEmail("MailBeacon", sendgridSenderEmail)
	subject := "Confirm your email address"
	to := mail.NewEmail("", email)
	plainTextContent := "Click the link below to confirm your email"
	htmlContent := fmt.Sprintf("Click <a href='%s'>here</a> to confirm your email", link)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	_, err := s.sendgridClient.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}
	return nil
}

func (s *sendgridMailer) SendWelcomeEmail(email string) error {
	from := mail.NewEmail("MailBeacon", sendgridSenderEmail)
	subject := "Welcome to MailBeacon!"
	to := mail.NewEmail("", email)
	plainTextContent := "Welcome to MailBeacon!"
	htmlContent := `Welcome to MailBeacon!🎉 We’re thrilled to have you on board.`
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := s.sendgridClient.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}
