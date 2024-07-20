package service

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/redis_model"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/database/mapper"
    "blog-backend/app/model"
    "blog-backend/global"
    "blog-backend/pkg/email"
    "blog-backend/pkg/helper"
    "blog-backend/pkg/jwt"
    "errors"
    "fmt"
    jwtlib "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "time"
)

func GetUserInfoByUUID(requestData *request.GetByUUIDRequest) (
        responseData *response.UserResponse, code error_code.ErrorCode) {
    user, err := mapper.GetUserByUUID(requestData.UUID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, error_code.UsernameIsNotExist
        }
        return nil, error_code.DatabaseError
    }

    responseData = &response.UserResponse{
        UUID:           user.UUID,
        Username:       user.Username,
        Nickname:       user.Nickname,
        Email:          user.Email,
        AvatarImageURL: user.AvatarImageURL,
    }

    return responseData, error_code.SUCCESS
}

func SendVerifyCodeWithEmail(requestData *request.GetRegisterVerifyCodeWithEmailRequest, requestId string) (responseData *response.SendVerifyCodeWithEmailResponse, code error_code.ErrorCode) {
    if _, ok := IsUsernameExist(requestData.Username); ok {
        return nil, error_code.UsernameAlreadyExists
    }

    if _, ok := IsEmailExist(requestData.Email); ok {
        return nil, error_code.EmailAlreadyInUse
    }

    // 生成验证码
    verifyCode := helper.MakeStr(global.CONFIG.VerifyCodeConfig.Length, helper.DigitAlpha)

    // 将验证码存入 Redis
    key := global.VerifyCodeKeyPrefix + requestId
    value := &redis_model.RegisterUserInfo{
        Username:   requestData.Username,
        Password:   requestData.Password,
        Email:      requestData.Email,
        VerifyCode: verifyCode,
    }
    expire := time.Second * time.Duration(global.CONFIG.VerifyCodeConfig.Expire)
    global.RDB.Set(key, value, expire)

    // 发送邮件
    email.SendEmail(requestData.Email, struct {
        Username   string
        Email      string
        VerifyCode string
    }{
        Username:   requestData.Username,
        Email:      requestData.Email,
        VerifyCode: verifyCode,
    })

    responseData = &response.SendVerifyCodeWithEmailResponse{
        RequestID: requestId,
    }
    return responseData, error_code.SUCCESS
}

func CreateUserWithEmailVerifyCode(requestData *request.CreateUserByEmailVerifyCodeRequest) error_code.ErrorCode {
    key := global.VerifyCodeKeyPrefix + requestData.RequestID
    value := &redis_model.RegisterUserInfo{}
    if !global.RDB.GetWithScan(key, value) {
        return error_code.RedisError
    }

    fmt.Printf("%#v\n", value)

    if value.VerifyCode != requestData.VerifyCode {
        return error_code.VerifyCodeExpired
    }

    if value.Username != requestData.Username ||
            value.Password != requestData.Password ||
            value.Email != requestData.Email {
        return error_code.RegisterInfoMismatch
    }

    salt := helper.MakeStr(16, helper.DigitAlphaPunct)
    password := helper.MD5(requestData.Password + salt)
    user := &model.User{
        Username: requestData.Username,
        Password: password,
        Email:    requestData.Email,
    }
    if err := mapper.CreateUser(user); err != nil {
        return error_code.DatabaseError
    }

    return error_code.SUCCESS
}

func LoginByUsernameAndPassword(requestData *request.LoginByUsernameAndPasswordRequest) (responseData *response.LoginResponse, code error_code.ErrorCode) {
    user, isExist := IsUsernameExist(requestData.Username)
    if !isExist {
        return nil, error_code.UsernameIsNotExist
    }

    // 校验密码
    passwordMd5 := helper.MD5(requestData.Password + user.Salt)
    if passwordMd5 != user.Password {
        return nil, error_code.PasswordVerifyFailed
    }

    tokenStr, code := CreateToken(user)
    if !error_code.IsSuccess(code) {
        return nil, code
    }

    responseData = &response.LoginResponse{
        User: response.UserResponse{
            UUID:           user.UUID,
            Username:       user.Username,
            Nickname:       user.Nickname,
            Email:          user.Email,
            AvatarImageURL: user.AvatarImageURL,
        },
        Token: tokenStr,
    }

    return responseData, error_code.SUCCESS
}

func IsUsernameExist(username string) (*model.User, bool) {
    user, err := mapper.GetUserByUsername(username)
    return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}

func IsEmailExist(email string) (*model.User, bool) {
    user, err := mapper.GetUserByEmail(email)
    return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}

func CreateToken(user *model.User) (tokenStr string, code error_code.ErrorCode) {
    j := &jwt.JWT{
        SigningKey: []byte(global.CONFIG.JWTConfig.SigningKey),
    }

    claims := jwt.CustomClaims{
        UUID:     user.UUID,
        Username: user.Username,
        RegisteredClaims: jwtlib.RegisteredClaims{
            ID:        uuid.New().String(),
            ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(time.Second * time.Duration(global.CONFIG.JWTConfig.ExpiresTime))),
            NotBefore: jwtlib.NewNumericDate(time.Now()),
            Issuer:    "Luminous",
            IssuedAt:  jwtlib.NewNumericDate(time.Now()),
        },
    }

    tokenStr, err := j.GenToken(claims)
    if err != nil {
        return "", error_code.AuthTokenCreateFailed
    }

    return tokenStr, error_code.SUCCESS
}
