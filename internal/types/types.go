package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email" validate:"required,email"`
	Token     string             `bson:"token"`
	Confirmed bool               `bson:"confirmed"`
	Active    bool               `bson:"active"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type SendConfirmationEmailRequest struct {
	Email string `json:"email"`
	Token string `bson:"token"`
}

func ConvertSendConfirmationEmailRequestToUser(sendConfirmationEmailRequest SendConfirmationEmailRequest) User {
	return User{
		Email: sendConfirmationEmailRequest.Email,
		Token: sendConfirmationEmailRequest.Token,
	}
}

type SendWelcomeEmailRequest struct {
	Email string `json:"email"`
}

func ConvertSendWelcomeEmailRequestToUser(sendWelcomeEmailRequest SendWelcomeEmailRequest) User {
	return User{
		Email: sendWelcomeEmailRequest.Email,
	}
}
