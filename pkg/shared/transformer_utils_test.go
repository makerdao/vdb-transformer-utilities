package shared_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/makerdao/vdb-transformer-utilities/pkg/fakes"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shared transformer utils", func() {
	Describe("VerifyLog", func() {
		It("returns err if log is missing topics", func() {
			log := types.Log{Data: fakes.FakeHash().Bytes()}

			err := shared.VerifyLog(log, 1, true)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(shared.ErrLogMissingTopics(1, 0)))
		})

		It("returns error if log has fewer than required number of topics", func() {
			log := types.Log{
				Data: fakes.FakeHash().Bytes(),
				Topics: []common.Hash{
					fakes.FakeHash(),
				},
			}

			err := shared.VerifyLog(log, 2, true)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(shared.ErrLogMissingTopics(2, 1)))
		})

		It("returns err if log is missing required data", func() {
			log := types.Log{
				Topics: []common.Hash{{}, {}, {}, {}},
			}

			err := shared.VerifyLog(log, 4, true)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(shared.ErrLogMissingData))
		})

		It("does not return error for missing data if data not required", func() {
			log := types.Log{
				Topics: []common.Hash{{}, {}, {}, {}},
			}

			err := shared.VerifyLog(log, 4, false)

			Expect(err).NotTo(HaveOccurred())
		})

		It("does not return error for valid log", func() {
			log := types.Log{
				Topics: []common.Hash{{}, {}},
				Data:   fakes.FakeHash().Bytes(),
			}

			err := shared.VerifyLog(log, 2, true)

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
