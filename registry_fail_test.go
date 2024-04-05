package migration_test

import (
	"errors"
)

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
