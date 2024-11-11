package middlewaresHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/premwitthawas/basic-api/config"
	"github.com/premwitthawas/basic-api/modules/entities"
	"github.com/premwitthawas/basic-api/modules/middlewares/middlewaresUsecases"
)

type middlewareHandlersErrCode string

const (
	routerCheckErrorCode middlewareHandlersErrCode = "middleware-router-check-error"
)

type IMiddlewareHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
}

type MiddlewareHandler struct {
	cfg                 config.IConfig
	middlewaresUsecases middlewaresUsecases.IMiddlewareUsecase
}

func MiddlewareHandlerInit(cfg config.IConfig, mu middlewaresUsecases.IMiddlewareUsecase) IMiddlewareHandler {
	return &MiddlewareHandler{middlewaresUsecases: mu, cfg: cfg}
}

func (mh *MiddlewareHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,HEAD,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (mh *MiddlewareHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewReponse(c).Error(fiber.StatusNotFound, string(routerCheckErrorCode), "router not found").Res()
	}
}

func (mh *MiddlewareHandler) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:        "${pid} ${time} [${ip}] ${status} - ${method} ${path}\n",
		TimeFormat:    "01/02/2006",
		TimeZone:      "Bangkok/Asia",
		DisableColors: false,
	})
}
