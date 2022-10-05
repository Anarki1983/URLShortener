package service

import (
	"context"

	"URLShortener/common/errorx"
	"URLShortener/model/bo"
	"URLShortener/model/po"
	"URLShortener/repo"
)

type IRedirectURLService interface {
	GetOriginURL(ctx context.Context, boReq *bo.RedirectURLRequest) (boResp *bo.RedirectURLResponse, err *errorx.ServiceError)
}

func NewRedirectURLService(RedirectURLRepo repo.IRedirectURLRepo) *RedirectURLService {
	return &RedirectURLService{
		RedirectURLRepo: RedirectURLRepo,
	}
}

type RedirectURLService struct {
	RedirectURLRepo repo.IRedirectURLRepo
}

func (s *RedirectURLService) GetOriginURL(ctx context.Context, boReq *bo.RedirectURLRequest) (boResp *bo.RedirectURLResponse, err *errorx.ServiceError) {
	poResp := &po.RedirectURLResponse{}
	poReq := &po.RedirectURLRequest{
		UrlId: boReq.UrlId,
	}

	poResp, err = s.RedirectURLRepo.GetOriginURL(ctx, poReq)
	if err != nil {
		return nil, err
	}

	return &bo.RedirectURLResponse{
		Url: poResp.Url,
	}, nil
}
