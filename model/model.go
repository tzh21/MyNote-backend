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
	Avatar   []byte `json:"avatar"`
}
