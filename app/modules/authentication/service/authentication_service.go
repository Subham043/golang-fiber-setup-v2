package authentication_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/subham043/golang-fiber-setup/app/middlewares/jwt"
	authentication_dto "github.com/subham043/golang-fiber-setup/app/modules/authentication/dto"
	"github.com/subham043/golang-fiber-setup/bootstrap/database/ent"
	"github.com/subham043/golang-fiber-setup/bootstrap/database/ent/user"
)

type AuthenticationService struct {
	db  *ent.Client
	jwt *jwt.JWTMiddleware
}

type IAuthenticationService interface {
	Index() error
}

func NewAuthenticationService(db *ent.Client, jwt *jwt.JWTMiddleware) *AuthenticationService {
	return &AuthenticationService{db: db, jwt: jwt}
}

func (h *AuthenticationService) userExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	return h.db.User.
		Query().
		Where(user.EmailEQ(email)).
		Exist(ctx)
}

func (h *AuthenticationService) userExistsByPhone(
	ctx context.Context,
	phone string,
) (bool, error) {
	return h.db.User.
		Query().
		Where(user.PhoneEQ(phone)).
		Exist(ctx)
}

type RegisterResult struct {
	User         *ent.User
	AccessToken  *string
	RefreshToken *string
}

func (h *AuthenticationService) Register(ctx context.Context, payload *authentication_dto.SignUpPayload) (*RegisterResult, error) {
	var phone *string
	if payload.Phone != "" {
		phone = &payload.Phone
	}
	userExistsByEmail, err := h.userExistsByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if userExistsByEmail {
		return nil, fiber.NewError(fiber.StatusBadRequest, "user with this email already exists")
	}
	if phone != nil {
		userExistsByPhone, err := h.userExistsByPhone(ctx, *phone)
		if err != nil {
			return nil, err
		}
		if userExistsByPhone {
			return nil, fiber.NewError(fiber.StatusBadRequest, "user with this phone already exists")
		}
	}
	user, err := h.db.User.
		Create().
		SetName(payload.Name).
		SetEmail(payload.Email).
		SetPassword(payload.Password).
		SetNillablePhone(phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	token_payload := jwt.JWTUserDTO{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
	access_token, err := h.jwt.GenerateAccessToken(token_payload)
	if err != nil {
		return nil, err
	}
	refresh_token, err := h.jwt.GenerateRefreshToken(token_payload)
	if err != nil {
		return nil, err
	}
	return &RegisterResult{User: user, AccessToken: &access_token, RefreshToken: &refresh_token}, nil
}
