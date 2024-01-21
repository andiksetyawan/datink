package handler

import (
	"go.uber.org/dig"

	"datink/internal/delivery/rest/handler/packages"
	"datink/internal/delivery/rest/handler/swipe"
	"datink/internal/delivery/rest/handler/user"
	"datink/internal/delivery/rest/handler/userPackage"
)

type Handler struct {
	dig.In

	User        user.Handler
	Swipe       swipe.Handler
	Packages    packages.Handler
	UserPackage userPackage.Handler
}
