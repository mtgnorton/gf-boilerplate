package model

type JwtConfig struct {
	SecretKey       string `json:"secretKey"`
	Expires         int    `json:"expires"`
	AutoRefresh     bool   `json:"autoRefresh"`
	RefreshInterval int    `json:"refreshInterval"`
	MaxRefreshTimes int    `json:"maxRefreshTimes"`
	MultiLogin      bool   `json:"multiLogin"`
}

type Identity struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
