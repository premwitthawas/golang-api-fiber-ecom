package entities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/premwitthawas/basic-api/pkg/logger"
)

type IResponse interface {
	Success(code int, data any) IResponse
	Error(code int, TraceId string, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"message"`
}

func NewReponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}
}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	logger.LoggerProjectInit(r.Context, &r.Data).Print().Save()
	return r
}
func (r *Response) Error(code int, traceId string, msg string) IResponse {
	r.ErrorRes = &ErrorResponse{
		TraceId: traceId,
		Msg:     msg,
	}
	r.StatusCode = code
	r.IsError = true
	return r
}
func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if r.IsError {
			return &r.ErrorRes
		}
		return &r.Data
	}())
}
