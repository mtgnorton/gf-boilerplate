package consts

const (
	// CacheTokenKeyPrefix 缓存token的key前缀,存储形式为:
	// 1. token:statistics
	// 2. userID:token
	CacheTokenKeyPrefix = "token"
	// CacheTokenUserPrefix 缓存用户绑定的token的key前缀,存储形式为:
	// 1. userID:[]token
	CacheTokenUserPrefix = "token_user"
)
