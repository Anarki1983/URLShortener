package repo

import (
	"context"

	lru "github.com/hashicorp/golang-lru"

	"URLShortener/common/errorx"
	"URLShortener/model/po"
	"URLShortener/repo/cacheHelper"
	"URLShortener/repo/redisHelper"
)

// mockgen -source ./urlshortener.go -destination ./mock/urlshortener_mock.go -package mock
type IURLShortenerRepo interface {
	CreateShortenURL(ctx context.Context, poReq *po.CreateShortenURLRequest) (poResp *po.CreateShortenURLResponse, err *errorx.ServiceError)
}

func NewURLShortenerRepo() *URLShortenerRepo {
	return &URLShortenerRepo{}
}

type URLShortenerRepo struct {
	lruCache *lru.Cache
}

func (s *URLShortenerRepo) CreateShortenURL(ctx context.Context, poReq *po.CreateShortenURLRequest) (poResp *po.CreateShortenURLResponse, err *errorx.ServiceError) {
	success := false
	success, err = redisHelper.SetNX(ctx, poReq.UrlId, poReq.Url, poReq.Duration)
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errorx.CreateHashIdFailedError
	}

	// add to cache
	cacheHelper.Add(ctx, poReq.UrlId, poReq.Url)

	return &po.CreateShortenURLResponse{
		UrlId: poReq.UrlId,
	}, nil
}
