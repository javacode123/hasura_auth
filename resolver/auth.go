package resolver

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/hasura_auth/model"
	"github.com/hasura_auth/service"
	"github.com/spf13/cast"
)

func (r *Resolver) Login(ctx context.Context, loginParam model.LoginParam) (*LoginResolver, error) {
	userInfo, err := service.UserInfo(loginParam.Name, loginParam.Password)
	if err != nil {
		return nil, err
	}
	if userInfo.UserId == 0 {
		return nil, errors.Newf("name: %s, password: %s no exist", loginParam.Name, loginParam.Password)
	}
	defaultRole := ""
	if len(userInfo.Roles) > 0 {
		defaultRole = userInfo.Roles[0]
	}

	token, err := service.CreateToken(
		ctx,
		model.CustomClaims{
			AllowedRoles: userInfo.Roles,
			UserId:       cast.ToString(userInfo.UserId),
			DefaultRole:  defaultRole},
	)
	if err != nil {
		return nil, err
	}
	return &LoginResolver{&model.JsonWebToken{
		Token: token,
	}}, nil
}

type LoginResolver struct {
	token *model.JsonWebToken
}

func (l *LoginResolver) Token() string {
	return l.token.Token
}
