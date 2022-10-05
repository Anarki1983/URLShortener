package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"URLShortener/common/errorx"
	"URLShortener/model/bo"
	"URLShortener/model/po"
	"URLShortener/repo/mock"
)

func TestRedirectURLService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("GetOriginURL", t, func() {
		Convey("success", func() {
			ctx := context.Background()
			mockRepo := mock.NewMockIRedirectURLRepo(mockCtrl)
			url := "https://google.com"
			mockRepo.EXPECT().GetOriginURL(ctx, gomock.Any()).Return(&po.RedirectURLResponse{Url: url}, nil)

			s := &RedirectURLService{
				RedirectURLRepo: mockRepo,
			}

			boResp, err := s.GetOriginURL(ctx, &bo.RedirectURLRequest{})
			So(boResp.Url, ShouldEqual, url)
			So(err, ShouldBeNil)
		})

		Convey("url not found error", func() {
			ctx := context.Background()
			mockRepo := mock.NewMockIRedirectURLRepo(mockCtrl)
			mockRepo.EXPECT().GetOriginURL(ctx, gomock.Any()).Return(nil, errorx.UrlNotFoundError)

			s := &RedirectURLService{
				RedirectURLRepo: mockRepo,
			}

			_, err := s.GetOriginURL(ctx, &bo.RedirectURLRequest{})
			So(err, ShouldBeError, errorx.UrlNotFoundError)
		})

		Convey("fetch database failed error", func() {
			ctx := context.Background()
			mockRepo := mock.NewMockIRedirectURLRepo(mockCtrl)
			mockRepo.EXPECT().GetOriginURL(ctx, gomock.Any()).Return(nil, errorx.FetchDatabaseFailedError)

			s := &RedirectURLService{
				RedirectURLRepo: mockRepo,
			}

			_, err := s.GetOriginURL(ctx, &bo.RedirectURLRequest{})
			So(err, ShouldBeError, errorx.FetchDatabaseFailedError)
		})
	})
}
