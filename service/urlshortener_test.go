package service

import (
	"URLShortener/model/po"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"URLShortener/common/define"
	"URLShortener/common/errorx"
	"URLShortener/model/bo"
	"URLShortener/repo/mock"
)

func TestURLShortenerService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("CreateShortenURL", t, func() {
		Convey("success", func() {
			ctx := context.Background()
			mockRepo := mock.NewMockIURLShortenerRepo(mockCtrl)
			urlId := "A1b2C3"
			mockRepo.EXPECT().CreateShortenURL(ctx, gomock.Any()).Return(&po.CreateShortenURLResponse{UrlId: urlId}, nil)

			s := &URLShortenerService{
				URLShortenerRepo: mockRepo,
			}

			boResp, err := s.CreateShortenURL(ctx, &bo.CreateShortenURLRequest{})
			So(boResp.UrlId, ShouldEqual, urlId)
			So(err, ShouldBeNil)
		})

		Convey("create hashId Failed", func() {
			ctx := context.Background()
			mockRepo := mock.NewMockIURLShortenerRepo(mockCtrl)
			mockRepo.EXPECT().CreateShortenURL(ctx, gomock.Any()).Return(nil, errorx.InsertDataBaseFailedError).MaxTimes(define.CreateShortenUrlRetryMax)

			s := &URLShortenerService{
				URLShortenerRepo: mockRepo,
			}

			_, err := s.CreateShortenURL(ctx, &bo.CreateShortenURLRequest{})
			So(err, ShouldBeError, errorx.CreateHashIdFailedError)
		})
	})
}
