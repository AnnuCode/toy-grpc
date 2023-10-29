package memory

import (
	"context"
	"fmt"

	"sync"

	userpb "github.com/AnnuCode/toy-grpc/gen/proto/user/v1"
)

type Repository struct {
	sync.RWMutex
	data []*userpb.User
}

func New() *Repository {
	return &Repository{
		data: []*userpb.User{
			{
				Id:      1,
				Fname:   "Lionel Messi",
				City:    "Barcelona",
				Phone:   1234567891,
				Height:  5.7,
				Married: true,
			},
			{
				Id:      2,
				Fname:   "Cristiano Ronaldo",
				City:    "Madrid",
				Phone:   1234737891,
				Height:  6.1,
				Married: false,
			},
			{
				Id:      3,
				Fname:   "Neymar",
				City:    "Paris",
				Phone:   1234888891,
				Height:  5.9,
				Married: false,
			},
			{
				Id:      4,
				Fname:   "Robert Lewandowski",
				City:    "Munich",
				Phone:   1234544491,
				Height:  6.1,
				Married: true,
			},
		},
	}
}
func (r *Repository) GetUser(_ context.Context, id int64) (*userpb.GetUserResponse, error) {
	r.RLock()
	defer r.RUnlock()
	for _, user := range r.data {
		if user.Id == id {
			return &userpb.GetUserResponse{User: user}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
func (r *Repository) GetUsers(_ context.Context, id []int64) (*userpb.GetUsersResponse, error) {
	r.Lock()
	defer r.Unlock()

	response := []*userpb.User{}
	for _, id := range id {
		for _, user := range r.data {
			if user.Id == id {
				response = append(response, user)
				break
			}
		}
	}
	return &userpb.GetUsersResponse{Users: response}, nil
}
