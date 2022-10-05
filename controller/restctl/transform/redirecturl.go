package transform

import (
	"github.com/gin-gonic/gin"

	"URLShortener/common/define"
	"URLShortener/common/errorx"
	"URLShortener/model/bo"
	"URLShortener/model/dto"
)

func RedirectURLReq(ctx *gin.Context) (boReq *bo.RedirectURLRequest, err *errorx.ServiceError) {
	dtoReq := &dto.RedirectURLRequest{}
	if err := ctx.ShouldBindUri(dtoReq); err != nil {
		return nil, errorx.BadRequestError
	}

	return redirectURLReq(dtoReq)
}

func redirectURLReq(dtoReq *dto.RedirectURLRequest) (boReq *bo.RedirectURLRequest, err *errorx.ServiceError) {
	// verify urlId format
	if len(dtoReq.UrlId) != define.UrlIdLength {
		return nil, errorx.InvalidParameterError
	}

	boReq = &bo.RedirectURLRequest{
		UrlId: dtoReq.UrlId,
	}

	return boReq, nil
}

func RedirectURLResp(ctx *gin.Context, bo *bo.RedirectURLResponse) (dtoResp *dto.RedirectURLResponse, err *errorx.ServiceError) {
	dtoResp = &dto.RedirectURLResponse{
		Url: bo.Url,
	}

	return dtoResp, nil
}
