package authentication_dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignInPayload() *SignInPayload {
	return new(SignInPayload)
}

func (payload *SignInPayload) Validate() error {
	return validation.ValidateStruct(payload,
		validation.Field(&payload.Email, validation.Required.Error("email is required"), is.Email.Error("email is invalid"), validation.Length(1, 255).Error("email length must be between 1 and 255")),
		validation.Field(&payload.Password, validation.Required.Error("password is required"), validation.Length(1, 255).Error("password length must be between 1 and 255")),
	)
}
