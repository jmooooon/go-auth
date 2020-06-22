package auth

import (
	"context"

	"github.com/jmooooon/go-auth/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{
}

// SignIn ...
func (s *GRPCServer) SignIn (ctx context.Context, req *api.SignInRequest) (*api.SignInResponce, error) {
	return nil, nil
}

// SignUp ...
func (s *GRPCServer) SignUp (ctx context.Context, req *api.SignUpRequest) (*api.SignUpResponce, error) {
	return &api.SignUpResponce{ Status: true, Message:""}, nil
}

// TokenDecrypt ...
func (s *GRPCServer) TokenDecrypt (ctx context.Context, req *api.TokenRequest) (*api.TokenResponce, error) {
	return nil, nil
}