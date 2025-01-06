package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 用于对密码进行哈希加密
func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), nil
}

// CompareHashPwdAndPwd 用于比较哈希密码和密码是否匹配
func CompareHashPwdAndPwd(hashedPwd, pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) { //密码不匹配
		return false, nil
	} else if err != nil { //其他错误
		return false, err
	} else { //密码匹配
		return true, nil
	}
}
