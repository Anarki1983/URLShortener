package transform

import (
	"URLShortener/common/errorx"
	"URLShortener/model/dto"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRedirectURLTransform(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("redirectURLReq", t, func() {
		Convey("success", func() {
			urlId := "1A2b3C"
			dtoReq := &dto.RedirectURLRequest{
				UrlId: urlId,
			}

			boReq, err := redirectURLReq(dtoReq)
			So(boReq.UrlId, ShouldEqual, urlId)
			So(err, ShouldBeNil)
		})

		Convey("invalid urlId format", func() {
			urlId := "1A2b3C4D"
			dtoReq := &dto.RedirectURLRequest{
				UrlId: urlId,
			}

			_, err := redirectURLReq(dtoReq)
			So(err, ShouldEqual, errorx.InvalidParameterError)
		})
	})
}
