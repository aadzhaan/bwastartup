package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userId int) (string, error) //string adalah token yang dihasilkan
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedtoken string) (*jwt.Token, error) {
	//param pertama ada string token
	//param kedua adalah function
	token, err := jwt.Parse(encodedtoken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // HMAC sama dengan H256
		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
