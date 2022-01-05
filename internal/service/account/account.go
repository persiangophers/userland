package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/entiry/model"
	"github.com/persiangophers/userland/internal/pkg/logger"
	"github.com/persiangophers/userland/internal/repository"
	"github.com/persiangophers/userland/internal/service"
	"github.com/persiangophers/userland/pkg/myjwt"
	"github.com/persiangophers/userland/pkg/password_util"
)

type acc struct {
	cfg    config.Account
	db     repository.Postgres
	logger logger.Logger
	jwt    myjwt.Myjwt
}

func New(cfg config.Account, postgres repository.Postgres, logger logger.Logger) service.Account {

	// instantiate jwt
	myjwtIns := myjwt.GetInstance()
	// setup jwt token generator
	myjwtIns.SetSecret([]byte(cfg.TokenSecret))
	myjwtIns.SetClaims("userId", "exp")

	return &acc{
		cfg:    cfg,
		db:     postgres,
		logger: logger,
		jwt:    myjwtIns,
	}
}

func (a *acc) Login(ctx context.Context, username, password string) (string, error) {

	user, err := a.db.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	_ = user
	// check password
	if ok := password_util.ComparePassword(password, user.Password); !ok {
		return "", errors.New("password is wrong")
	}
	// generating token with calims
	return a.jwt.NewToken(user.ID, time.Now().Add(time.Hour*72).Unix())
}

func (a *acc) Register(ctx context.Context, user *model.User) error {

	if len(user.Username) < a.cfg.MinUsernameLength {
		return fmt.Errorf("minimum username length: %d", a.cfg.MinUsernameLength)
	}

	if _, err := a.db.GetUserByUsername(ctx, user.Username); err == nil {
		return errors.New("username exists")
	}

	if _, err := a.db.GetUserByEmail(ctx, user.Email); err == nil {
		return errors.New("email exists")
	}

	if len(user.Password) < 8 {
		return errors.New("password minimum length: 8")
	}

	hashedPass, err := password_util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPass

	return a.db.CreateUser(ctx, user)
}
