package migration_test

import (
	"context"
	"testing"

	migration "github.com/r-usenko/go-migrator"
	"github.com/r-usenko/go-migrator/drivers/postgres"
	"github.com/r-usenko/go-migrator/drivers/processor"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	src := &srcSuccess{}

	driver1, err := processor.New(src)
	assert.NoError(t, err)

	reg := new(migration.Registry)

	err = reg.AddDriver(driver1)
	assert.NoError(t, err)

	err = reg.Run(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, src.Control, []int{1, 2, 3})
}
func TestMultiple(t *testing.T) {
	src := &srcSuccess{}

	driver1, err := processor.New(src)
	assert.NoError(t, err)

	driver2, err := postgres.New(src)
	assert.NoError(t, err)

	reg := new(migration.Registry)

	err = reg.AddDriver(driver1)
	assert.NoError(t, err)

	err = reg.AddDriver(driver2)
	assert.NoError(t, err)

	err = reg.Run(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, src.Control, []int{1, 2, 3, 1, 2, 3})
}

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
