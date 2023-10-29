package user

import (
	"context"
	"errors"

	userpb "github.com/AnnuCode/toy-grpc/gen/proto/user/v1"
	"github.com/AnnuCode/toy-grpc/internal/repository"
)

var ErrNotFound = errors.New("not found")

type userRepository interface {
	GetUser(ctx context.Context, id int64) (*userpb.GetUserResponse, error)
	GetUsers(ctx context.Context, id []int64) (*userpb.GetUsersResponse, error)
}
type Controller struct {
	repo userRepository
}

func New(repo userRepository) *Controller {
	return &Controller{repo}
}
func (c *Controller) GetUser(ctx context.Context, id int64) (*userpb.GetUserResponse, error) {
	res, err := c.repo.GetUser(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, nil
}
func (c *Controller) GetUsers(ctx context.Context, id []int64) (*userpb.GetUsersResponse, error) {
	res, err := c.repo.GetUsers(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, nil
}
