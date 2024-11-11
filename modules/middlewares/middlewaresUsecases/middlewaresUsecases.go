package middlewaresUsecases

import "github.com/premwitthawas/basic-api/modules/middlewares/middlewaresRepositories"

type IMiddlewareUsecase interface {
}

type MiddlewareUsecase struct {
	middlewareRepository middlewaresRepositories.IMiddlewareRepository
}

func MiddlewareUsecaseInit(mr middlewaresRepositories.IMiddlewareRepository) IMiddlewareUsecase {
	return &MiddlewareUsecase{middlewareRepository: mr}
}
