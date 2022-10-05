package transform

import (
	"URLShortener/common/errorx"
	"URLShortener/model/dto"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestURLShortenerTransform(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("createShortenURLReq", t, func() {
		Convey("success", func() {
			url := "https://abc.org"
			expiredAt := time.Now().Add(time.Second).Format(time.RFC3339)
			dtoReq := &dto.CreateShortenURLRequest{
				Url:       url,
				ExpiredAt: expiredAt,
			}

			boReq, err := createShortenURLReq(dtoReq)
			So(boReq.Url, ShouldEqual, url)
			So(err, ShouldBeNil)
		})

		Convey("invalid url format", func() {
			url := "abc.org"
			expiredAt := time.Now().Add(time.Second).Format(time.RFC3339)
			dtoReq := &dto.CreateShortenURLRequest{
				Url:       url,
				ExpiredAt: expiredAt,
			}

			_, err := createShortenURLReq(dtoReq)
			So(err, ShouldEqual, errorx.InvalidParameterError)
		})

		Convey("invalid expireAt format", func() {
			url := "https://abc.org"
			expiredAt := time.Now().Add(time.Second).Format(time.RFC822)
			dtoReq := &dto.CreateShortenURLRequest{
				Url:       url,
				ExpiredAt: expiredAt,
			}

			_, err := createShortenURLReq(dtoReq)
			So(err, ShouldEqual, errorx.InvalidParameterError)
		})

		Convey("invalid expireAt time", func() {
			url := "https://abc.org"
			expiredAt := time.Now().AddDate(0, 0, -1).Format(time.RFC3339)
			dtoReq := &dto.CreateShortenURLRequest{
				Url:       url,
				ExpiredAt: expiredAt,
			}

			_, err := createShortenURLReq(dtoReq)
			So(err, ShouldEqual, errorx.InvalidParameterError)
		})
	})
}
