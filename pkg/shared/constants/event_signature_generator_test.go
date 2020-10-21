package constants_test

import (
	"errors"

	"github.com/makerdao/vdb-transformer-utilities/pkg/fakes"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared/constants"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("event signature generator", func() {
	BeforeEach(func() {
		fakes.SetFakeConfig()
	})

	Describe("findSignatureInAbi", func() {
		It("returns the signature if it exists in the ABI", func() {
			abi := constants.GetContractABI("MCD_JUG")
			signature, _ := constants.FindSignatureInAbi(abi, "file", []string{"bytes32", "bytes32", "uint256"})

			Expect(signature).To(Equal("file(bytes32,bytes32,uint256)"))
		})

		It("returns error if signature not found in ABI", func() {
			expectedError := errors.New("method file(bytes32,bytes32) does not exist in ABI")
			abi := constants.GetContractABI("MCD_JUG")

			_, err := constants.FindSignatureInAbi(abi, "file", []string{"bytes32", "bytes32"})

			Expect(err).To(MatchError(expectedError))
		})
	})

	Describe("GetOverloadedFunctionSignature", func() {
		It("panics if it encounters an error", func() {
			abi := constants.GetContractABI("MCD_JUG")
			Expect(func() { constants.GetOverloadedFunctionSignature(abi, "file", []string{"bytes32", "bytes32"}) }).To(Panic())
		})
	})
})
