package database

import (
	"MailBeacon/internal/types"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type NewsletterStore interface {
	AddUser(context.Context, types.User) error
	ConfirmUser(context.Context, types.User) error
}

const (
	DbName   = "mailbeacon"
	CollName = "users"
)

func (d *Database) AddUser(ctx context.Context, user types.User) error {
	col := d.db.Database(DbName).Collection(CollName)

	result := col.FindOne(ctx, bson.M{"email": user.Email})
	if result.Err() == nil {
		return errors.New("user has already signed up")
	}

	_, err := col.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) ConfirmUser(ctx context.Context, user types.User) error {
	col := d.db.Database(DbName).Collection(CollName)

	result := col.FindOne(ctx, bson.M{"email": user.Email, "token": user.Token})
	if result.Err() != nil {
		return errors.New("user not found")
	}

	_, err := col.UpdateOne(ctx, bson.M{"email": user.Email, "token": user.Token}, bson.M{"$set": bson.M{"confirmed": true}})
	if err != nil {
		return err
	}

	return nil
}
