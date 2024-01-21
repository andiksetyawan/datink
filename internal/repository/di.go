package repository

import (
	"fmt"

	"go.uber.org/dig"

	"datink/internal/repository/packages"
	"datink/internal/repository/swipe"
	"datink/internal/repository/user"
	"datink/internal/repository/userPackages"
)

func Register(container *dig.Container) error {
	// user
	if err := container.Provide(user.NewRepository); err != nil {
		return fmt.Errorf("failed to provide user respository: %w", err)
	}

	// swipe
	if err := container.Provide(swipe.NewRepository); err != nil {
		return fmt.Errorf("failed to provide swipe respository: %w", err)
	}

	// packages
	if err := container.Provide(packages.NewRepository); err != nil {
		return fmt.Errorf("failed to provide packages respository: %w", err)
	}

	// swipe
	if err := container.Provide(userPackages.NewRepository); err != nil {
		return fmt.Errorf("failed to provide userPackage respository: %w", err)
	}

	return nil
}
