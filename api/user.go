package api

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"time"
	"vvblog/config"
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/model"
	"vvblog/service"
	"vvblog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var User = new(userApi)

type userApi struct{}

func (u *userApi) Login(c *gin.Context) {
	var req model.UserLoginReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("%+v", err.Error())
		panic(verror.WrapCode(vcode.CodeInvalidParameter, err))
	}

	passHash := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))

	var user model.User
	if err := model.DB.Select("id, username").First(&user, "username = ? and password = ?", req.Username, passHash).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(verror.NewCode(vcode.CodeBusinessValidationFailed, "用户名或密码错误"))
		}

		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	tokenString, err := service.Auth.GenToken(req.Username)
	if err != nil {
		panic(err)
	}

	_, err = model.RDB.Set(context.Background(), "token_"+req.Username, tokenString, time.Duration(config.App.TokenEffectiveDuration)*time.Second).Result()
	if err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	if err := model.DB.Model(&user).Update("logined_at", time.Now()).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	utils.JsonSuccess(c, gin.H{
		"token": tokenString,
	})
}

func (u *userApi) Logout(c *gin.Context) {
	_, err := model.RDB.Del(context.Background(), "token_"+c.GetString("username")).Result()
	if err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	utils.JsonSuccess(c)
}
