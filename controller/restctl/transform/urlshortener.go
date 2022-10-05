package transform

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"URLShortener/common/define"
	"URLShortener/common/errorx"
	"URLShortener/model/bo"
	"URLShortener/model/dto"
)

func CreateShortenURLReq(ctx *gin.Context) (boReq *bo.CreateShortenURLRequest, err *errorx.ServiceError) {
	dtoReq := &dto.CreateShortenURLRequest{}
	if err := ctx.ShouldBindJSON(dtoReq); err != nil {
		return nil, errorx.BadRequestError
	}

	// verify url format
	if _, err := url.ParseRequestURI(dtoReq.Url); err != nil {
		return nil, errorx.InvalidParameterError
	}

	// verify expiredAt format
	expireAt, tErr := time.Parse(time.RFC3339, dtoReq.ExpiredAt)
	if tErr != nil {
		return nil, errorx.InvalidParameterError
	}
	// verify expiredAt is already expired
	if expireAt.Before(time.Now()) {
		return nil, errorx.InvalidParameterError
	}

	boReq = &bo.CreateShortenURLRequest{
		Url:       dtoReq.Url,
		ExpiredAt: expireAt,
	}

	return boReq, nil
}

func CreateShortenURLResp(ctx *gin.Context, bo *bo.CreateShortenURLResponse) (dtoResp *dto.CreateShortenURLResponse, err *errorx.ServiceError) {
	dtoResp = &dto.CreateShortenURLResponse{
		UrlId:      bo.UrlId,
		ShortenUrl: define.Domain + bo.UrlId,
	}

	return dtoResp, nil
}
