package service

import (
	"time"
	"vvblog/config"
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/model"

	jwt "github.com/dgrijalva/jwt-go"
)

var Auth = new(auth)

type auth struct{}

func (auth *auth) GenToken(username string) (string, error) {
	c := model.UserClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.App.TokenEffectiveDuration) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    config.App.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(config.App.TokenSecret))
	if err != nil {
		return "", verror.NewCode(vcode.CodeInternalError, "生成JWT出错")
	}
	return tokenString, nil
}

func (auth *auth) ParseToken(tokenString string) (*model.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.App.TokenSecret), nil
	})
	if err != nil {
		return nil, verror.WrapCode(vcode.CodeInternalError, err, "解析JWT出错")
	}
	if userClaim, ok := token.Claims.(*model.UserClaim); ok && token.Valid {
		return userClaim, nil
	}
	return nil, verror.NewCode(vcode.CodeNotAuthorized, "JWT失效")
}
