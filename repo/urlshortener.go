package repo

import (
	"URLShortener/common/errorx"
	"URLShortener/model/po"
	"URLShortener/repo/redisHelper"
	"context"
)

// mockgen -source ./urlshortener.go -destination ./mock/urlshortener_mock.go -package mock
type IURLShortenerRepo interface {
	CreateShortenURL(ctx context.Context, poReq *po.CreateShortenURLRequest) (poResp *po.CreateShortenURLResponse, err *errorx.ServiceError)
}

func NewURLShortenerRepo() *URLShortenerRepo {
	return &URLShortenerRepo{}
}

type URLShortenerRepo struct {
}

func (s *URLShortenerRepo) CreateShortenURL(ctx context.Context, poReq *po.CreateShortenURLRequest) (poResp *po.CreateShortenURLResponse, err *errorx.ServiceError) {
	_, err = redisHelper.SetNX(ctx, poReq.UrlId, poReq.Url, poReq.Duration)
	if err != nil {
		return nil, err
	}

	return &po.CreateShortenURLResponse{
		UrlId: poReq.UrlId,
	}, nil
}
