package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"datink/internal/models/auth"
	"datink/internal/models/user/dto"
	"datink/pkg/util"
	"datink/pkg/util/httputil"
)

func (s service) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repository.User.GetByEmail(ctx, s.resource.Database.GetMaster(), req.Email)
	if err != nil {
		return dto.LoginResponse{}, httputil.ErrorWrap(httputil.ErrBadRequest, errors.New("invalid email or password"))
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return dto.LoginResponse{}, httputil.ErrorWrap(httputil.ErrBadRequest, errors.New("invalid email or password"))
	}

	claims := auth.UserJwtAuth{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.resource.Config.JwtSigningKey))

	return dto.LoginResponse{
		AccessToken: ss,
	}, nil

}
