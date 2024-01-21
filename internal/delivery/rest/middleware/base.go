package middleware

import (
	"go.uber.org/dig"

	"datink/internal/resource"
)

type (
	// Middlewares .
	Middlewares struct {
		dig.In

		Resource resource.Resource
	}
)
