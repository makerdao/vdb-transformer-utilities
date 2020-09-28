package shared_test

import (
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared"
	"github.com/makerdao/vulcanizedb/pkg/config"
	"github.com/makerdao/vulcanizedb/pkg/core"
	"github.com/makerdao/vulcanizedb/pkg/datastore/postgres"
	"github.com/makerdao/vulcanizedb/pkg/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Foreign key utils", func() {
	var db *postgres.DB

	BeforeEach(func() {
		node := core.Node{
			GenesisBlock: "GENESIS",
			NetworkID:    1,
			ID:           "b6f90c0fdd8ec9607aed8ee45c69322e47b7063f0bfb7a29c8ecafab24d0a22d24dd2329b5ee6ed4125a03cb14e57fd584e67f9e53e6c631055cbbd82f080845",
			ClientName:   "Geth/v1.7.2-stable-1db4ecdc/darwin-amd64/go1.9",
		}
		dbConfig := config.Database{
			Hostname: "localhost",
			Name:     "vulcanize_testing",
			Port:     5432,
		}
		var err error
		db, err = postgres.NewDB(dbConfig, node)
		Expect(err).NotTo(HaveOccurred())
		db.MustExec(`DELETE FROM public.addresses`)
	})

	Describe("GetOrCreateAddress", func() {
		It("creates an address record", func() {
			_, err := shared.GetOrCreateAddress(fakes.FakeAddress.Hex(), db)
			Expect(err).NotTo(HaveOccurred())

			var address string
			db.Get(&address, `SELECT address from addresses LIMIT 1`)
			Expect(address).To(Equal(fakes.FakeAddress.Hex()))
		})

		It("returns the id for an address that already exists", func() {
			//create the address record
			createAddressId, createErr := shared.GetOrCreateAddress(fakes.FakeAddress.Hex(), db)
			Expect(createErr).NotTo(HaveOccurred())

			//get the address record
			getAddressId, getErr := shared.GetOrCreateAddress(fakes.FakeAddress.Hex(), db)
			Expect(getErr).NotTo(HaveOccurred())

			Expect(createAddressId).To(Equal(getAddressId))

			var addressCount int
			db.Get(&addressCount, `SELECT count(*) from addresses`)
			Expect(addressCount).To(Equal(1))
		})
	})

	Describe("GetOrCreateAddressInTransaction", func() {
		It("creates an address record", func() {
			tx, txErr := db.Beginx()
			Expect(txErr).NotTo(HaveOccurred())

			_, createErr := shared.GetOrCreateAddressInTransaction(fakes.FakeAddress.Hex(), tx)
			Expect(createErr).NotTo(HaveOccurred())

			commitErr := tx.Commit()
			Expect(commitErr).NotTo(HaveOccurred())

			var address string
			db.Get(&address, `SELECT address from addresses LIMIT 1`)
			Expect(address).To(Equal(fakes.FakeAddress.Hex()))
		})

		It("returns the id for an address that already exists", func() {
			tx, txErr := db.Beginx()
			Expect(txErr).NotTo(HaveOccurred())

			//create the address record
			createAddressId, createErr := shared.GetOrCreateAddressInTransaction(fakes.FakeAddress.Hex(), tx)
			Expect(createErr).NotTo(HaveOccurred())

			//get the address record
			getAddressId, getErr := shared.GetOrCreateAddressInTransaction(fakes.FakeAddress.Hex(), tx)
			Expect(getErr).NotTo(HaveOccurred())

			commitErr := tx.Commit()
			Expect(commitErr).NotTo(HaveOccurred())

			Expect(createAddressId).To(Equal(getAddressId))

			var addressCount int
			db.Get(&addressCount, `SELECT count(*) from addresses`)
			Expect(addressCount).To(Equal(1))
		})

		It("doesn't persist the address if the transaction is rolled back", func() {
			tx, txErr := db.Beginx()
			Expect(txErr).NotTo(HaveOccurred())

			_, createErr := shared.GetOrCreateAddressInTransaction(fakes.FakeAddress.Hex(), tx)
			Expect(createErr).NotTo(HaveOccurred())

			commitErr := tx.Rollback()
			Expect(commitErr).NotTo(HaveOccurred())

			var addressCount int
			db.Get(&addressCount, `SELECT count(*) from addresses`)
			Expect(addressCount).To(Equal(0))
		})
	})
})
