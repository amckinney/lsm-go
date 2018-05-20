package lsm

import (
	"context"

	"github.com/amckinney/lsm-go/gen/idl/lsm-go"
)

// Handler is a LSM Storage handler.
type Handler = lsmgopb.LSMGoServer

type handler struct{}

// New returns a new LSM Storage Handler.
func New() (Handler, error) {
	return &handler{}, nil
}

func (h *handler) Get(context.Context, *lsmgopb.GetRequest) (*lsmgopb.GetResponse, error) { return nil, nil }

func (h *handler) Set(context.Context, *lsmgopb.SetRequest) (*lsmgopb.SetResponse, error) { return nil, nil }
