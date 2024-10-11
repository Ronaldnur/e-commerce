package dto

type NewUserRequest struct {
	Email    string            `json:"email" valid:"email,required"`
	Username string            `json:"username" valid:"required"`
	Password string            `json:"password" valid:"required"`
	Address  NewAddressRequest `json:"address"`
}

type UserDataResponse struct {
	Id       string             `json:"id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Verified bool               `json:"verified"`
	Address  GetAddressResponse `json:"address"`
}

type GetAddressResponse struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
type NewUserResponse struct {
	Result     string           `json:"result"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       UserDataResponse `json:"data"`
}

type NewLoginRequest struct {
	EmailorUsername string `json:"email/username" valid:"required"`
	Password        string `json:"password" valid:"required"`
}

type NewLoginResponse struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"token"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
