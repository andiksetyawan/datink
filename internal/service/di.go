package service

import (
	"fmt"

	"go.uber.org/dig"

	"datink/internal/service/packages"
	"datink/internal/service/swipe"
	"datink/internal/service/user"
	"datink/internal/service/userPackage"
)

func Register(container *dig.Container) error {
	// user
	if err := container.Provide(user.NewService); err != nil {
		return fmt.Errorf("failed to provide user service: %w", err)
	}

	// swipe
	if err := container.Provide(swipe.NewService); err != nil {
		return fmt.Errorf("failed to provide swipe service: %w", err)
	}

	// packages
	if err := container.Provide(packages.NewService); err != nil {
		return fmt.Errorf("failed to provide packages service: %w", err)
	}

	// user package
	if err := container.Provide(userPackage.NewService); err != nil {
		return fmt.Errorf("failed to provide userPackage service: %w", err)
	}

	return nil
}
