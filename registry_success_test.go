package migration_test

type srcSuccess struct {
	Control []int
}

func (m *srcSuccess) M1() error {
	m.Control = append(m.Control, 1)
	return nil
}
func (m *srcSuccess) M3() error {
	m.Control = append(m.Control, 3)
	return nil
}
func (m *srcSuccess) M2() error {
	m.Control = append(m.Control, 2)
	return nil
}
