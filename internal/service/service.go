package service

import (
	"go.uber.org/dig"

	"datink/internal/service/packages"
	"datink/internal/service/swipe"
	"datink/internal/service/user"
	"datink/internal/service/userPackage"
)

type Service struct {
	dig.In

	User        user.Service
	Swipe       swipe.Service
	Package     packages.Service
	UserPackage userPackage.Service
}
