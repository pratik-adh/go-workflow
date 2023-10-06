package service

import (
	"CRUD/config"
	"CRUD/constants"
	"CRUD/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// //jwt service
type JWTService interface {
	GenerateToken( email string, password string) string
	ValidateToken(token string) (*jwt.Token, error)
}



type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Pratik",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = constants.SecretKey
	}
	return secret
}

type User struct {
    ID uuid.UUID
}

func (jwtService *jwtServices) GenerateToken( email string, password string) string {

	var user User

	if err := config.DB.Model(&models.Registration{}).Select("id").Where("email = ?", email).Scan(&user).Error; err != nil {
		panic(err)
	}

	claims := &models.AuthCustomClaims{
		Email: email,
		Password: password,
		ID: user.ID, 
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    jwtService.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtService.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtService *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); 
		!isvalid {
			return nil, fmt.Errorf("Invalid token: %v", token.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})
}
