package lsm

import "go.uber.org/fx"

// Module provides an LSM Handler into an Fx application.
var Module = fx.Provide(New)
