package controller

import (
	"context"
	"github.com/danisusanto/go-dating-service/app/domain/exception"
	"github.com/danisusanto/go-dating-service/app/usecase"
	schemav1 "github.com/danisusanto/go-dating-service/proto/schema/v1"
)

func (s *SchemaV1API) Login(ctx context.Context, req *schemav1.LoginRequest) (*schemav1.AuthResponse, error)  {


	if err := req.Validate(); err != nil {
		return nil, exception.Invalid(err.Error())
	}
	res, err := s.LoginUsecase.Call(ctx, &usecase.LoginUsecaseParams{
		Email: req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &schemav1.AuthResponse{
		Data: &schemav1.AuthResponseData{
			AccessToken: res.Token,
		},
	}, nil
}
