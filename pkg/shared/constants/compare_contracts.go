package constants

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type MismatchedConstructorsError struct {
	constructorOne, constructorTwo string
}

func NewMismatchedConstructorsError(a, b abi.ABI) error {
	return &MismatchedConstructorsError{
		constructorOne: a.Constructor.String(),
		constructorTwo: b.Constructor.String(),
	}
}

func (err *MismatchedConstructorsError) Error() string {
	return fmt.Sprintf("constructors don't match constructorOne: %s, constructorTwo: %s", err.constructorOne, err.constructorTwo)
}

type MismatchedMethodsError struct {
	method, methodOne, methodTwo string
}

func NewMismatchedMethodsError(a, b abi.ABI, method string) error {
	return &MismatchedMethodsError{
		method:    method,
		methodOne: a.Methods[method].String(),
		methodTwo: b.Methods[method].String(),
	}
}

func (err *MismatchedMethodsError) Error() string {
	return fmt.Sprintf("outer methods don't match for method %s, method one: %s, method two: %s", err.method, err.methodOne, err.methodTwo)
}

type MismatchedEventsError struct {
	event, eventOne, eventTwo string
}

func NewMismatchedEventsError(a, b abi.ABI, event string) error {
	return &MismatchedEventsError{
		eventOne: a.Events[event].String(),
		eventTwo: a.Events[event].String(),
		event:    event,
	}
}

func (err *MismatchedEventsError) Error() string {
	return fmt.Sprintf("outer events don't match for event %s, event one: %s, event two: %s", err.event, err.eventOne, err.eventTwo)
}

func CompareContractABI(a, b abi.ABI) error {
	if a.Constructor.String() != b.Constructor.String() {
		return NewMismatchedConstructorsError(a, b)
	}
	for ak, av := range a.Methods {
		if av.String() != b.Methods[ak].String() {
			return NewMismatchedMethodsError(a, b, ak)
		}
	}
	for ak, av := range a.Events {
		if av.String() != b.Events[ak].String() {
			return NewMismatchedEventsError(a, b, ak)
		}
	}
	return nil
}
