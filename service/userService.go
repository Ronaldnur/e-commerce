package service

import (
	"fmt"
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/user_repository"
	"net/http"
)

type userService struct {
	userRepo user_repository.Repository
}

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(userLogin dto.NewLoginRequest) (*dto.NewLoginResponse, errs.MessageErr)
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	if len(payload.Password) < 6 {
		return nil, errs.NewBadRequest("Password harus 6 Characters")
	}

	existingEmail, err := us.userRepo.GetUserByEmail(payload.Email)

	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingEmail != nil {
		return nil, errs.NewDuplicateDataError("Email telah digunakan, silakan coba email lain")
	}

	existingUsername, err := us.userRepo.GetUserByUsername(payload.Username)

	if err != nil && err.Status() == http.StatusInternalServerError {
		return nil, err
	}

	if existingUsername != nil {
		return nil, errs.NewDuplicateDataError("Username telah digunakan, silakan coba email lain")
	}

	user := entity.User{
		Email:    payload.Email,
		Username: payload.Username,
		Password: payload.Password,
		Verified: false,
		Address:  entity.Address(payload.Address),
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	userId, err := us.userRepo.CreateUser(user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
		Data: dto.UserDataResponse{
			Id:       userId,
			Email:    payload.Email,
			Username: payload.Username,
			Verified: user.Verified,
			Address:  dto.GetAddressResponse(payload.Address),
		},
	}
	return &response, nil
}

func (u *userService) Login(userLogin dto.NewLoginRequest) (*dto.NewLoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userLogin)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepo.GetUserByEmail(userLogin.EmailorUsername)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			user, err = u.userRepo.GetUserByUsername(userLogin.EmailorUsername)
			if err != nil {
				if err.Status() == http.StatusNotFound {
					return nil, errs.NewBadRequest("invalid email/username")
				}
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	isValidPassword := user.ComparePassword(userLogin.Password)
	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid password")
	}

	token := user.GenerateToken()

	response := dto.NewLoginResponse{
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
		Data: dto.TokenResponse{
			Token: token,
		},
	}
	return &response, nil
}
