package request

type LoginByUsernameAndPasswordRequest struct {
    Username string `json:"username" binding:"required,min=3,max=16,username-charset"`
    Password string `json:"password" binding:"required,min=8,max=16,password-charset"`
}

type GetRegisterVerifyCodeWithEmailRequest struct {
    Username string `json:"username" binding:"required,min=3,max=16,username-charset"`
    Password string `json:"password" binding:"required,min=8,max=16,password-charset"`
    Email    string `json:"email" binding:"required,min=4,max=256,email"`
}

type CreateUserByEmailVerifyCodeRequest struct {
    Username   string `json:"username" binding:"required,min=3,max=16,username-charset"`
    Password   string `json:"password" binding:"required,min=8,max=16,password-charset"`
    Email      string `json:"email" binding:"required,min=4,max=256,email"`
    VerifyCode string `json:"verifyCode" binding:"required,min=6,max=6"`
    RequestID  string `json:"requestID" binding:"required"`
}
