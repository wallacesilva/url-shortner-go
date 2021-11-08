package repositories

import (
	"database/sql"

	"github.com/wallacesilva/url-shortnet-go/app/models"
)

type UrlRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (repo *UrlRepository) Get(id int64) (*models.Url, error) {
	query := `
		SELECT url, code
		FROM urls
		WHERE id = ?;
	`

	row := repo.db.QueryRow(query, id)

	var url string
	var code string

	err := row.Scan(&url, &code)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	urlShorten := &models.Url{
		ID:   id,
		Url:  url,
		Code: code,
	}

	return urlShorten, nil
}

func (repo *UrlRepository) Create(url *models.Url) (*models.Url, error) {
	query := `
		INSERT INTO url (url, code, title)
		VALUES (?, ?, ?);
	`

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(url.Url, url.Code)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdUrlShorten, err := repo.Get(lastID)
	if err != nil {
		return nil, err
	}

	return createdUrlShorten, nil
}
