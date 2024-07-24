package response

// LoginResponse 登录返回,user token 和过期时间
type LoginResponse struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

type CreateTokenByRefreshTokenResponse struct {
    AccessToken string `json:"accessToken"`
}
