package grpc

import (
	"context"
	"errors"
	"fmt"

	userpb "github.com/AnnuCode/toy-grpc/gen/proto/user/v1"
	"github.com/AnnuCode/toy-grpc/internal/controller/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	userpb.UnimplementedGetUserServiceServer
	userpb.UnimplementedGetUsersServiceServer
	ctrl *user.Controller
}

func New(ctrl *user.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}
func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if req == nil || req.GetId() <= 0 {
		return nil, fmt.Errorf("invalid req or User Id")
	}
	u, err := h.ctrl.GetUser(ctx, req.GetId())
	if err != nil && errors.Is(err, user.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &userpb.GetUserResponse{User: u.User}, nil

}
func (h *Handler) GetUsers(ctx context.Context, req *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil req")
	}
	for _, id := range req.GetId() {
		if id <= 0 {
			return nil, fmt.Errorf("invalid req or User Id")
		}
	}
	u, err := h.ctrl.GetUsers(ctx, req.GetId())
	if err != nil && errors.Is(err, user.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &userpb.GetUsersResponse{Users: u.Users}, nil

}
