package services


import (
	"errors"
	"time"

	"github.com/wahyujatirestu/simple-procurement-system/config"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	modelutil "github.com/wahyujatirestu/simple-procurement-system/utils/models"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	CreateToken(user models.User) (string, error)
	VerifyToken(tokenString string) (modelutil.JwtPayloadClaim, error)
}

type jwtService struct {
	cfg config.JWTConfig
}

func NewJwtService(cfg config.JWTConfig) JwtService {
	return &jwtService{cfg: cfg}
}

func (j *jwtService) CreateToken(user models.User) (string, error) {
	claims := modelutil.JwtPayloadClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.AccessTokenLifetime)),
		},
		UserId: user.ID,
		Role:   user.Role,
	}

	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	return token.SignedString(j.cfg.JwtSignatureKey)
}

func (j *jwtService) VerifyToken(tokenString string) (modelutil.JwtPayloadClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&modelutil.JwtPayloadClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return j.cfg.JwtSignatureKey, nil
		},
	)

	if err != nil {
		return modelutil.JwtPayloadClaim{}, err
	}

	claims, ok := token.Claims.(*modelutil.JwtPayloadClaim)
	if !ok || !token.Valid {
		return modelutil.JwtPayloadClaim{}, errors.New("invalid token")
	}

	return *claims, nil
}
