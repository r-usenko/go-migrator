package migration

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
)

var (
	sourceMethodNameValidation = regexp.MustCompile(`^M\d+$`)

	typeOfCtx   = reflect.TypeFor[context.Context]()
	typeOfError = reflect.TypeFor[error]()
	typeOfTx    = reflect.TypeFor[ITransaction]()
)

type source struct {
	receiver reflect.Value
	methods  []reflect.Method
	id       string
}

func reflectSource(instance any) (*source, error) {
	receiver := reflect.ValueOf(instance)
	var receiverPtr = receiver

	if receiver.Kind() == reflect.Struct {
		receiverPtr = reflect.New(receiver.Type())
		receiverPtr.Elem().Set(receiver)
	}

	src := &source{
		id: receiver.String(),
	}

	if receiverPtr.Kind() != reflect.Ptr {
		return src, fmt.Errorf("reflect source: %w", ErrInvalidTypeOfSource)
	}

	if receiverValue := receiverPtr.Elem(); receiverValue.Kind() != reflect.Struct {
		return src, fmt.Errorf("reflect source. pointer not struct: %w", ErrInvalidTypeOfSource)
	}

	src.receiver = receiverPtr

	receiverType := receiverPtr.Type()

	for mi := 0; mi < receiverType.NumMethod(); mi++ {
		method := receiverType.Method(mi)

		if sourceMethodNameValidation.MatchString(method.Name) {
			methodType := method.Type
			if methodType.NumOut() != 1 || !methodType.Out(0).Implements(typeOfError) {
				return src, ErrInvalidSourceMethod
			}

			src.methods = append(src.methods, method)
		}
	}

	if len(src.methods) == 0 {
		return src, ErrEmptySource
	}

	return src, nil
}
