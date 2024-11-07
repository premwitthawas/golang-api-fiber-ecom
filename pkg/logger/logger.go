package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/premwitthawas/basic-api/pkg/utils"
)

type ILoggerProject interface {
	Print() ILoggerProject
	Save()
	SetQueryString(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(data any)
}

type LoggerProject struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Metode     string `json:"metode"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func LoggerProjectInit(c *fiber.Ctx, res any) ILoggerProject {
	log := &LoggerProject{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Metode:     c.Method(),
		Path:       c.Path(),
		StatusCode: c.Response().StatusCode(),
	}
	log.SetQueryString(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

func (l *LoggerProject) Print() ILoggerProject {
	utils.Debug(l)
	return l
}
func (l *LoggerProject) Save() {
	data := utils.Output(l)
	filename := fmt.Sprintf("./assets/logs/loger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02 15:04:05"), ":", "-"))
	if err := os.MkdirAll("./assets/logs", os.ModePerm); err != nil {
		log.Fatalf("error creating directory: %v", err)
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}
func (l *LoggerProject) SetQueryString(c *fiber.Ctx) {
	var querystring any
	if err := c.QueryParser(&querystring); err != nil {
		log.Printf("error parsing querystring: %v", err)
	}
	l.Query = querystring
}
func (l *LoggerProject) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("error parsing body: %v", err)
		l.Body = "Unprocessable Entity"
		return
	}
	switch l.Path {
	case "v1/users/signup":
		l.Body = "never gonna give you up"
	default:
		l.Body = body
	}
}
func (l *LoggerProject) SetResponse(res any) {
	l.Response = res
}
