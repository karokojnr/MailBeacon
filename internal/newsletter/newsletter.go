package newsletter

import (
	"MailBeacon/internal/database"
	"MailBeacon/internal/mailer"
	"MailBeacon/internal/pubsub"
	"MailBeacon/internal/types"
	"context"
)

type NewsletterSevice interface {
	SignUp(context.Context, types.User) error
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
		return err
	}

	err = n.pubSub.Publish("newsletter-signup", user)
	if err != nil {
		return err
	}

	err = n.mailer.SendConfirmationEmail(user.Email, user.Token)
	if err != nil {
		return err
	}
	return nil
}
