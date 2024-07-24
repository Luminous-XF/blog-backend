package jwt

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/model"
    "blog-backend/config"
    "errors"
    "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"
    "time"
)

type CustomClaims struct {
    UUID       uuid.UUID `json:"uuid"`
    Username   string    `json:"username"`
    BufferTime int64     `json:"buffer_time"`
    jwt.RegisteredClaims
}

type JWT struct {
    SigningKey []byte
}

func NewJWT() *JWT {
    return &JWT{
        SigningKey: []byte(config.CONFIG.JWTConfig.SigningKey),
    }
}

// GenAccessToken 生成 Access Token
func (j *JWT) GenAccessToken(user *model.User, tokenID string) (string, error) {
    claims := CustomClaims{
        UUID:     user.UUID,
        Username: user.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ID: tokenID,
            ExpiresAt: jwt.NewNumericDate(
                time.Now().Add(time.Second * time.Duration(config.CONFIG.JWTConfig.ExpiresTime)),
            ),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    config.CONFIG.JWTConfig.Issuer,
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return tokenStr.SignedString(j.SigningKey)
}

// GenRefreshToken 生成 Refresh Token
func (j *JWT) GenRefreshToken(user *model.User, tokenID string) (string, error) {
    claims := CustomClaims{
        UUID:     user.UUID,
        Username: user.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ID: tokenID,
            ExpiresAt: jwt.NewNumericDate(
                time.Now().Add(time.Second * time.Duration(config.CONFIG.JWTConfig.StorageTime)),
            ),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    config.CONFIG.JWTConfig.Issuer,
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return tokenStr.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error_code.ErrorCode) {
    t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        var ve *jwt.ValidationError
        if errors.As(err, &ve) {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return nil, error_code.AuthTokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                return nil, error_code.AuthTokenExpired
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return nil, error_code.AuthTokenNotValidYet
            } else {
                return nil, error_code.AuthTokenInvalid
            }
        }
    }

    if t != nil {
        if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
            return claims, error_code.SUCCESS
        }
        return nil, error_code.AuthTokenInvalid
    }
    return nil, error_code.AuthTokenInvalid
}
