package entity

import (
	"mongo-api/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         string    `bson:"_id"`
	Email      string    `bson:"email"`
	Username   string    `bson:"username"`
	Password   string    `bson:"password"`
	Verified   bool      `bson:"veridied"`
	Address    Address   `bson:"address"`
	Created_at time.Time `bson:"created_at"`
	Updated_at time.Time `bson:"updated_at"`
}

type Address struct {
	Street     string `bson:"street"`
	City       string `bson:"city"`
	Province   string `bson:"province"`
	Country    string `bson:"country"`
	PostalCode string `bson:"postal_code"`
}

var secret_key = "RAHASIA"

var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Id = string(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = string(email)
	}
	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err

}

func (u User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret_key))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
