package service

import (
	"math"
	"time"
	"user_be_go/internal/config"
	"user_be_go/internal/constants"
	"user_be_go/internal/domain"
	"user_be_go/internal/dto"
	"user_be_go/internal/repository"
	"user_be_go/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUsers 获取所有用户
func (s *UserService) GetUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.NewNotFoundError(constants.MsgUserNotFound)
	}
	return user, nil
}

// GetUsersWithPageAndCount 分页查询用户
func (s *UserService) GetUsersWithPageAndCount(query dto.UserQueryDTO) (*dto.PageResponse, error) {
	users, total, err := s.repo.FindWithPageAndCount(query)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))

	return &dto.PageResponse{
		Data:       users,
		Total:      total,
		Page:       query.Page,
		Limit:      query.Limit,
		TotalPages: totalPages,
	}, nil
}

// CreateUser 创建用户
func (s *UserService) CreateUser(dto dto.UserCreateDTO) (*domain.User, error) {
	// 检查邮箱是否已存在
	existing, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, utils.NewConflictError(constants.MsgUserEmailExists)
	}

	// 创建用户
	user := &domain.User{
		Name:      dto.Name,
		Email:     dto.Email,
		BirthDate: dto.BirthDate.ToTimePtr(),
		BeType:    config.AppConfig.BeType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if dto.Age != nil {
		user.Age = *dto.Age
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(id int, dto dto.UserUpdateDTO) (*domain.User, error) {
	// 检查用户是否存在
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.NewNotFoundError(constants.MsgUserNotFound)
	}

	// 如果更新邮箱，检查新邮箱是否已存在
	if dto.Email != nil && *dto.Email != user.Email {
		existing, err := s.repo.FindByEmail(*dto.Email)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, utils.NewConflictError(constants.MsgUserEmailExists)
		}
		user.Email = *dto.Email
	}

	// 更新字段
	if dto.Name != nil {
		user.Name = *dto.Name
	}
	if dto.Age != nil {
		user.Age = *dto.Age
	}
	if dto.BirthDate != nil {
		user.BirthDate = dto.BirthDate.ToTimePtr()
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(id, user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id int) error {
	// 检查用户是否存在
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return utils.NewNotFoundError(constants.MsgUserNotFound)
	}

	return s.repo.Delete(id)
}
