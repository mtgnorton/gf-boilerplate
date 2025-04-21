// Package consts 定义系统常量
package consts

const (
	// ContextKey 上下文键名
	ContextKeyUser    = "UserInfo" // 用户信息
	ContextKeyRoleID  = "RoleID"   // 角色ID
	ContextKeyRoleKey = "RoleKey"  // 角色标识

	// CacheKey 缓存键名
	CacheKeyUser = "user:%d"      // 用户信息缓存
	CacheKeyMenu = "menu:role:%d" // 角色菜单缓存

	// TokenKey token相关
	TokenType    = "Bearer"
	TokenExpire  = 86400 // token过期时间(秒)
	TokenRefresh = 43200 // token刷新时间(秒)

	// Status 状态
	StatusEnabled  = "enabled"  // 启用
	StatusDisabled = "disabled" // 禁用
)
