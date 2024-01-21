package application

import (
	"fmt"
	"sync"

	"go.uber.org/dig"

	"datink/config"
	"datink/internal/repository"
	"datink/internal/resource"
	"datink/internal/service"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	var outer error

	once.Do(func() {
		container = dig.New()

		if err := config.Register(container); err != nil {
			outer = err
			return
		}

		if err := resource.Register(container); err != nil {
			outer = err
			return
		}

		if err := repository.Register(container); err != nil {
			outer = err
			return
		}

		if err := service.Register(container); err != nil {
			outer = err
			return
		}
	})

	if outer != nil {
		return nil, fmt.Errorf("failed to initialize container: %w", outer)
	}

	return container, nil
}
