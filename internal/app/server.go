package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"go.uber.org/dig"
	"go.uber.org/zap"

	"datink/internal/delivery"
	"datink/internal/resource"
)

func RunServer() {
	invoke := func(container *dig.Container) error {
		return container.Invoke(func(infrastructure delivery.Package, resource resource.Resource) error {
			if err := infrastructure.Register(resource.Echo); err != nil {
				resource.Logger.Error("error on register http")
				return err
			}

			go func() {
				if err := resource.Echo.Start(fmt.Sprintf(":%d", resource.Config.AppPort)); err != nil {
					resource.Logger.Error("http server interrupted", zap.Error(err))
				} else {
					resource.Logger.Info("starting apps")
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)

			<-quit
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := resource.Echo.Shutdown(ctx); err != nil {
				return err
			}

			resource.Logger.Info("done, server exited properly :)")
			return nil
		})
	}

	onError := func(container *dig.Container, err error) {
		_ = container.Invoke(func() error {
			if err != nil {
				panic(err)
			}
			return nil
		})
	}

	container, err := Container()
	if err != nil {
		panic(err)
	}

	if err := invoke(container); err != nil {
		onError(container, err)
	}

	return
}
