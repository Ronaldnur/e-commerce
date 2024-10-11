package handler

import (
	"encoding/json"
	"mongo-api/dto"
	"mongo-api/pkg/errs"
	"mongo-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UserService service.UserService
}

func NewUserHandler(UserService service.UserService) userHandler {
	return userHandler{
		UserService: UserService,
	}
}

func (uh *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUserRegister dto.NewUserRequest

	if err := json.NewDecoder(r.Body).Decode(&newUserRegister); err != nil {

		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		w.WriteHeader(errBindJson.Status())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": errBindJson.Status(),
			"error":      errBindJson.Message(),
		})

		return
	}

	result, err := uh.UserService.CreateNewUser(newUserRegister)

	if err != nil {
		// Handle any errors from the service
		w.WriteHeader(err.Status()) // Status code dari error
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     false,
			"statusCode": err.Status(),
			"message":    err.Message(),
			"error":      err.Error(),
		})
		return
	}
	w.WriteHeader(result.StatusCode)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": http.StatusInternalServerError,
			"error":      "Failed to encode response",
		})
		return
	}

}
func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUserLogin dto.NewLoginRequest

	// Cek apakah request body valid
	if err := json.NewDecoder(r.Body).Decode(&newUserLogin); err != nil {
		// Handle error untuk request body yang tidak valid
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		w.WriteHeader(errBindJson.Status())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": errBindJson.Status(),
			"error":      errBindJson.Message(),
		})
		return
	}

	// Memproses login
	result, err := uh.UserService.Login(newUserLogin)
	if err != nil {
		// Handle error dari service
		w.WriteHeader(err.Status()) // Status code dari error
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": err.Status(),
			"message":    err.Message(),
			"error":      err.Error(),
		})
		return
	}

	// Berhasil login, kirim response dengan status code yang benar
	w.WriteHeader(result.StatusCode)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		// Handle error jika gagal mengencode response
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": http.StatusInternalServerError,
			"error":      "Failed to encode response",
		})
		return
	}
}

func (uh *userHandler) RegisterGin(ctx *gin.Context) {
	var newUserRegister dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRegister); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.UserService.CreateNewUser(newUserRegister)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (uh *userHandler) LoginGin(ctx *gin.Context) {
	var newUserRequest dto.NewLoginRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	result, err := uh.UserService.Login(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
