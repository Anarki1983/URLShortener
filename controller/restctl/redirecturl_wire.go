//go:build wireinject
// +build wireinject

package restctl

import (
	"context"

	"github.com/google/wire"

	"URLShortener/repo"
	"URLShortener/service"
)

// provider set of RedirectURL
var RedirectURLSet = wire.NewSet(
	NewRedirectURLController,
	service.NewRedirectURLService,
	wire.Bind(new(service.IRedirectURLService), new(*service.RedirectURLService)), // bind interface with service imp
	repo.NewRedirectURLRepo,
	wire.Bind(new(repo.IRedirectURLRepo), new(*repo.RedirectURLRepo)), // bind interface with repo imp
)

func initRedirectURLController(ctx context.Context) (*RedirectURLController, error) {
	wire.Build(RedirectURLSet)

	return &RedirectURLController{}, nil
}
