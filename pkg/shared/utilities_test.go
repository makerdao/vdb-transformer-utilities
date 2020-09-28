package shared_test

import (
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utilities", func() {
	Describe("converting int256 hex to big int", func() {
		It("correctly converts positive number", func() {
			result := shared.ConvertInt256HexToBigInt("0x00000000000000000000000000000000000000000000000007a1fe1602770000")

			Expect(result.String()).To(Equal("550000000000000000"))
		})

		It("correctly converts negative number", func() {
			result := shared.ConvertInt256HexToBigInt("0xffffffffffffffffffffffffffffffffffffffffffffffffff4e5d43d13b0000")

			Expect(result.String()).To(Equal("-50000000000000000"))
		})

		It("correctly converts another negative number", func() {
			result := shared.ConvertInt256HexToBigInt("0xfffffffffffffffffffffffffffffffffffffffffffffffffe9cba87a2760000")

			Expect(result.String()).To(Equal("-100000000000000000"))
		})
	})

	Describe("GetFullTableName", func() {
		It("Concatenates a schema and table name", func() {
			schema := "schema_name"
			table := "table_name"
			fullTableName := shared.GetFullTableName(schema, table)
			Expect(fullTableName).To(Equal("schema_name.table_name"))
		})
	})
})
