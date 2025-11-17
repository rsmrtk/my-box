package jwt

import (
	"fmt"
	"time"

	gojwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	gojwt.RegisteredClaims
	CustomerID string `json:"customer_id"`
	OS         string `json:"os ,omitempty"`
	AppVersion string `json:"app_version ,omitempty"`
}

type JWT interface {
	Generate(customerID string, os string, appVersion string) (string, error)
	Verify(token string) (*Claims, error)
}

type jwt struct {
	secret   string
	duration time.Duration
}

func New(secret string, duration time.Duration) JWT {
	return &jwt{secret: secret, duration: duration}
}

func (m *jwt) Generate(customerID string, os string, appVersion string) (string, error) {
	claims := Claims{
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(m.duration)),
			IssuedAt:  gojwt.NewNumericDate(time.Now()),
		},
		CustomerID: customerID,
		OS:         os,
		AppVersion: appVersion,
	}
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *jwt) Verify(token string) (*Claims, error) {
	tkn, err := gojwt.ParseWithClaims(token, &Claims{}, func(token *gojwt.Token) (any, error) {
		if _, ok := token.Method.(*gojwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid access token signing method: %v", token.Header["alg"])
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid access token claims type: %T", tkn.Claims)
	}

	if err := claims.Valid(); err != nil {
		return nil, fmt.Errorf("invalid access token claims: %w", err)
	}

	return claims, nil
}
