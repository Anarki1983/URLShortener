package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/speps/go-hashids/v2"

	"URLShortener/common/define"
	"URLShortener/common/errorx"
	"URLShortener/model/bo"
	"URLShortener/model/po"
	"URLShortener/repo"
)

type IURLShortenerService interface {
	CreateShortenURL(ctx context.Context, boReq *bo.CreateShortenURLRequest) (boResp *bo.CreateShortenURLResponse, err *errorx.ServiceError)
}

func NewURLShortenerService(URLShortenerRepo repo.IURLShortenerRepo) *URLShortenerService {
	return &URLShortenerService{
		URLShortenerRepo: URLShortenerRepo,
	}
}

type URLShortenerService struct {
	URLShortenerRepo repo.IURLShortenerRepo
}

func (s *URLShortenerService) CreateShortenURL(ctx context.Context, boReq *bo.CreateShortenURLRequest) (boResp *bo.CreateShortenURLResponse, err *errorx.ServiceError) {
	poResp := &po.CreateShortenURLResponse{}
	poReq := &po.CreateShortenURLRequest{
		Url:      boReq.Url,
		Duration: boReq.ExpiredAt.Sub(time.Now()),
	}

	success := false
	for i := 0; i < define.CreateShortenUrlRetryMax; i++ {
		poReq.UrlId = genUrlId()
		if poResp, err = s.URLShortenerRepo.CreateShortenURL(ctx, poReq); err == nil {
			success = true
			break
		}
	}

	if !success {
		return nil, errorx.CreateHashIdFailedError
	}

	return &bo.CreateShortenURLResponse{
		UrlId: poResp.UrlId,
	}, nil
}

func genUrlId() string {
	hd := hashids.NewData()
	hd.Salt = "URLShortener"
	hd.MinLength = define.UrlIdMaxLength
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{rand.Int(), rand.Int(), rand.Int(), rand.Int()})

	return e[:define.UrlIdMaxLength]
}
