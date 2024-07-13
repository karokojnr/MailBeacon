package newsletter

import (
	"MailBeacon/internal/database"
	"MailBeacon/internal/mailer"
	"MailBeacon/internal/pubsub"
	"MailBeacon/internal/types"
	"context"
	"errors"
	"log"
)

type NewsletterSevice interface {
	SignUp(context.Context, types.User) error
	SendConfirmationEmail(context.Context, types.User) error
	ConfirmSubscription(context.Context, types.User) error
	SendWelcomeEmail(context.Context, types.User) error
}

type newsletterService struct {
	store  database.NewsletterStore
	pubSub pubsub.PubSub
	mailer mailer.Mailer
}

func NewNewsletterService(store database.NewsletterStore, pubSub pubsub.PubSub, mailer mailer.Mailer) *newsletterService {
	return &newsletterService{
		store:  store,
		pubSub: pubSub,
		mailer: mailer,
	}
}

func (n *newsletterService) SignUp(ctx context.Context, user types.User) error {
	err := n.store.AddUser(ctx, user)
	if err != nil {
		log.Println("Error adding user to database: ", err)
		return err
	}

	err = n.pubSub.Publish("newsletter-signup", user)
	if err != nil {
		log.Println("Error publishing newsletter signup event: ", err)
		return err
	}
	return nil
}

func (n *newsletterService) SendConfirmationEmail(ctx context.Context, user types.User) error {
	err := n.mailer.SendConfirmationEmail(user.Email, user.Token)
	if err != nil {
		log.Println("Error sending confirmation email: ", err)
		return errors.New("error sending confirmation email. Please try again later")
	}
	return nil
}

func (n *newsletterService) ConfirmSubscription(ctx context.Context, user types.User) error {
	err := n.store.ConfirmUser(ctx, user)
	if err != nil {
		return err
	}

	err = n.pubSub.Publish("newsletter-email-confirmed", user)
	if err != nil {
		return err
	}
	return nil
}

func (n *newsletterService) SendWelcomeEmail(ctx context.Context, user types.User) error {
	err := n.mailer.SendWelcomeEmail(user.Email)
	if err != nil {
		return err
	}
	return nil
}
