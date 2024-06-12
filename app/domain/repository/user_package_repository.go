package repository

import (
	"context"
	"github.com/danisusanto/go-dating-service/app/domain/entity"
)

type UserPackageRepository interface {
	Create(ctx context.Context, userPackage *entity.UserPackage) (*entity.UserPackage, error)
}