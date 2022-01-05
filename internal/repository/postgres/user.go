package postgres

import (
	"context"

	"github.com/persiangophers/userland/internal/entiry/model"
)

func (p *postgresql) CreateUser(ctx context.Context, user *model.User) error {
	return p.db.Create(user).Error
}

func (p *postgresql) UpdateUser(ctx context.Context, user *model.User) error {
	return p.db.Save(user).Error
}

func (p *postgresql) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := p.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (p *postgresql) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := p.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (p *postgresql) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := p.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (p *postgresql) DeleteUser(ctx context.Context, id int) error {
	return p.db.Where("id = ?", id).Unscoped().Delete(&model.User{}).Error
}
