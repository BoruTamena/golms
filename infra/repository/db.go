package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/BoruTamena/internal/core/port/repository"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbName   = "usermanagement"
)

type database struct {
	*sql.DB
}

func NewDB(cfg *viper.Viper) (repository.DataBase, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Get("HOST"), cfg.Get("HOST"), cfg.Get("PORT"), cfg.Get("DBUSER"), cfg.Get("DBNAME"))

	db, err := sql.Open("postgres", dsn)

	if err != nil {

		log.Fatal(err, err.Error())

		return nil, err
	}
	return &database{
		db,
	}, nil

}

func (da database) Close() error {
	return da.DB.Close()
}

func (da database) GetDB() *sql.DB {
	return da.DB
}
