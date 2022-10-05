package restctl

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"URLShortener/controller/restctl/transform"
	"URLShortener/log"
	"URLShortener/service"
)

const RedirectURLControllerKey = "RedirectURLController"

func init() {
	ctl, err := initRedirectURLController(context.Background())
	if err != nil {
		log.Panic(context.Background(), "%v", err)
	}

	registerController(RedirectURLControllerKey, ctl)
}

type RedirectURLController struct {
	RedirectURLService service.IRedirectURLService
}

func NewRedirectURLController(RedirectURLService service.IRedirectURLService) *RedirectURLController {
	return &RedirectURLController{
		RedirectURLService: RedirectURLService,
	}
}

func (ctl *RedirectURLController) SetupRouters(router *gin.Engine) {
	router.GET("/:url_id", ctl.Redirect)
}

func (ctl *RedirectURLController) Redirect(ginCtx *gin.Context) {
	boReq, err := transform.RedirectURLReq(ginCtx)
	if err != nil {
		JSONError(ginCtx, err)
		return
	}

	boResp, err := ctl.RedirectURLService.GetOriginURL(ginCtx, boReq)
	if err != nil {
		JSONError(ginCtx, err)
		return
	}

	dtoResp, err := transform.RedirectURLResp(ginCtx, boResp)
	if err != nil {
		JSONError(ginCtx, err)
		return
	}

	ginCtx.Redirect(http.StatusFound, dtoResp.Url)
}
