package main

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"URLShortener/controller/restctl"
	"URLShortener/log"
	"URLShortener/repo/redisHelper"
)

func main() {
	defer func() {
		log.Error(context.Background(), "Server shutdown...")
		if err := recover(); err != nil {
			log.Error(context.Background(), "error: %v", err)
			log.Error(context.Background(), string(debug.Stack()))
		}
	}()

	if err := Run(); err != nil {
		log.Error(context.Background(), "%v", err)
	}
}

func Run() error {
	stop := make(chan error, 1)
	ch := make(chan os.Signal, 1)

	// init rand seed
	rand.Seed(time.Now().UnixNano())

	log.Init()

	// init infra
	redisHelper.Init()

	ctx := context.Background()
	if router, err := restctl.InitController(ctx); err != nil {
		return err
	} else {
		go run(router, stop)
	}

	log.Info(ctx, "Server Start")

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	var err error
	select {
	case err = <-stop:
		gracefulShutdown(ctx)
	case <-ch:
		gracefulShutdown(ctx)
	}

	return err
}

func run(router *gin.Engine, stop chan error) {
	if err := router.Run(":80"); err != nil {
		stop <- errors.New(" Doesn't has valid port. ")
	}
}

func gracefulShutdown(ctx context.Context) {
	redisHelper.Stop()
}
