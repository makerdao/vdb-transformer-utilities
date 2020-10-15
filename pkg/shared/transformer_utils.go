package shared

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/makerdao/vulcanizedb/libraries/shared/constants"
)

const (
	OneTopicRequired    = 1
	TwoTopicsRequired   = 2
	ThreeTopicsRequired = 3
	FourTopicsRequired  = 4
	LogDataRequired     = true
	LogDataNotRequired  = false
)

var (
	ErrLogMissingTopics = func(expectedNumTopics, actualNumTopics int) error {
		return fmt.Errorf("log missing topics: has %d, want %d", actualNumTopics, expectedNumTopics)
	}
	ErrLogMissingData   = errors.New("log missing data")
	ErrCouldNotCreateFK = func(err error) error {
		return fmt.Errorf("transformer could not create FK: %v", err)
	}
)

func VerifyLog(log types.Log, expectedNumTopics int, isDataRequired bool) error {
	actualNumTopics := len(log.Topics)
	if actualNumTopics < expectedNumTopics {
		return ErrLogMissingTopics(expectedNumTopics, actualNumTopics)
	}
	if isDataRequired && len(log.Data) < constants.DataItemLength {
		return ErrLogMissingData
	}
	return nil
}
