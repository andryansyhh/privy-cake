package usecase

import "privy/repository"

type PrivyUsecase struct {
	repo repository.RepoInterface
}

type UsecaseInterface interface {
	CakeUsecase
}

func NewUsecase(repo repository.RepoInterface) UsecaseInterface {
	return &PrivyUsecase{
		repo: repo,
	}
}
