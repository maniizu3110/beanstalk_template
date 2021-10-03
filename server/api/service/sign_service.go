package service

import (
	"context"
	"server/api/sqlc"
	"server/api/util/lib"
)

type SignService interface {
	CreateUser(params *sqlc.CreateUserParams) (*userResponse, error)
}

type signServiceImpl struct {
	store sqlc.Store
}

func NewSignService(store sqlc.Store) SignService {
	res := &signServiceImpl{}
	res.store = store
	return res
}


func (s *signServiceImpl) CreateUser(params *sqlc.CreateUserParams) (*userResponse, error) {
	//TODO:ここではまだhashedされてないのでPasswordで扱いたい
	hashedPassword, err := lib.HashPassword(params.HashedPassword)
	if err != nil {
		return &userResponse{}, err
	}
	newParams := sqlc.CreateUserParams{
		Name:       params.Name,
		HashedPassword: hashedPassword,
		Email:          params.Email,
	}
	user, err := s.store.CreateUser(context.Background(), newParams)
	if err != nil {
		return &userResponse{}, err
	}

	res := newUserResponse(user)
	return res, nil
}