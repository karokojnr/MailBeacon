package newsletter

import (
	"MailBeacon/internal/database"
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
}

func NewNewsletterService(store database.NewsletterStore, pubSub pubsub.PubSub) *newsletterService {
	return &newsletterService{
		store:  store,
		pubSub: pubSub,
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
	return nil
}
