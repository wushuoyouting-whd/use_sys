package constants

// API 路径常量
const (
	APIPrefix = "/api"
	
	// 用户模块
	UserPrefix = "/users"
	
	// 系统模块
	SysPrefix = "/sys"
)

// HTTP 状态码
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusConflict            = 409
	StatusUnprocessableEntity = 422
	StatusInternalServerError = 500
)

// 业务错误码
const (
	CodeSuccess = 0
	
	// 通用错误码 1xxxx
	CodeInvalidParams = 10001
	CodeInternalError = 10002
	
	// 用户相关错误码 2xxxx
	CodeUserNotFound    = 20001
	CodeUserEmailExists = 20002
)

// 错误消息
const (
	MsgSuccess           = "success"
	MsgInvalidParams     = "请求参数错误"
	MsgInternalError     = "服务器内部错误"
	MsgUserNotFound      = "User not found"
	MsgUserEmailExists   = "Email already exists"
	MsgUserCreated       = "user created success"
	MsgUserUpdated       = "user updated success"
	MsgUserDeleted       = "user deleted success"
)

