package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Model
	Username  string `gorm:"type:varchar(20);not null" json:"username"`
	Password  string `gorm:"type:varchar(64);not null" json:"password"`
	LoginedAt *time.Time
}

type UserLoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
