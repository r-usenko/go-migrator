package migration

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type IHaveTransaction interface {
	BeginTx(ctx context.Context) (ITransaction, error)
}

type IDriver interface {
	GetSources() []any
	WithCtx(ctx context.Context) context.Context
	//Name() string
}

type ITransaction interface {
	Commit() error
	Rollback() error
}

func NewRegistry(drivers ...IDriver) (*Registry, error) {
	m := new(Registry)

	for idx, drv := range drivers {
		if err := m.AddDriver(drv); err != nil {
			return nil, fmt.Errorf("add driver to registry[%d]: %w", idx, err)
		}
	}

	return m, nil
}

type Registry struct {
	drivers []IDriver
}

func (m *Registry) AddDriver(drv IDriver) error {
	m.drivers = append(m.drivers, drv)

	return nil
}

func (m *Registry) Run(ctx context.Context) error {
	for idxDrv, drv := range m.drivers {
		sources := drv.GetSources()
		drvCtx := drv.WithCtx(ctx)
		var transactedDriver IHaveTransaction

		for idxSrc, src := range sources {
			reflectSrc, err := reflectSource(src)
			if err != nil {
				return fmt.Errorf("get methods of source[%d.%d]: %w", idxDrv, idxSrc, err)
			}

			//sort methods
			sort.Slice(reflectSrc.methods, func(i, j int) bool {
				return reflectSrc.methods[i].Name < reflectSrc.methods[j].Name
			})

			for _, method := range reflectSrc.methods {
				if err = reflectCall(drvCtx, reflectSrc.receiver, method, transactedDriver); err != nil {
					return fmt.Errorf("call migration[%d.%d]: %w", idxDrv, idxSrc, err)
				}
			}
		}
	}

	return nil
}

func reflectCall(
	ctx context.Context,
	srcReceiver reflect.Value,
	method reflect.Method,
	transactedDriver IHaveTransaction,
) (err error) {
	srcIn := []reflect.Value{
		srcReceiver,
	}

	var tx ITransaction

	//first[0] In always must be receiver
	for i := 1; i < method.Type.NumIn(); i++ {
		//inject ctx
		if method.Type.In(i).Implements(typeOfCtx) {
			srcIn = append(srcIn, reflect.ValueOf(ctx))

			continue
		}

		//inject tx
		if transactedDriver != nil && method.Type.In(i).Implements(typeOfTx) {
			tx, err = transactedDriver.BeginTx(ctx)
			if err != nil {
				err = fmt.Errorf("driver BeginTx: %w", err)
				return
			}

			srcIn = append(srcIn, reflect.ValueOf(tx))

			continue
		}
	}

	//Commit or Rollback Tx on exit
	if tx != nil {
		defer func() {
			if err == nil {
				err = tx.Commit()
			} else {
				err = errors.Join(err, tx.Rollback())
			}
		}()
	}

	out := method.Func.Call(srcIn)

	//This is safe. We check signature of method in reflectSource
	errI := out[0].Interface()
	if errI != nil {
		err = errI.(error)
	}

	return

}
