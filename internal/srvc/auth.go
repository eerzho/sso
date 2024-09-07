package srvc

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"sso/internal/def"
	"sso/internal/dto"
	"sso/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	rTokenExpiresHour int
	aTokenExpiresHour int
	rTokenLength      int
	jwtSecret         []byte
	userSrvc          UserSrvc
	refreshTokenSrvc  RefreshTokenSrvc
}

func NewAuth(jwtSecret string, userSrvc UserSrvc, refreshTokenSrvc RefreshTokenSrvc) *Auth {
	return &Auth{
		rTokenExpiresHour: 24,
		aTokenExpiresHour: 2,
		rTokenLength:      50,
		jwtSecret:         []byte(jwtSecret),
		userSrvc:          userSrvc,
		refreshTokenSrvc:  refreshTokenSrvc,
	}
}

func (a *Auth) Login(ctx context.Context, email, password, ip string) (*dto.Token, error) {
	const op = "srvc.Auth.Login"

	user, err := a.validateCredential(ctx, email, password)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, rToken, err := a.createRTokenByUser(ctx, user, ip)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	aToken, err := a.generateAToken(user, refreshToken, ip)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.Token{
		AToken: aToken,
		RToken: rToken,
	}, nil
}

func (a *Auth) DecodeAToken(ctx context.Context, aToken string) (*dto.Claims, error) {
	const op = "srvc.Auth.DecodeAToken"

	token, err := jwt.ParseWithClaims(aToken, &dto.Claims{}, a.getSigningKey)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	claims, ok := token.Claims.(*dto.Claims)
	if !ok {
		return nil, fmt.Errorf("%s: %w", op, def.ErrInvalidClaimsType)
	}

	return claims, nil
}

func (a *Auth) validateCredential(ctx context.Context, email, password string) (*model.User, error) {
	const op = "srvc.Auth.validateCredential"

	user, err := a.userSrvc.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, def.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, def.ErrInvalidCredentials)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, fmt.Errorf("%s: %w", op, def.ErrInvalidCredentials)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (a *Auth) createRTokenByUser(ctx context.Context, user *model.User, ip string) (*model.RefreshToken, string, error) {
	const op = "srvc.Auth.createRTokenByUser"

	err := a.refreshTokenSrvc.DeleteByUser(ctx, user)
	if err != nil && !errors.Is(err, def.ErrNotFound) {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	random := make([]byte, a.rTokenLength)
	_, err = rand.Read(random)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	hash, err := bcrypt.GenerateFromPassword(random, bcrypt.DefaultCost)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err := a.refreshTokenSrvc.CreateByUser(
		ctx,
		user,
		ip,
		string(hash),
		time.Now().Add(time.Duration(a.rTokenExpiresHour)*time.Hour),
	)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	return refreshToken, base64.StdEncoding.EncodeToString(random), nil
}

func (a *Auth) generateAToken(user *model.User, refreshToken *model.RefreshToken, ip string) (string, error) {
	const op = "srvc.Auth.generateAToken"

	claims := dto.Claims{
		IP:             ip,
		UserID:         user.ID,
		RefreshTokenID: refreshToken.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.aTokenExpiresHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(a.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) getSigningKey(token *jwt.Token) (interface{}, error) {
	const op = "srv.Auth.getSigningKey"

	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("%s: %w", op, def.ErrInvalidSigningMethod)
	}

	return a.jwtSecret, nil
}
