package usersHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/premwitthawas/basic-api/config"
	"github.com/premwitthawas/basic-api/modules/entities"
	"github.com/premwitthawas/basic-api/modules/users"
	"github.com/premwitthawas/basic-api/modules/users/usersUsecases"
)

type UserHandleStatusCode string

const (
	BodyParserError           UserHandleStatusCode = "bad-request-body-parser"
	IsEmailValidError         UserHandleStatusCode = "bad-request-email-invalid"
	SignUpCustomerError       UserHandleStatusCode = "bad-request-sign-up-customer"
	InternalServerErrorSignup UserHandleStatusCode = "internal-server-error-sign-up-customer"
)

type IUsersHandler interface {
	SignUpCustomer(ctx *fiber.Ctx) error
}

type UsersHandler struct {
	cfg          config.IConfig
	usersUsecase usersUsecases.IUserUsecase
}

func UsersHandlerInit(cfg config.IConfig, usersUsecase usersUsecases.IUserUsecase) IUsersHandler {
	return &UsersHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}

func (u *UsersHandler) SignUpCustomer(ctx *fiber.Ctx) error {
	req := new(users.UserRegisterReq)
	if err := ctx.BodyParser(req); err != nil {
		return entities.NewReponse(ctx).Error(fiber.StatusBadRequest, string(BodyParserError), err.Error()).Res()
	}
	if !req.IsEmailValid() {
		return entities.NewReponse(ctx).Error(fiber.StatusBadRequest, string(IsEmailValidError), "Email is invalid").Res()
	}

	result, err := u.usersUsecase.InsertCustomer(req)
	if err != nil {
		switch err.Error() {
		case "email is already exist":
			return entities.NewReponse(ctx).Error(fiber.StatusBadRequest, string(SignUpCustomerError), err.Error()).Res()
		case "username is already exist":
			return entities.NewReponse(ctx).Error(fiber.StatusBadRequest, string(SignUpCustomerError), err.Error()).Res()
		default:
			return entities.NewReponse(ctx).Error(fiber.StatusInternalServerError, string(InternalServerErrorSignup), err.Error()).Res()
		}
	}
	return entities.NewReponse(ctx).Success(fiber.StatusCreated, result).Res()
}
