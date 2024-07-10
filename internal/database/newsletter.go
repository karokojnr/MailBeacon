package database

import (
	"MailBeacon/internal/types"
	"context"
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

	// result := col.FindOne(ctx, bson.M{"email": user.Email})
	// if result.Err() == nil {
	// 	return errors.New("user already exists")
	// }

	_, err := col.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
