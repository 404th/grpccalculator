package adder

import (
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return &CalcResponse{Res: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) Sub(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return &CalcResponse{Res: req.GetX() - req.GetY()}, nil
}

func (s *GRPCServer) Mult(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return &CalcResponse{Res: req.GetX() * req.GetY()}, nil
}

func (s *GRPCServer) Div(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return &CalcResponse{Res: req.GetX() / req.GetY()}, nil
}
