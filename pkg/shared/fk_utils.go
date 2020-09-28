package shared

import (
	"github.com/jmoiron/sqlx"
	"github.com/makerdao/vulcanizedb/libraries/shared/repository"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres"
)

func GetOrCreateAddress(address string, db *postgres.DB) (int64, error) {
	return repository.GetOrCreateAddress(db, address)
}

func GetOrCreateAddressInTransaction(address string, tx *sqlx.Tx) (int64, error) {
	addressId, addressErr := repository.GetOrCreateAddressInTransaction(tx, address)
	return addressId, addressErr
}
