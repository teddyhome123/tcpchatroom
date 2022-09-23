package model

import (
	"errors"
)

//根據業務邏輯的需要 自定義錯誤

var (
	ERROR_USER_NOTEXISTS = errors.New("該用戶不存在")
	ERROR_USER_EXISTS = errors.New("該用戶已存在")
	ERROR_USER_PWD = errors.New("密碼錯誤")
)