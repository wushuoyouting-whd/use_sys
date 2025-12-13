package controller

import (
	"strconv"
	"user_be_go/internal/constants"
	"user_be_go/internal/dto"
	"user_be_go/internal/service"
	"user_be_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

// List 获取用户列表
// @Summary 获取用户列表
// @Description 支持分页和条件查询的用户列表接口
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param limit query int false "每页数量" default(10)
// @Param name query string false "姓名（模糊查询）"
// @Param email query string false "邮箱（模糊查询）"
// @Param age query int false "年龄"
// @Param beType query string false "后端类型"
// @Param startDate query string false "开始日期"
// @Param endDate query string false "结束日期"
// @Success 200 {object} response.ApiResponse
// @Router /api/users [get]
func (c *UserController) List(ctx *gin.Context) {
	var query dto.UserQueryDTO
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Error(ctx, constants.StatusBadRequest, constants.MsgInvalidParams+": "+err.Error())
		return
	}

	// 设置默认值
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	data, err := c.service.GetUsersWithPageAndCount(query)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response.Success(ctx, data)
}

// GetByID 获取用户详情
// @Summary 获取用户详情
// @Description 根据ID获取用户详细信息
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.ApiResponse
// @Router /api/users/{id} [get]
func (c *UserController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, constants.StatusBadRequest, "无效的用户ID")
		return
	}

	user, err := c.service.GetUserByID(id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response.Success(ctx, user)
}

// Create 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserCreateDTO true "用户信息"
// @Success 200 {object} response.ApiResponse
// @Router /api/users [post]
func (c *UserController) Create(ctx *gin.Context) {
	var dto dto.UserCreateDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.Error(ctx, constants.StatusBadRequest, constants.MsgInvalidParams+": "+err.Error())
		return
	}

	user, err := c.service.CreateUser(dto)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response.Success(ctx, user, constants.MsgUserCreated)
}

// Update 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body dto.UserUpdateDTO true "用户信息"
// @Success 200 {object} response.ApiResponse
// @Router /api/users/{id} [put]
func (c *UserController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, constants.StatusBadRequest, "无效的用户ID")
		return
	}

	var dto dto.UserUpdateDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.Error(ctx, constants.StatusBadRequest, constants.MsgInvalidParams+": "+err.Error())
		return
	}

	user, err := c.service.UpdateUser(id, dto)
	if err != nil {
		handleError(ctx, err)
		return
	}

	response.Success(ctx, user, constants.MsgUserUpdated)
}

// Delete 删除用户
// @Summary 删除用户
// @Description 根据ID删除用户
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.ApiResponse
// @Router /api/users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, constants.StatusBadRequest, "无效的用户ID")
		return
	}

	if err := c.service.DeleteUser(id); err != nil {
		handleError(ctx, err)
		return
	}

	response.Success(ctx, nil, constants.MsgUserDeleted)
}

