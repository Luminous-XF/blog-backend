package error_code

type ErrorCode int // 错误码

//go:generate stringer -type ErrorCode -linecomment

const (
    SUCCESS ErrorCode = iota + 20000 // 成功
)

// 40xxx 客户端错误
const (
    ERROR          ErrorCode = iota + 40000 // 失败
    ParamBindError                          // 参数信息错误
)

// 41xxx User 模块错误
const (
    UsernameIsNotExist    ErrorCode = iota + 41000 // 用户名不存在
    PasswordVerifyFailed                           // 密码校验失败
    UsernameAlreadyExists                          // 用户名已存在
    EmailAlreadyInUse                              // 邮箱已被使用
    RegisterInfoMismatch                           // 注册信息不符
    VerifyCodeExpired                              // 验证码不存在或已失效
    AuthFailed                                     // 认证失败
    AuthTokenNULL                                  // 没有 Token
    AuthTokenExpired                               // Token 已过期
    AuthTokenNotValidYet
    AuthTokenMalformed
    AuthTokenInvalid
    AuthTokenCreateFailed // Token 创建失败
)

// 6xxxx 数据库相关错误
// 60xxx MySQL 数据库错误
const (
    DatabaseError       ErrorCode = iota + 60000 // MySQL 数据库错误
    QueryPostListFailed                          // 查询帖子列表错误
)

// 62xxx Redis 错误

const (
    RedisError ErrorCode = iota + 61000
)

var errorMsg = map[ErrorCode]string{
    SUCCESS: "Ok!",

    ERROR:          "Error!",
    ParamBindError: "There was an error with the parameters provided.",

    UsernameIsNotExist:    "The entered username does not exist.",
    PasswordVerifyFailed:  "The password you entered is incorrect. Please try again.",
    UsernameAlreadyExists: "The username already exists.",
    EmailAlreadyInUse:     "The email address is already in use.",
    RegisterInfoMismatch:  "Information mismatch.",
    VerifyCodeExpired:     "The Verification code does not exist or has expired.",

    AuthFailed:            "Auth failed.",
    AuthTokenNULL:         "No authorization token found.",
    AuthTokenExpired:      "Auth token is expired.",
    AuthTokenNotValidYet:  "Auth token is not valid.",
    AuthTokenMalformed:    "Auth token malformed.",
    AuthTokenInvalid:      "Auth token is invalid.",
    AuthTokenCreateFailed: "Token create failed.",

    DatabaseError:       "MySQL Database Error.",
    QueryPostListFailed: "Unable to Fetch Post List.",

    RedisError: "Redis Error.",
}

func ErrMsg(code ErrorCode) string {
    return errorMsg[code]
}

func IsSuccess(code ErrorCode) bool {
    return code >= 20000 && code < 30000
}
