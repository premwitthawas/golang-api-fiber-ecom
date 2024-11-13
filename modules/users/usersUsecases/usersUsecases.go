package usersUsecases

import (
	"github.com/premwitthawas/basic-api/modules/users"
	"github.com/premwitthawas/basic-api/modules/users/usersRepositories"
)

type IUserUsecase interface {
	InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
}

type UserUsecase struct {
	userRepository usersRepositories.IUserRepository
}

func UserUsecaseInit(usersRepository usersRepositories.IUserRepository) IUserUsecase {
	return &UserUsecase{
		userRepository: usersRepository,
	}
}

func (u *UserUsecase) InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
	// fmt.Print(req)
	if err := req.BcryptHashing(); err != nil {
		return nil, err
	}
	result, err := u.userRepository.InsertUser(req, false)
	if err != nil {
		return nil, err
	}
	// return nil, nil
	return result, nil
}
