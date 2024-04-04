package migration

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
)

var (
	sourceMethodNameValidation = regexp.MustCompile(`^M\d+$`)

	typeOfCtx     = reflect.TypeFor[context.Context]()
	typeOfError   = reflect.TypeFor[error]()
	typeOfTx      = reflect.TypeFor[ITransaction]()
	typeOfBeginTx = reflect.TypeFor[IHaveTransaction]()
)

type source struct {
	receiver reflect.Value
	methods  []reflect.Method
	id       string
}

func reflectSource(instance any) (*source, error) {
	receiver := reflect.ValueOf(instance)
	var receiverPtr = receiver

	id := receiver.String()
	_ = id

	if receiver.Kind() == reflect.Struct {
		receiverPtr = reflect.New(receiver.Type())
		receiverPtr.Elem().Set(receiver)
	}

	if receiverPtr.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("reflect source: %w", ErrInvalidTypeOfSource)
	}

	if receiverValue := receiverPtr.Elem(); receiverValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("reflect source. pointer not struct: %w", ErrInvalidTypeOfSource)
	}

	src := &source{
		receiver: receiverPtr,
	}

	receiverType := receiverPtr.Type()

	for mi := 0; mi < receiverType.NumMethod(); mi++ {
		method := receiverType.Method(mi)

		if sourceMethodNameValidation.MatchString(method.Name) {
			methodType := method.Type
			if methodType.NumOut() != 1 || !methodType.Implements(typeOfError) {
				return nil, ErrInvalidSourceMethod
			}

			src.methods = append(src.methods, method)
		}
	}

	if len(src.methods) == 0 {
		return nil, ErrEmptySource
	}

	return src, nil
}
