package service

import (
	"context"
	"server/api/sqlc"
)

type UserService interface {
	ListUser(params sqlc.ListUserParams) ([]sqlc.User, error)
}

type userServiceImpl struct {
	store sqlc.Store
}

func NewUserService(store sqlc.Store) UserService {
	res := &userServiceImpl{}
	res.store = store
	return res
}

type userResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newUserResponse(user sqlc.User) *userResponse {
	return &userResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (s *userServiceImpl) ListUser(params sqlc.ListUserParams) ([]sqlc.User, error) {
	listUser, err := s.store.ListUser(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return listUser, nil

}
