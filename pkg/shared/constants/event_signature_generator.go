package constants

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/makerdao/vulcanizedb/pkg/eth"
)

func GetEventTopicZero(solidityEventSignature string) string {
	eventSignature := []byte(solidityEventSignature)
	hash := crypto.Keccak256Hash(eventSignature)
	return hash.Hex()
}

func GetSolidityFunctionSignature(abi, name string) string {
	parsedAbi, _ := eth.ParseAbi(abi)

	if method, ok := parsedAbi.Methods[name]; ok {
		return method.Sig
	} else if event, ok := parsedAbi.Events[name]; ok {
		return GetEventSignature(event)
	}
	panic("Error: could not get Solidity method signature for: " + name)
}

func GetEventSignature(event abi.Event) string {
	types := make([]string, len(event.Inputs))
	for i, input := range event.Inputs {
		types[i] = input.Type.String()
		i++
	}

	return fmt.Sprintf("%v(%v)", event.Name, strings.Join(types, ","))
}

type ContractMethod struct {
	Name   string
	Inputs []MethodInput
}

type MethodInput struct {
	Type string
}

func GetOverloadedFunctionSignature(rawAbi, name string, paramTypes []string) string {
	result, err := FindSignatureInAbi(rawAbi, name, paramTypes)
	if err != nil {
		panic(err)
	}
	return result
}

func FindSignatureInAbi(rawAbi, name string, paramTypes []string) (string, error) {
	contractMethods := make([]ContractMethod, 0)
	err := json.Unmarshal([]byte(rawAbi), &contractMethods)
	if err != nil {
		return "", errors.New("unable to parse ABI")
	}
	signature := fmt.Sprintf("%v(%v)", name, strings.Join(paramTypes, ","))
	if containsMatchingMethod(contractMethods, name, paramTypes) == false {
		return "", errors.New("method " + signature + " does not exist in ABI")
	}
	return signature, nil
}

func containsMatchingMethod(methods []ContractMethod, name string, paramTypes []string) bool {
	for _, method := range methods {
		if method.Name == name && hasMatchingParams(method, paramTypes) {
			return true
		}
	}
	return false
}

func hasMatchingParams(method ContractMethod, expectedParamTypes []string) bool {
	params := method.Inputs
	actualParamTypes := make([]string, len(params))
	for i, param := range params {
		actualParamTypes[i] = param.Type
	}
	return areEqual(expectedParamTypes, actualParamTypes)
}

func areEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
