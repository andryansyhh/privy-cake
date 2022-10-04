package repository

import (
	"privy/models"
	query "privy/repository/query"
)

type CakeRepo interface {
	PostCakeData(input models.CakeRequest) error
	GetAllCakes() (res []models.CakeResponse, err error)
	GetCakeByID(ID int) (res models.CakeResponse, err error)
	UpdateCakeByID(ID int, req models.CakeRequest) error
	DeleteCakeByID(ID int) error
}

func (b *Repo) PostCakeData(req models.CakeRequest) error {
	_, err := b.db.Exec(
		query.PostCake,
		req.Title,
		req.Description,
		req.Rating,
		req.Image,
	)

	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) GetAllCakes() (res []models.CakeResponse, err error) {
	var qry = query.GetAllCakes

	row, err := b.db.Query(
		qry,
	)
	if err != nil {
		return res, err
	}

	for row.Next() {
		temp := models.CakeResponse{}
		err = row.Scan(
			&temp.ID,
			&temp.Title,
			&temp.Description,
			&temp.Rating,
			&temp.Image,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}
	return res, nil
}

func (b *Repo) GetCakeByID(ID int) (res models.CakeResponse, err error) {
	err = b.db.QueryRow(
		query.GetCakeByID,
		ID,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
		&res.Rating,
		&res.Image,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (b *Repo) UpdateCakeByID(ID int, req models.CakeRequest) error {
	_, err := b.db.Exec(
		query.UpdateCakeByID,
		ID,
		req.Title,
		req.Description,
		req.Rating,
		req.Image,
	)

	if err != nil {
		return err
	}
	return nil
}

func (b *Repo) DeleteCakeByID(ID int) error {
	_, err := b.db.Exec(
		query.DeleteCakeByID,
		ID,
	)

	if err != nil {
		return err
	}
	return nil
}
