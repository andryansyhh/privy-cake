package repository

import (
	"database/sql"
	"log"
	"privy/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var res = models.CakeResponse{
	ID:          1,
	Title:       "Lemon cheesecake",
	Description: "A cheesecake made of lemon",
	Rating:      7,
	Image:       "<https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg>",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

var req = models.CakeRequest{
	Title:       "Lemon cheesecake",
	Description: "A cheesecake made of lemon",
	Rating:      7,
	Image:       "<https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg>",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestPostCakeData(t *testing.T) {
	db, mock := NewMock()
	repository := &Repo{db}
	defer func() {
		repository.db.Close()
	}()

	query := `
	INSERT INTO
		cake(
			title,
			description,
			rating,
			image,
			created_at,
			updated_at		
		)
		VALUES(
			?, ?, ?, ?, now(), now()
		)
	`
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(req.Title, req.Description, req.Rating, req.Image).WillReturnResult(sqlmock.NewResult(0, 1))
	err := repository.PostCakeData(req)
	assert.NoError(t, err)
}

func TestGetAllCakes(t *testing.T) {
	db, mock := NewMock()
	repository := &Repo{db}
	defer func() {
		repository.db.Close()
	}()

	query := `
	SELECT
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at	
	FROM cake WHERE deleted_at is null
	`

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(res.ID, res.Title, res.Description, res.Rating, res.Image, res.CreatedAt, res.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	cakes, err := repository.GetAllCakes()
	assert.NotEmpty(t, cakes)
	assert.NoError(t, err)
	assert.Len(t, cakes, 1)
}

func TestGetCakeByID(t *testing.T) {
	db, mock := NewMock()
	repository := &Repo{db}
	defer func() {
		repository.db.Close()
	}()

	query := `
	SELECT
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at	
	FROM cake 
	WHERE deleted_at is null 
	AND id = ?
	`

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(res.ID, res.Title, res.Description, res.Rating, res.Image, res.CreatedAt, res.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	cakes, err := repository.GetCakeByID(int(res.ID))
	assert.NotEmpty(t, cakes)
	assert.NoError(t, err)
}

func TestUpdateCakeByID(t *testing.T) {
	db, mock := NewMock()
	repository := &Repo{db}
	defer func() {
		repository.db.Close()
	}()

	query := `
	UPDATE cake SET
			title = ?,
			description = ?,
			rating = ?,
			image = ?,
			updated_at	= now()	
	WHERE id = ?
	`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(res.ID, req.Title, req.Description, req.Rating, req.Image).WillReturnResult(sqlmock.NewResult(0, 1))
	err := repository.UpdateCakeByID(int(res.ID), req)
	assert.NoError(t, err)
}

func TestDeleteCakeByID(t *testing.T) {
	db, mock := NewMock()
	repository := &Repo{db}
	defer func() {
		repository.db.Close()
	}()

	query := `
	UPDATE cake SET deleted_at = now()
	WHERE id = ?
	`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(res.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	err := repository.DeleteCakeByID(int(res.ID))
	assert.NoError(t, err)
}
