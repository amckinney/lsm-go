package server

import (
	"net"

	"github.com/amckinney/lsm-go/gen/idl/lsm-go"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Module starts a gRPC, LSM Storage server.
var Module = fx.Options(
	fx.Provide(
		func() (net.Listener, error) { return net.Listen("tcp", "0.0.0.0:8080") },
		grpc.NewServer,
	),
	fx.Invoke(
		lsmgopb.RegisterLSMGoServer,
		func(s *grpc.Server, l net.Listener) { s.Serve(l) },
	),
)
