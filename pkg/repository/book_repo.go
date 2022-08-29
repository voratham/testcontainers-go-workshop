package repository

import (
	"github.com/jmoiron/sqlx"
)

type bookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) BookRepo {
	return &bookRepo{
		db: db,
	}
}

func (br *bookRepo) Create(b Book) (*Book, error) {
	query := `insert into book (name, price, author, description, image_url) values (?,?,?,?,?)`
	result, err := br.db.Exec(query, b.Name, b.Price, b.Author, b.Description, b.ImageURL)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	b.ID = id
	return &b, nil

}

func (br *bookRepo) GetByID(id int64) (*Book, error) {
	var book Book
	err := br.db.Get(&book, `select id, name, price, author, description, image_url from book where id= ? `, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (br *bookRepo) DeleteByID(id int64) error {
	_, err := br.db.Exec(`delete from book where id=?`, id)
	if err != nil {
		return err
	}
	return nil
}

func (br *bookRepo) Update(id int64, b Book) (*Book, error) {
	_, err := br.db.NamedExec(`update book SET name=:name, price=:price, author=:author, description=:description, image_url=:image_url`, &b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}