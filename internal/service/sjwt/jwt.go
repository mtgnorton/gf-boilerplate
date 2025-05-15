package sjwt

import (
	"context"
	"gf-boilerplate/internal/consts"
	"gf-boilerplate/internal/model"
	"gf-boilerplate/internal/service/errctx"
	"gf-boilerplate/internal/service/st"
	"gf-boilerplate/internal/service/stime"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	*model.Identity
	jwt.RegisteredClaims
}

type Statistics struct {
	ExpireAt     int64 `json:"exp"` // token过期时间
	RefreshAt    int64 `json:"ra"`  // 刷新时间
	RefreshCount int64 `json:"rc"`  // 刷新次数
}

// 启用后，同一账号在设备A登录后再于设备B登录，将立即踢出设备A；
func GenerateToken(ctx context.Context, identity *model.Identity) (string, error) {
	config, err := st.GetConfig().JwtConfig(ctx)
	if err != nil {
		return "", err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Identity: identity,
	}).SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", errctx.Wrap(ctx, err)
	}
	duration := time.Duration(config.Expires) * time.Second
	statistics := &Statistics{
		ExpireAt:     st.NowTime().Add(duration).Unix(),
		RefreshAt:    st.NowTime().Unix(),
		RefreshCount: 0,
	}

	tokenUserKey := GetTokenUserKey(identity.Id)
	tokenKey := GetTokenKey(token)

	st.GetJWTCache().Set(ctx, tokenKey, statistics, duration)
	tokensVar, err := st.GetJWTCache().Get(ctx, tokenUserKey)
	if err != nil {
		return "", err
	}
	if tokensVar.IsEmpty() {
		tokens := map[string]bool{token: true}
		st.GetJWTCache().Set(ctx, tokenUserKey, tokens, duration)
		return token, nil
	}

	var tokens map[string]bool
	if subErr := tokensVar.Scan(&tokens); subErr != nil {
		return "", subErr
	}
	// 如果开启了禁止多端登录,删除掉其他token
	if config.MultiLogin {
		for token := range tokens {
			st.GetJWTCache().Remove(ctx, GetTokenKey(token))
		}
		tokens = map[string]bool{}
	}
	tokens[token] = true
	st.GetJWTCache().Set(ctx, tokenUserKey, tokens, duration)
	return token, nil
}

func ParseToken(ctx context.Context, token string) (*model.Identity, error) {
	if token == "" {
		return nil, errctx.New(ctx, "cg.auth.token.invalid")
	}
	config, err := st.GetConfig().JwtConfig(ctx)
	if err != nil {
		return nil, err
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errctx.New(ctx, "cg.auth.token.invalid")
	}
	claims, ok := jwtToken.Claims.(*Claims)
	if !ok || claims.Identity == nil {
		return nil, errctx.New(ctx, "cg.auth.token.invalid")
	}
	statisticsVar, err := st.GetJWTCache().Get(ctx, GetTokenKey(token))
	if err != nil {
		return nil, err
	}
	if statisticsVar.IsEmpty() {
		return nil, errctx.New(ctx, "cg.auth.token.invalid")
	}
	var statistics *Statistics
	if subErr := statisticsVar.Scan(&statistics); subErr != nil {
		return nil, subErr
	}
	if statistics == nil || statistics.ExpireAt < st.NowTime().Unix() {
		return nil, errctx.New(ctx, "cg.auth.token.invalid")
	}
	// 如果开启了自动刷新,则刷新token
	if !config.AutoRefresh {
		return claims.Identity, nil
	}
	if config.MaxRefreshTimes > 0 && statistics.RefreshCount >= int64(config.MaxRefreshTimes) {
		return claims.Identity, nil
	}
	if statistics.RefreshAt+int64(config.RefreshInterval) < stime.Now().Unix() {
		return claims.Identity, nil
	}
	return nil, errctx.New(ctx, "cg.auth.token.invalid")
}

func GetTokenKey(tokenKey string) string {
	return consts.CacheTokenKeyPrefix + tokenKey
}

func GetTokenUserKey(userID uint64) string {
	return consts.CacheTokenUserPrefix + strconv.FormatUint(userID, 10)
}
