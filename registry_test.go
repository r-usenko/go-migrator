package migration_test

import (
	"context"
	"errors"
	"testing"

	migration "github.com/r-usenko/go-migrator"
	"github.com/r-usenko/go-migrator/drivers/processor"
)

type script struct {
}

func (m *script) M1() error {
	return nil
}
func (m *script) M200() error {
	return errors.New("migration error")
}

func TestSetup(t *testing.T) {
	driver, err := processor.New(new(script))
	if err != nil {
		panic(err)
	}

	reg := new(migration.Registry)
	if err = reg.AddDriver(driver); err != nil {
		panic(err)
	}
	if err = reg.Run(context.Background()); err != nil {
		panic(err)
	}
}
