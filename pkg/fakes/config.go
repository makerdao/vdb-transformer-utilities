package fakes

import (
	"io/ioutil"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetFakeConfig() {
	logrus.SetOutput(ioutil.Discard)
	viper.SetConfigName("config")
	viper.AddConfigPath("$GOPATH/src/github.com/makerdao/vdb-transformer-utilities/pkg/fakes/")
}

func FakeHash() common.Hash {
	return common.HexToHash(randomString(64))
}

func randomString(length int) string {
	charset := "abcdef1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
