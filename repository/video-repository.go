package repository

import (
	"example/web-service-gin/entity"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type VideoRepository interface {
	Save(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *sqlx.DB
}

func NewVideoRepository() VideoRepository {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entity.Video) error {
	_, err := db.NamedExec(`
		INSERT INTO videos (title, description, url, person_id, created_at, updated_at)
		VALUES (:title, :description, :url, :person_id, :created_at, :updated_at)`,
		&video)
	if err != nil {
		return err
	}

	return nil
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	err := db.Select(&videos, "SELECT * FROM videos")
	if err != nil {
		panic("Failed to close database")
	}

	return videos
}
