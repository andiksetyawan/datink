package swipe

import (
	"go.uber.org/dig"

	"datink/internal/resource"
	"datink/internal/service"
)

type Handler struct {
	dig.In

	Resource resource.Resource
	Service  service.Service
}
