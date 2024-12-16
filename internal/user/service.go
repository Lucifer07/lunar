package user

import (
	"context"

	"github.com/Lucifer07/lunar/internal/utils"
)
type UserService interface {
	GetAll(ctx context.Context) ([]users, error)
}
type userService struct {
	userRepo UserRepository
	transaction utils.Transactor
}
func NewUserService(userRepo UserRepository, transaction utils.Transactor) UserService {
	return &userService{
		userRepo: userRepo,
		transaction: transaction,
	}
}
func (u *userService) GetAll(ctx context.Context) ([]users, error) {
	return u.userRepo.FindAll(ctx, nil, 0, 0, "")
}