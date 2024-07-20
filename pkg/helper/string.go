package helper

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	digitLetterBytes         = "0123456789"
	alphaLetterBytes         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitAndAlphaLetterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allLetterBytes           = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ()`!@#$%^&*_-+=|{}[]:;'<>,.?"
)

var (
	letterIdxBits int64
	letterIdxMask int64
	letterIdxMax  int64
)

type MakeStrMode int

const (
	// Alpha 仅包含大小写字母
	Alpha MakeStrMode = iota
	// Digit 仅包含数字
	Digit
	// DigitAlpha 包含数字和大小写字母
	DigitAlpha
	// DigitAlphaPunct 包含数字、字母、特殊符号
	DigitAlphaPunct
)

var src = rand.NewSource(time.Now().UnixNano())

func makeStr(length int, letterBytes string) string {
	calMakeStrInfo()
	str := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			str[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&str))
}

func calMakeStrInfo() {
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax = 63 / letterIdxBits
}

func MakeStr(length int, mode MakeStrMode) string {
	switch mode {
	case Digit:
		letterIdxBits = 4
		return makeStr(length, digitLetterBytes)
	case Alpha:
		letterIdxBits = 6
		return makeStr(length, alphaLetterBytes)
	case DigitAlpha:
		letterIdxBits = 6
		return makeStr(length, digitAndAlphaLetterBytes)
	case DigitAlphaPunct:
		letterIdxBits = 7
		return makeStr(length, allLetterBytes)
	default:
		letterIdxBits = 7
		return makeStr(length, allLetterBytes)
	}
}
