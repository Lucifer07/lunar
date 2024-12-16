package user

import (
	"context"

	"github.com/Lucifer07/lunar/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context, filters map[string]interface{}, offset int, limit int, orderby string) ([]users, error)
	FindAllUnscoped(ctx context.Context, filters map[string]interface{}, offset int, limit int, orderby string) ([]users, error)
	FindAllDeleted(ctx context.Context, filters map[string]interface{}, offset int, limit int, orderby string) ([]users, error)
	FindByID(ctx context.Context, id int) (users, error)
	Create(ctx context.Context, user users) (users, error)
	Update(ctx context.Context, id int, user users) (users, error)
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
}
type userRepo struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepo{conn: conn}
}
func (u *userRepo) FindAll(ctx context.Context, filters map[string]interface{}, offset int, limit int, orderby string) ([]users, error) {
	var users []users
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Where(filters).Offset(offset).Limit(limit).Order(orderby).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepo) FindAllUnscoped(ctx context.Context, filters map[string]interface{}, offset int, limit int, orderby string) ([]users, error) {
	var users []users
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Unscoped().Where(filters).Offset(offset).Limit(limit).Order(orderby).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepo) FindAllDeleted(ctx context.Context, filters map[string]interface{}, offset int, limit int, orderby string) ([]users, error) {
	var users []users
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Unscoped().Where(filters).Offset(offset).Limit(limit).Order(orderby).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepo) FindByID(ctx context.Context, id int) (users, error) {
	var user users
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.First(&user, id).Error; err != nil {
		return users{}, err
	}
	return user, nil
}
func (u *userRepo) Create(ctx context.Context, user users) (users, error) {
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Create(&user).Error; err != nil {
		return users{}, err
	}
	return user, nil
}
func (u *userRepo) Update(ctx context.Context, id int, user users) (users, error) {
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Model(&users{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return users{}, err
	}
	return user, nil
}
func (u *userRepo) Delete(ctx context.Context, id int) error {
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Delete(&users{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepo) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var count int64
	db := utils.GetQueryRunner(ctx, u.conn)
	if err := db.Model(&users{}).Where(filters).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
