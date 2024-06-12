package controller_test

import (
	"github.com/danisusanto/go-dating-service/mocks"
)

type SchemaV1APIMock struct {
	loginUsecase mocks.LoginUsecase
	registerUsecase mocks.RegisterUsecase
	listUserUsecase mocks.ListUserUsecase
	createUserPreferenceUsecase mocks.CreateUserPreferenceUsecase
	listPackageUsecase mocks.ListPackageUsecase
	orderPackageUsecase mocks.OrderPackageUsecase
	meUsecase mocks.MeUsecase

	tokenManager mocks.TokenManager
}