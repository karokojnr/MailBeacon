package mailer

type Mailer interface {
	SendConfirmationEmail(email string, token string) error
	SendWelcomeEmail(email string) error
}
