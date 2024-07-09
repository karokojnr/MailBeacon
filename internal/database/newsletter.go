package database

import (
	"MailBeacon/internal/types"
	"context"
	"fmt"
	"log"
)

type NewsletterStore interface {
	AddUser(context.Context, types.User) error
}

const (
	DbName   = "mailbeacon"
	CollName = "users"
)

func (d *Database) AddUser(ctx context.Context, user types.User) error {
	col := d.db.Database(DbName).Collection(CollName)

	_, err := col.InsertOne(ctx, user)
	if err != nil {
		fmt.Fprintf(log.Writer(), "could not insert user: %v", err)
		return err
	}
	return nil
}
