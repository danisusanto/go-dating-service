package controller

import (
	"context"
	"github.com/danisusanto/go-dating-service/app/domain/exception"
	"github.com/danisusanto/go-dating-service/app/usecase"
	schemav1 "github.com/danisusanto/go-dating-service/proto/schema/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SchemaV1API) OrderPackage(ctx context.Context, req *schemav1.OrderPackageRequest) (*emptypb.Empty, error)  {
	if err := req.Validate(); err != nil {
		return nil, exception.Invalid(err.Error())
	}

	userClaim, err := s.TokenManager.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	err = s.OrderPackageUsecase.Call(ctx, &usecase.OrderPackageUsecaseParams{
		UserID:       userClaim.ID,
		PackageID:    uint(req.GetId()),
		TotalPayment: uint(req.GetTotalPayment()),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}