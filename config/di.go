package config

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewConfig); err != nil {
		return errors.Wrap(err, "can not initialize config")
	}
	return nil
}
