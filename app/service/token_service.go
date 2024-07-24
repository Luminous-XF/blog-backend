package service

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/model"
    "blog-backend/config"
    "blog-backend/pkg/jwt"
    "github.com/google/uuid"
)

func CreateTokenByRefreshToken(
        req *request.CreateTokenByRefreshTokenRequest,
) (rsp *response.CreateTokenByRefreshTokenResponse, code error_code.ErrorCode) {
    refreshToken := req.RefreshToken

    j := &jwt.JWT{
        SigningKey: []byte(config.CONFIG.JWTConfig.SigningKey),
    }
    var claims *jwt.CustomClaims
    if claims, code = j.ParseToken(refreshToken); !code.IsSuccess() {
        return nil, code
    }

    accessTokenStr, code := CreateAccessToken(&model.User{
        UUID:     claims.UUID,
        Username: claims.Username,
    })
    if !code.IsSuccess() {
        return nil, code
    }

    rsp = &response.CreateTokenByRefreshTokenResponse{
        AccessToken: accessTokenStr,
    }

    return rsp, error_code.CreateSuccess
}

// CreateAccessToken 创建 Token
func CreateAccessToken(user *model.User) (AccessTokenStr string, code error_code.ErrorCode) {
    j := &jwt.JWT{
        SigningKey: []byte(config.CONFIG.JWTConfig.SigningKey),
    }

    AccessTokenStr, err := j.GenAccessToken(user, uuid.New().String())
    if err != nil {
        return "", error_code.AuthTokenCreateFailed
    }

    return AccessTokenStr, error_code.CreateSuccess
}

// RefreshToken 创建 Token
func RefreshToken(user *model.User) (RefreshTokenStr string, code error_code.ErrorCode) {
    j := &jwt.JWT{
        SigningKey: []byte(config.CONFIG.JWTConfig.SigningKey),
    }

    RefreshTokenStr, err := j.GenAccessToken(user, uuid.New().String())
    if err != nil {
        return "", error_code.AuthTokenCreateFailed
    }

    return RefreshTokenStr, error_code.CreateSuccess
}
