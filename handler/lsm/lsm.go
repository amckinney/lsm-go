package lsm

import (
	"context"

	"github.com/amckinney/lsm-go/gen/idl/lsm"
	"github.com/amckinney/lsm-go/storage"
)

// Handler is a LSM Storage handler.
type Handler = lsmpb.LSMServer

type handler struct {
	storage storage.Storage
}

// New returns a new LSM Storage Handler.
func New(s storage.Storage) (Handler, error) {
	return &handler{
		storage: s,
	}, nil
}

func (h *handler) Get(ctx context.Context, req *lsmpb.GetRequest) (*lsmpb.GetResponse, error) {
	value, err := h.storage.Get(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return &lsmpb.GetResponse{
		Value: value,
	}, nil
}

func (h *handler) Set(ctx context.Context, req *lsmpb.SetRequest) (*lsmpb.SetResponse, error) {
	if err := h.storage.Set(ctx, req.GetKey(), req.GetValue()); err != nil {
		return nil, err
	}
	return &lsmpb.SetResponse{}, nil
}
