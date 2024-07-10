package types

import (
	"MailBeacon/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func ConvertSignUpRequestToUser(signUpRequest SignUpRequest) User {
	return User{
		Email: signUpRequest.Email,
		Token: utils.GenerateRandomToken(),
	}
}

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

// SendConfirmationEmailRequest to user
func ConvertSendConfirmationEmailRequestToUser(sendConfirmationEmailRequest SendConfirmationEmailRequest) User {
	return User{
		Email: sendConfirmationEmailRequest.Email,
		Token: sendConfirmationEmailRequest.Token,
	}
}
