package provider

import (
	"github.com/danisusanto/go-dating-service/app/domain/manager"
	"github.com/danisusanto/go-dating-service/config"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenManager = manager.NewTokenManager([]byte(config.ENV("JWT_SECRET_KEY", "secret")))
	passwordManager = manager.NewPasswordManager(bcrypt.DefaultCost)
)
