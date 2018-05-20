package lsm

import (
	"context"

	"github.com/amckinney/lsm-go/gen/idl/lsm"
)

// Handler is a LSM Storage handler.
type Handler = lsmpb.LSMServer

type handler struct{}

// New returns a new LSM Storage Handler.
func New() (Handler, error) {
	return &handler{}, nil
}

func (h *handler) Get(context.Context, *lsmpb.GetRequest) (*lsmpb.GetResponse, error) { return nil, nil }

func (h *handler) Set(context.Context, *lsmpb.SetRequest) (*lsmpb.SetResponse, error) { return nil, nil }
