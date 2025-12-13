package repository

import (
	"user_be_go/internal/domain"
	"user_be_go/internal/dto"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindAll 查询所有用户
func (r *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

// FindByID 根据ID查询用户
func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

// FindByEmail 根据邮箱查询用户
func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

// FindWithPageAndCount 分页查询用户
func (r *UserRepository) FindWithPageAndCount(query dto.UserQueryDTO) ([]domain.User, int64, error) {
	var users []domain.User
	var total int64

	queryBuilder := r.db.Model(&domain.User{})

	// 构建查询条件
	if query.Name != "" {
		queryBuilder = queryBuilder.Where("name ILIKE ?", "%"+query.Name+"%")
	}
	if query.Email != "" {
		queryBuilder = queryBuilder.Where("email ILIKE ?", "%"+query.Email+"%")
	}
	if query.Age != nil {
		queryBuilder = queryBuilder.Where("age = ?", *query.Age)
	}
	if query.BeType != "" {
		queryBuilder = queryBuilder.Where("be_type = ?", query.BeType)
	}
	if query.StartDate != nil {
		queryBuilder = queryBuilder.Where("birth_date >= ?", *query.StartDate)
	}
	if query.EndDate != nil {
		queryBuilder = queryBuilder.Where("birth_date <= ?", *query.EndDate)
	}

	// 统计总数
	if err := queryBuilder.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.Limit
	err := queryBuilder.
		Order("updated_at DESC").
		Offset(offset).
		Limit(query.Limit).
		Find(&users).Error

	return users, total, err
}

// Create 创建用户
func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// Update 更新用户
func (r *UserRepository) Update(id int, user *domain.User) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(id int) error {
	return r.db.Delete(&domain.User{}, id).Error
}

