package usecase_test

import (
	"github.com/danisusanto/go-dating-service/mocks"
	"google.golang.org/grpc/grpclog"
)

type UsecaseMock struct {
	userRepository mocks.UserRepository
	preferenceRepository mocks.PreferenceRepository
	packageRepository mocks.PackageRepository
	userPackageRepository mocks.UserPackageRepository

	passwordManager mocks.PasswordManager
	tokenManager mocks.TokenManager

	logger grpclog.LoggerV2
}
