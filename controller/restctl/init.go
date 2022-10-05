package restctl

import (
	"URLShortener/common/errorx"
	"URLShortener/log"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IController interface {
	SetupRouters(router *gin.Engine)
}

var controllerMap map[string]IController

func registerController(key string, controller IController) {
	if controllerMap == nil {
		controllerMap = make(map[string]IController)
	}

	if _, ok := controllerMap[key]; !ok {
		controllerMap[key] = controller
	} else {
		panic(fmt.Sprintf("register with duplicate key[%s]", key))
	}
}

func InitController(ctx context.Context) (router *gin.Engine, err error) {
	router = gin.Default()
	router.Use(cors.New(CorsConfig()))
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return router, err
	}

	for _, c := range controllerMap {
		c.SetupRouters(router)
	}

	return router, nil
}

func CorsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Origin"}
	return corsConf
}

func JSONError(ginCtx *gin.Context, err *errorx.ServiceError) {
	log.Error(ginCtx, "%v", err)
	ginCtx.AbortWithStatusJSON(err.Status, err)
}
