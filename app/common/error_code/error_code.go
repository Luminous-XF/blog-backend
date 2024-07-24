// Package error_code 错误码包
// 错误码规则:
// 错误码由 9 位数字构成, AAA-B-CC-DDD
// AAA: 对应 HTTP 状态码
// B: 系统级别模块编号
// CC: 资源级别模块编号
// DDD: 业务错误码
package error_code

// ErrorCode 错误码
type ErrorCode int

// 更具注释中的内容自动生产错误码对应的信息
//go:generate stringer -type ErrorCode -linecomment

const (
    SUCCESS ErrorCode = iota + 200000000 // Ok!
)

const (
    CreateSuccess ErrorCode = iota + 201000000 // Create Successful!
)

// 4xx000xxx 客户端错误
const (
    ERROR          ErrorCode = iota + 400000000 // Error!
    ParamBindError                              // There was an error with the parameters provided.
)

// 4xx001xxx User 模块错误
const (
    UsernameIsNotExist    ErrorCode = iota + 400001000 // The entered username does not exist.
    PasswordVerifyFailed                               // The password you entered is incorrect. Please try again.
    UsernameAlreadyExists                              // The username already exists.
    EmailAlreadyInUse                                  // The email address is already in use.
    RegisterInfoMismatch                               // Information mismatch.
    VerifyCodeExpired                                  // The Verification code does not exist or has expired.
    RefreshTokenInvalid                                // The refresh token is invalid.
)

const (
    AuthFailed            ErrorCode = iota + 401001000 // Auth failed.
    AuthTokenNULL                                      // No authorization token found.
    AuthTokenExpired                                   // Auth token is expired.
    AuthTokenNotValidYet                               // Auth token is not valid.
    AuthTokenMalformed                                 // Auth token malformed.
    AuthTokenInvalid                                   // Auth token is invalid.
    AuthTokenCreateFailed                              // Token create failed.
)

// 数据库 MySQL 错误
const (
    DatabaseError       ErrorCode = iota + 500001000 // MySQL Database Error.
    QueryPostListFailed                              // Unable to Fetch Post List.
)

const (
    RedisError ErrorCode = iota + 500002000 // Redis Error.
)

// IsSuccess 判断错误码是否为成功
func (code ErrorCode) IsSuccess() bool {
    return code >= 200000000 && code < 300000000
}

// IsClientError 判断错误码是否属于客户端错误
func (code ErrorCode) IsClientError() bool {
    return code >= 400000000 && code < 500000000
}

// IsSystemError 判断错误码是否属于系统错误
func (code ErrorCode) IsSystemError() bool {
    return code >= 500000000 && code < 600000000
}

// Status 返回该错误码对应的 HTTP 状态码
func (code ErrorCode) Status() int {
    return int(code) / 1000000
}
