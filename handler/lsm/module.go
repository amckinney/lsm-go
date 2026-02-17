package lsm

import (
	"github.com/amckinney/lsm-go/storage"
	"go.uber.org/fx"
)

// Module provides an LSM Handler into an Fx application.
var Module = fx.Options(
	fx.Provide(
		// Provide an in-memory storage implementation
		// TODO: Replace with LSM storage once implemented
		func() storage.Storage {
			return storage.NewMemoryStorage()
		},
		New,
	),
)
