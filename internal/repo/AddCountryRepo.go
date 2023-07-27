package repo

import (
	"context"
	"database/sql"
	"time"

	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type AddCountryRepo struct {
	query string
	dbs   *sql.DB
}

func NewAddCountryRepo(db *sql.DB) contract.AddCountryRepoInterface {
	return &AddCountryRepo{
		query: "INSERT INTO Country(CountryName, ProvinceName) values(nullif(?,''),nullif(?,''))",
		dbs:   db,
	}
}

func (repo *AddCountryRepo) AddCountry(data model.Country) (lastinserted int64, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := trx.ExecContext(timeoutctx, repo.query, data.CountryName, data.Region)
	if err != nil {
		panic(err)
	}

	lastinserted, err = res.LastInsertId()

	return
}
