package usersRepositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/premwitthawas/basic-api/modules/users"
	usersPatterns "github.com/premwitthawas/basic-api/modules/users/usersRepositories/patterns"
)

type IUserRepository interface {
	InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error)
}

type UserRepository struct {
	db *sqlx.DB
}

func UserRepositoryInit(db *sqlx.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := usersPatterns.InsertUser(ur.db, req, isAdmin)
	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}
	user, err := result.Result()
	if err != nil {
		return nil, err
	}
	return user, nil
}
