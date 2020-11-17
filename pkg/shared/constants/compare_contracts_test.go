package constants_test

import (
	"github.com/makerdao/vdb-transformer-utilities/pkg/fakes"
	"github.com/makerdao/vdb-transformer-utilities/pkg/shared/constants"
	"github.com/makerdao/vulcanizedb/pkg/eth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compare ABIs", func() {
	BeforeEach(func() {
		fakes.SetFakeConfig()
	})

	It("has no error if the ABIs are equal", func() {
		flipEthABI, parseErr := eth.ParseAbi(constants.GetContractABI("MCD_FLIP_ETH_A_1.0.0"))
		Expect(parseErr).NotTo(HaveOccurred())
		flipBatABI, parseErr := eth.ParseAbi(constants.GetContractABI("MCD_FLIP_BAT_A_1.0.0"))
		Expect(parseErr).NotTo(HaveOccurred())

		Expect(constants.CompareContractABI(flipEthABI, flipBatABI)).To(Succeed())
	})

	It("has a constructor error if ABIs differ on the constructor", func() {
		medianBalABI, parseErr := eth.ParseAbi(constants.GetContractABI("MEDIAN_BAL"))
		Expect(parseErr).NotTo(HaveOccurred())
		medianBatABI, parseErr := eth.ParseAbi(constants.GetContractABI("MEDIAN_BAT"))
		Expect(parseErr).NotTo(HaveOccurred())

		expectedError := constants.NewMismatchedConstructorsError(medianBalABI, medianBatABI)
		Expect(constants.CompareContractABI(medianBalABI, medianBatABI)).To(MatchError(expectedError))
	})

	It("has a method error if the ABIs have the same constructor but differing methods", func() {
		medianBalABI, parseErr := eth.ParseAbi(constants.GetContractABI("MEDIAN_BAL"))
		Expect(parseErr).NotTo(HaveOccurred())
		medianBalWithoutDenyABI, parseErr := eth.ParseAbi(constants.GetContractABI("MEDIAN_BAL_WITHOUT_DENY"))
		Expect(parseErr).NotTo(HaveOccurred())

		expectedError := constants.NewMismatchedMethodsError(medianBalABI, medianBalWithoutDenyABI, "deny")
		Expect(constants.CompareContractABI(medianBalABI, medianBalWithoutDenyABI)).To(MatchError(expectedError))
	})

	It("has an event error if the ABIs have the same constructor and methods but differing events", func() {
		medianBalABI, parseErr := eth.ParseAbi(constants.GetContractABI("MEDIAN_BAL"))
		Expect(parseErr).NotTo(HaveOccurred())
		medianBalWithoutLogMedianPriceABI, parseErr := eth.ParseAbi(constants.GetContractABI("MEDIAN_BAL_WITHOUT_LOG_MEDIAN_PRICE"))
		Expect(parseErr).NotTo(HaveOccurred())

		expectedError := constants.NewMismatchedEventsError(medianBalABI, medianBalWithoutLogMedianPriceABI, "LogMedianPrice")
		Expect(constants.CompareContractABI(medianBalABI, medianBalWithoutLogMedianPriceABI)).To(MatchError(expectedError))
	})

})
