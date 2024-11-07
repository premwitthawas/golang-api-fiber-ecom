package servers

import (
	"github.com/gofiber/fiber/v2"
	monitorHandlers "github.com/premwitthawas/basic-api/modules/monitor/handlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type ModuleFactory struct {
	router fiber.Router
	server *server
}

func ModuleFactoryInit(r fiber.Router, s *server) IModuleFactory {
	return &ModuleFactory{
		router: r,
		server: s,
	}
}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandlerInit(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}
