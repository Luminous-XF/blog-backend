package jwt

import (
    "blog-backend/app/common/error_code"
    "blog-backend/config"
    "errors"
    "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"
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

// GenToken 生成 Token
func (j *JWT) GenToken(claims CustomClaims) (string, error) {
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
