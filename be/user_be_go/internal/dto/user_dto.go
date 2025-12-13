package dto

import "time"

// UserCreateDTO 创建用户请求
type UserCreateDTO struct {
	Name      string `json:"name" binding:"required,max=100" example:"张三"`
	Age       *int   `json:"age" binding:"omitempty,min=0,max=150" example:"20"`
	Email     string `json:"email" binding:"required,email" example:"zhangsan@example.com"`
	BirthDate *Date  `json:"birthDate" binding:"omitempty" example:"2003-01-01"`
}

// UserUpdateDTO 更新用户请求
// @Description 更新用户请求参数
type UserUpdateDTO struct {
	Name      *string `json:"name" binding:"omitempty,max=100" example:"张三"`                    // 姓名
	Age       *int    `json:"age" binding:"omitempty,min=0,max=150" example:"20"`                // 年龄
	Email     *string `json:"email" binding:"omitempty,email" example:"zhangsan@example.com"`   // 邮箱
	BirthDate *Date   `json:"birthDate" binding:"omitempty" example:"2003-01-01"`                 // 出生日期
}

// UserQueryDTO 查询用户请求
type UserQueryDTO struct {
	Name      string     `form:"name"`
	Email     string     `form:"email"`
	Age       *int       `form:"age"`
	BeType    string     `form:"beType"`
	StartDate *time.Time `form:"startDate"`
	EndDate   *time.Time `form:"endDate"`
	Page      int        `form:"page,default=1"`
	Limit     int        `form:"limit,default=10"`
}

// PageResponse 分页响应
type PageResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"totalPages"`
}

