// Code generated by "stringer -type ErrorCode -linecomment"; DO NOT EDIT.

package error_code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SUCCESS-200000000]
	_ = x[ERROR-400000000]
	_ = x[ParamBindError-400000001]
	_ = x[UsernameIsNotExist-400001000]
	_ = x[PasswordVerifyFailed-400001001]
	_ = x[UsernameAlreadyExists-400001002]
	_ = x[EmailAlreadyInUse-400001003]
	_ = x[RegisterInfoMismatch-400001004]
	_ = x[VerifyCodeExpired-400001005]
	_ = x[AuthFailed-400001006]
	_ = x[AuthTokenNULL-400001007]
	_ = x[AuthTokenExpired-400001008]
	_ = x[AuthTokenNotValidYet-400001009]
	_ = x[AuthTokenMalformed-400001010]
	_ = x[AuthTokenInvalid-400001011]
	_ = x[AuthTokenCreateFailed-400001012]
	_ = x[DatabaseError-500001000]
	_ = x[QueryPostListFailed-500001001]
	_ = x[RedisError-500002000]
}

const (
	_ErrorCode_name_0 = "Ok!"
	_ErrorCode_name_1 = "Error!There was an error with the parameters provided."
	_ErrorCode_name_2 = "The entered username does not exist.The password you entered is incorrect. Please try again.The username already exists.The email address is already in use.Information mismatch.The Verification code does not exist or has expired.Auth failed.No authorization token found.Auth token is expired.Auth token is not valid.Auth token malformed.Auth token is invalid.Token create failed."
	_ErrorCode_name_3 = "MySQL Database Error.Unable to Fetch Post List."
	_ErrorCode_name_4 = "Redis Error."
)

var (
	_ErrorCode_index_1 = [...]uint8{0, 6, 54}
	_ErrorCode_index_2 = [...]uint16{0, 36, 92, 120, 156, 177, 229, 241, 270, 292, 316, 337, 359, 379}
	_ErrorCode_index_3 = [...]uint8{0, 21, 47}
)

func (i ErrorCode) String() string {
	switch {
	case i == 200000000:
		return _ErrorCode_name_0
	case 400000000 <= i && i <= 400000001:
		i -= 400000000
		return _ErrorCode_name_1[_ErrorCode_index_1[i]:_ErrorCode_index_1[i+1]]
	case 400001000 <= i && i <= 400001012:
		i -= 400001000
		return _ErrorCode_name_2[_ErrorCode_index_2[i]:_ErrorCode_index_2[i+1]]
	case 500001000 <= i && i <= 500001001:
		i -= 500001000
		return _ErrorCode_name_3[_ErrorCode_index_3[i]:_ErrorCode_index_3[i+1]]
	case i == 500002000:
		return _ErrorCode_name_4
	default:
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
