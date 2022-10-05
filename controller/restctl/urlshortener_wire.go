//go:build wireinject
// +build wireinject

package restctl

import (
	"context"

	"github.com/google/wire"

	"URLShortener/repo"
	"URLShortener/service"
)

// provider set of URLShortener
var URLShortenerSet = wire.NewSet(
	NewURLShortenerController,
	service.NewURLShortenerService,
	wire.Bind(new(service.IURLShortenerService), new(*service.URLShortenerService)), // bind interface with service imp
	repo.NewURLShortenerRepo,
	wire.Bind(new(repo.IURLShortenerRepo), new(*repo.URLShortenerRepo)), // bind interface with repo imp
)

func initURLShortenerController(ctx context.Context) (*URLShortenerController, error) {
	wire.Build(URLShortenerSet)

	return &URLShortenerController{}, nil
}
