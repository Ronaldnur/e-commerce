package user_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	GetUserByUsername(userUsername string) (*entity.User, errs.MessageErr)
	CreateUser(userPayload entity.User) (string, errs.MessageErr)
}
