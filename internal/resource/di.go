package resource

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"

	"datink/internal/resource/external"
	"datink/pkg/util/httputil"
)

func Register(container *dig.Container) error {
	if err := container.Provide(external.NewEcho); err != nil {
		return errors.Wrap(err, "can not initialize Echo")
	}

	if err := container.Provide(external.NewLogger); err != nil {
		return errors.Wrap(err, "can not initialize logger")
	}

	if err := container.Provide(external.NewDatabase); err != nil {
		return errors.Wrap(err, "can not initialize database")
	}

	if err := container.Provide(httputil.NewHttpResponse); err != nil {
		return errors.Wrap(err, "can not initialize response helper")
	}

	return nil
}
