package redis_model

import "encoding/json"

type RegisterUserInfo struct {
    Username   string `json:"username"`
    Password   string `json:"password"`
    Email      string `json:"email"`
    VerifyCode string `json:"verifyCode"`
}

func (registerUserInfo *RegisterUserInfo) MarshalBinary() (data []byte, err error) {
    return json.Marshal(registerUserInfo)
}

func (registerUserInfo *RegisterUserInfo) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, registerUserInfo)
}
