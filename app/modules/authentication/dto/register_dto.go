package authentication_dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type SignUpPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
	Phone           string `json:"phone,omitempty"`
}

func NewSignUpPayload() *SignUpPayload {
	return new(SignUpPayload)
}

func (payload *SignUpPayload) Validate() error {
	return validation.ValidateStruct(payload,
		validation.Field(&payload.Email, validation.Required.Error("email is required"), is.Email.Error("email is invalid"), validation.Length(1, 255).Error("email length must be between 1 and 255")),
		validation.Field(&payload.Password, validation.Required.Error("password is required"), validation.Length(1, 255).Error("password length must be between 1 and 255")),
		validation.Field(&payload.ConfirmPassword, validation.Required.Error("confirm password is required"), validation.Length(1, 255).Error("confirm password length must be between 1 and 255"), validation.By(func(value any) error {
			if value != payload.Password {
				return validation.NewError(
					"confirm_password_mismatch",
					"password and confirm password do not match",
				)
			}
			return nil
		})),
		validation.Field(&payload.Name, validation.Required.Error("name is required"), validation.Length(1, 255).Error("name length must be between 1 and 255")),
		validation.Field(&payload.Phone, validation.Length(10, 10).Error("phone length must be 10")),
	)
}

// make cost optional
func (payload *SignUpPayload) HashPassword(cost ...int) error {
	if len(cost) == 0 {
		cost = []int{bcrypt.DefaultCost}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), cost[0])
	if err != nil {
		return err
	}
	payload.Password = string(hashedPassword)
	return nil
}
