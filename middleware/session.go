package middleware

import (
	"fmt"
	"vvblog/config"
	"vvblog/vlog"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	redisCfg := config.Redis
	store, err := redis.NewStoreWithDB(10, "tcp", fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port),
		redisCfg.Secret, fmt.Sprintf("%d", redisCfg.Database), []byte("xxx"))
	if err != nil {
		vlog.Fatalf("[Session]创建session的redis存储出错：%+v", err)
	}
	return sessions.Sessions(config.App.Name+"_sessions", store)
}
