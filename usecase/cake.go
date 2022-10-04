package usecase

import (
	"privy/models"
	"privy/utils"
)

type CakeUsecase interface {
	PostCake(req models.CakeRequest) error
	GetAllCakes() (res []models.CakeResponse, err error)
	GetCakeByID(ID int) (res models.CakeResponse, err error)
	UpdateCakeByID(ID int, req models.CakeRequest) error
	DeleteCakeByID(ID int) error
}

func (b PrivyUsecase) PostCake(req models.CakeRequest) error {
	err := b.repo.PostCakeData(req)
	if err != nil {
		return err
	}
	return nil
}

func (b PrivyUsecase) GetAllCakes() (res []models.CakeResponse, err error) {
	res, err = b.repo.GetAllCakes()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (b PrivyUsecase) GetCakeByID(ID int) (res models.CakeResponse, err error) {
	res, err = b.repo.GetCakeByID(ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b PrivyUsecase) UpdateCakeByID(ID int, req models.CakeRequest) error {
	res, err := b.repo.GetCakeByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return utils.ErrIdNotFound
	}

	err = b.repo.UpdateCakeByID(ID, req)
	if err != nil {
		return err
	}

	return nil

}

func (b PrivyUsecase) DeleteCakeByID(ID int) error {
	res, err := b.repo.GetCakeByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return utils.ErrIdNotFound
	}

	err = b.repo.DeleteCakeByID(ID)
	if err != nil {
		return err
	}
	return nil
}
