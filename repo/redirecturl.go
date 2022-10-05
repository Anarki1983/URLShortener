package repo

import (
	"context"

	"URLShortener/common/errorx"
	"URLShortener/model/po"
	"URLShortener/repo/redisHelper"
)

// mockgen -source ./redirecturl.go -destination ./mock/redirecturl_mock.go -package mock
type IRedirectURLRepo interface {
	GetOriginURL(ctx context.Context, poReq *po.RedirectURLRequest) (poResp *po.RedirectURLResponse, err *errorx.ServiceError)
}

func NewRedirectURLRepo() *RedirectURLRepo {
	return &RedirectURLRepo{}
}

type RedirectURLRepo struct {
}

func (s *RedirectURLRepo) GetOriginURL(ctx context.Context, poReq *po.RedirectURLRequest) (poResp *po.RedirectURLResponse, err *errorx.ServiceError) {
	var url string
	url, err = redisHelper.Get(ctx, poReq.UrlId)
	if err != nil {
		if err == errorx.DataNotFoundError {
			return nil, errorx.UrlNotFoundError
		}

		return nil, errorx.FetchDatabaseFailedError
	}

	return &po.RedirectURLResponse{
		Url: url,
	}, nil
}
