package restctl

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"URLShortener/controller/restctl/transform"
	"URLShortener/log"
	"URLShortener/service"
)

const URLShortenerControllerKey = "URLShortenerController"

func init() {
	ctl, err := initURLShortenerController(context.Background())
	if err != nil {
		log.Panic(context.Background(), "%v", err)
	}

	registerController(URLShortenerControllerKey, ctl)
}

type URLShortenerController struct {
	URLShortenerService service.IURLShortenerService
}

// wire provider
func NewURLShortenerController(URLShortenerService service.IURLShortenerService) *URLShortenerController {
	return &URLShortenerController{
		URLShortenerService: URLShortenerService,
	}
}

func (ctl *URLShortenerController) SetupRouters(router *gin.Engine) {
	api := router.Group("api")
	apiv1 := api.Group("v1")
	apiv1.POST("/urls", ctl.CreateShortenURL)
}

func (ctl *URLShortenerController) CreateShortenURL(ginCtx *gin.Context) {
	// bind params to dto then convert to bo
	boReq, err := transform.CreateShortenURLReq(ginCtx)
	if err != nil {
		JSONError(ginCtx, err)
		return
	}

	// process create shorten url
	boResp, err := ctl.URLShortenerService.CreateShortenURL(ginCtx, boReq)
	if err != nil {
		JSONError(ginCtx, err)
		return
	}

	// transform bo to dto
	dtoResp, err := transform.CreateShortenURLResp(ginCtx, boResp)
	if err != nil {
		JSONError(ginCtx, err)
		return
	}

	ginCtx.JSON(http.StatusOK, dtoResp)
}
