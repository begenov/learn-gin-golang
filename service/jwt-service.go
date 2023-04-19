package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "pragmaticreviews.com",
	}
}

func (jwtSrv *jwtService) GenerateToken(name string, admin bool) string {
	claims := &jwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t

}
func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	log.Println("tokenString:", tokenString)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
	if err != nil {
		log.Println("error:", err)
	}
	return token, err
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
