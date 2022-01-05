package postgres

import (
	"fmt"

	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/entiry/model"
	"github.com/persiangophers/userland/internal/pkg/logger"
	"github.com/persiangophers/userland/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresql struct {
	db     *gorm.DB
	logger logger.Logger
}

func New(cfg config.Postgres, logger logger.Logger) (repository.Postgres, error) {
	db, err := gorm.Open(postgres.Open(dsn(cfg)), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return &postgresql{db: db, logger: logger}, nil
}

func dsn(cfg config.Postgres) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)
}
