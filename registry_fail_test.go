package migration_test

import (
	"context"
	"errors"
	"testing"

	migration "github.com/r-usenko/go-migrator"
	"github.com/r-usenko/go-migrator/drivers/processor"
	"github.com/stretchr/testify/assert"
)

func TestFail(t *testing.T) {
	src := new(srcFail)

	driver, err := processor.New(src)
	assert.NoError(t, err)

	reg := new(migration.Registry)

	err = reg.AddDriver(driver)
	assert.NoError(t, err)

	err = reg.Run(context.Background())
	assert.Error(t, err)

	//Only the M1 migration needs to be applied
	assert.Equal(t, src.Control, []int{1})
}

type srcFail struct {
	Control []int
}

func (m *srcFail) M1() error {
	m.Control = append(m.Control, 1)
	return nil
}
func (m *srcFail) M2() error {
	return errors.New("migration error")
}
func (m *srcFail) M3() error {
	m.Control = append(m.Control, 3)
	return nil
}
