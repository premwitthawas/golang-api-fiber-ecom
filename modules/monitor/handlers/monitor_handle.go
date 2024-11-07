package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/premwitthawas/basic-api/config"
	"github.com/premwitthawas/basic-api/modules/entities"
	"github.com/premwitthawas/basic-api/modules/monitor"
)

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type MonitorHandler struct {
	cfg config.IConfig
}

func MonitorHandlerInit(cfg config.IConfig) IMonitorHandler {
	return &MonitorHandler{cfg: cfg}
}

func (h *MonitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewReponse(c).Success(fiber.StatusOK, res).Res()
}
