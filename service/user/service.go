package user

import (
	"github.com/dezh-tech/go-gin-boilerplate/entity"
)

type Repository interface {
	Add(usr entity.User) error
	GetByPubkey(pubkey string) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
