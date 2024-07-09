package newsletter

import (
	"MailBeacon/internal/database"
	"MailBeacon/internal/types"
	"context"
)

type NewsletterSevice interface {
	SignUp(context.Context, types.User) error
}

type newsletterService struct {
	store database.NewsletterStore
}

func NewNewsletterService(store database.NewsletterStore) *newsletterService {
	return &newsletterService{
		store: store,
	}
}

func (n *newsletterService) SignUp(ctx context.Context, user types.User) error {
	return n.store.AddUser(ctx, user)
}
