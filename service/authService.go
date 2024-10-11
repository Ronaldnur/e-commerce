package service

import (
	"fmt"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/product_repository"
	"mongo-api/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentitaction() gin.HandlerFunc
}

type authService struct {
	userRepo    user_repository.Repository
	productRepo product_repository.Repository
}

func NewAuthService(userRepo user_repository.Repository, productRepo product_repository.Repository) AuthService {
	return &authService{
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}
func (a *authService) Authentitaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			fmt.Println("User tidak ditemukan:", err)
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result
		ctx.Set("userData", user)

		ctx.Next()
	}
}
