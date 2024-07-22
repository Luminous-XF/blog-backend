package service

import (
    "blog-backend/app/common/constant"
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/redis_model"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/database"
    "blog-backend/app/model"
    "blog-backend/config"
    "blog-backend/global"
    "blog-backend/pkg/email"
    "blog-backend/pkg/helper"
    "blog-backend/pkg/jwt"
    "errors"
    jwtlib "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "time"
)

// GetUserInfoByUUID 通过 UUID 获取用户信息
func GetUserInfoByUUID(
        req *request.GetByUUIDRequest,
) (rsp *response.UserResponse, code error_code.ErrorCode) {
    user, err := database.GetUserByUUID(req.UUID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, error_code.UsernameIsNotExist
        }
        return nil, error_code.DatabaseError
    }

    rsp = &response.UserResponse{
        UUID:           user.UUID,
        Username:       user.Username,
        Nickname:       user.Nickname,
        Email:          user.Email,
        AvatarImageURL: user.AvatarImageURL,
    }

    return rsp, error_code.SUCCESS
}

// GetRegisterVerifyCodeWithEmail 通过邮箱获取注册验证码
func GetRegisterVerifyCodeWithEmail(
        req *request.GetRegisterVerifyCodeWithEmailRequest,
        requestId string,
) (rsp *response.SendVerifyCodeWithEmailResponse, code error_code.ErrorCode) {
    if _, ok := IsUsernameExist(req.Username); ok {
        return nil, error_code.UsernameAlreadyExists
    }

    if _, ok := IsEmailExist(req.Email); ok {
        return nil, error_code.EmailAlreadyInUse
    }

    // 生成验证码
    verifyCode := helper.MakeStr(
        config.CONFIG.VerifyCodeConfig.Length,
        helper.DigitAlpha,
    )

    // 将验证码存入 Redis
    key := constant.VerifyCodeKeyPrefix + requestId
    value := &redis_model.RegisterUserInfo{
        Username:   req.Username,
        Password:   req.Password,
        Email:      req.Email,
        VerifyCode: verifyCode,
    }
    expire := time.Second * time.Duration(config.CONFIG.VerifyCodeConfig.Expire)
    global.RDB.Set(key, value, expire)

    // 发送邮件
    email.SendEmail(req.Email, struct {
        Username   string
        Email      string
        VerifyCode string
    }{
        Username:   req.Username,
        Email:      req.Email,
        VerifyCode: verifyCode,
    })

    rsp = &response.SendVerifyCodeWithEmailResponse{
        RequestID: requestId,
    }
    return rsp, error_code.SUCCESS
}

// CreateUserWithEmailVerifyCode 通过邮箱验证码创建账号
func CreateUserWithEmailVerifyCode(
        req *request.CreateUserByEmailVerifyCodeRequest,
) (*response.UserResponse, error_code.ErrorCode) {
    // 校验用户名是否已被使用
    if _, ok := IsUsernameExist(req.Username); ok {
        return nil, error_code.UsernameAlreadyExists
    }

    // 校验邮箱是否已被使用
    if _, ok := IsEmailExist(req.Email); ok {
        return nil, error_code.EmailAlreadyInUse
    }

    // 校验验证码
    key := constant.VerifyCodeKeyPrefix + req.RequestID
    value := &redis_model.RegisterUserInfo{}
    if !global.RDB.GetWithScan(key, value) {
        return nil, error_code.RedisError
    }

    if value.VerifyCode != req.VerifyCode {
        return nil, error_code.VerifyCodeExpired
    }

    // 校验注册信息, 需要保证获取验证码时的信息与提交注册时的信息一致
    if value.Username != req.Username ||
            value.Password != req.Password ||
            value.Email != req.Email {
        return nil, error_code.RegisterInfoMismatch
    }

    // 创建 User 实例
    user := &model.User{
        UUID:     uuid.New(),
        Username: req.Username,
        Password: req.Password,
        Email:    req.Email,
    }

    // 创建 User 记录, 入库
    if err := database.CreateUser(user); err != nil {
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            return nil, error_code.UsernameAlreadyExists
        }
        return nil, error_code.DatabaseError
    }

    // 返回新创建的 User 信息
    rsp := &response.UserResponse{
        UUID:           user.UUID,
        Username:       user.Username,
        Nickname:       user.Nickname,
        Email:          user.Email,
        AvatarImageURL: user.AvatarImageURL,
    }

    return rsp, error_code.CreateSuccess
}

// LoginByUsernameAndPassword 使用账号密码登录
func LoginByUsernameAndPassword(
        req *request.LoginByUsernameAndPasswordRequest,
) (rsp *response.LoginResponse, code error_code.ErrorCode) {
    user, isExist := IsUsernameExist(req.Username)
    if !isExist {
        return nil, error_code.UsernameIsNotExist
    }

    // 校验密码
    passwordMd5 := helper.MD5(req.Password + user.Salt)
    if passwordMd5 != user.Password {
        return nil, error_code.PasswordVerifyFailed
    }

    tokenStr, code := CreateToken(user)
    if !code.IsSuccess() {
        return nil, code
    }

    rsp = &response.LoginResponse{
        User: response.UserResponse{
            UUID:           user.UUID,
            Username:       user.Username,
            Nickname:       user.Nickname,
            Email:          user.Email,
            AvatarImageURL: user.AvatarImageURL,
        },
        Token: tokenStr,
    }

    return rsp, error_code.CreateSuccess
}

// IsUsernameExist 判断用户名是否存在
func IsUsernameExist(username string) (*model.User, bool) {
    user, err := database.GetUserByUsername(username)
    return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}

// IsEmailExist 判断邮箱是否存在
func IsEmailExist(email string) (*model.User, bool) {
    user, err := database.GetUserByEmail(email)
    return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}

// CreateToken 创建 Token
func CreateToken(user *model.User) (tokenStr string, code error_code.ErrorCode) {
    j := &jwt.JWT{
        SigningKey: []byte(config.CONFIG.JWTConfig.SigningKey),
    }

    claims := jwt.CustomClaims{
        UUID:     user.UUID,
        Username: user.Username,
        RegisteredClaims: jwtlib.RegisteredClaims{
            ID: uuid.New().String(),
            ExpiresAt: jwtlib.NewNumericDate(
                time.Now().Add(time.Second * time.Duration(config.CONFIG.JWTConfig.ExpiresTime)),
            ),
            NotBefore: jwtlib.NewNumericDate(time.Now()),
            Issuer:    "Luminous",
            IssuedAt:  jwtlib.NewNumericDate(time.Now()),
        },
    }

    tokenStr, err := j.GenToken(claims)
    if err != nil {
        return "", error_code.AuthTokenCreateFailed
    }

    return tokenStr, error_code.CreateSuccess
}
