package app

import (
	"context"
	"github.com/kloudlite/api/apps/auth/internal/entities"
	"github.com/kloudlite/api/pkg/repos"

	"github.com/kloudlite/api/apps/auth/internal/domain"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/auth"
	"github.com/kloudlite/api/pkg/errors"
)

type authGrpcServer struct {
	auth.UnimplementedAuthServer
	d domain.Domain
}

func (a *authGrpcServer) GetUser(ctx context.Context, in *auth.GetUserIn) (*auth.GetUserOut, error) {
	user, err := a.d.GetUserById(ctx, repos.ID(in.UserId))
	if err != nil {
		return nil, errors.NewE(err)
	}
	if user == nil {
		return nil, errors.Newf("could not find user with (id=%q)", in.UserId)
	}
	return &auth.GetUserOut{
		Id:    string(user.Id),
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (a *authGrpcServer) FromAccToken(token entities.AccessToken) *auth.AccessTokenOut {
	return &auth.AccessTokenOut{
		Id:       string(token.Id),
		UserId:   string(token.UserId),
		Email:    token.Email,
		Provider: token.Provider,
		OauthToken: &auth.OauthToken{
			AccessToken:  token.Token.AccessToken,
			TokenType:    token.Token.TokenType,
			RefreshToken: token.Token.RefreshToken,
			Expiry:       token.Token.Expiry.UnixMilli(),
		},
	}
}

func (a *authGrpcServer) EnsureUserByEmail(ctx context.Context, request *auth.GetUserByEmailRequest) (*auth.GetUserByEmailOut, error) {
	user, err := a.d.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if user == nil {
		user, err = a.d.EnsureUserByEmail(ctx, request.Email)
		if err != nil {
			return nil, errors.NewE(err)
		}
	}
	return &auth.GetUserByEmailOut{
		UserId: string(user.Id),
	}, nil
}

func (a *authGrpcServer) GetAccessToken(ctx context.Context, in *auth.GetAccessTokenRequest) (*auth.AccessTokenOut, error) {
	token, err := a.d.GetAccessToken(ctx, in.Provider, in.UserId, in.TokenId)
	if err != nil {
		return nil, errors.NewE(err)
	}
	if token == nil {
		return nil, errors.Newf("token is nil")
	}
	return a.FromAccToken(*token), nil
}

func fxRPCServer(d domain.Domain) auth.AuthServer {
	return &authGrpcServer{
		d: d,
	}
}
