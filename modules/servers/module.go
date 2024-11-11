package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/premwitthawas/basic-api/modules/middlewares/middlewaresHandlers"
	"github.com/premwitthawas/basic-api/modules/middlewares/middlewaresRepositories"
	"github.com/premwitthawas/basic-api/modules/middlewares/middlewaresUsecases"
	monitorHandlers "github.com/premwitthawas/basic-api/modules/monitor/handlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type ModuleFactory struct {
	router     fiber.Router
	server     *server
	middleware middlewaresHandlers.IMiddlewareHandler
}

func ModuleFactoryInit(r fiber.Router, s *server, middleware middlewaresHandlers.IMiddlewareHandler) IModuleFactory {
	return &ModuleFactory{
		router:     r,
		server:     s,
		middleware: middleware,
	}
}

func ModuleMiddlewareInit(s *server) middlewaresHandlers.IMiddlewareHandler {
	repository := middlewaresRepositories.MiddlewareRepositoryInit(s.db)
	usecase := middlewaresUsecases.MiddlewareUsecaseInit(repository)
	handler := middlewaresHandlers.MiddlewareHandlerInit(s.cfg, usecase)
	return handler
}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandlerInit(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}
