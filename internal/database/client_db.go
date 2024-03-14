package database

import (
	"database/sql"
	"walletcore/internal/entity"
)

type ClientDb struct {
	Db *sql.DB
}

func NewClientDb(db *sql.DB) *ClientDb {
	return &ClientDb{
		Db: db,
	}
}

func (c *ClientDb) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.Db.Prepare("select id, name, email, created_at from clients where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *ClientDb) Save(client *entity.Client) error {
	stmt, err := c.Db.Prepare("insert into clients (id, name, email, created_at) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
