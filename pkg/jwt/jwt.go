package jwt

import (
	"errors"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CompanyID uint   `json:"company_id"`
	gojwt.RegisteredClaims
}

type JWTManager struct {
	secret     []byte
	expiryHour int
}

func NewManager(secret string, expiryHour int) *JWTManager {
	return &JWTManager{
		secret:     []byte(secret),
		expiryHour: expiryHour,
	}
}

func (m *JWTManager) Generate(userID uint, email, role string, companyID uint) (string, error) {
	claims := Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		CompanyID: companyID,
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(time.Duration(m.expiryHour) * time.Hour)),
			IssuedAt:  gojwt.NewNumericDate(time.Now()),
		},
	}

	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *JWTManager) Verify(tokenStr string) (*Claims, error) {
	token, err := gojwt.ParseWithClaims(tokenStr, &Claims{}, func(t *gojwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*gojwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
