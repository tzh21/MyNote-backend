package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
	gorm.Model
	Username string `json:"username"`
	Motto    string `json:"motto"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"` // 头像文件名。图片需要另外请求
}
