package postgres

import "github.com/lib/pq"

type Driver struct {
	sources []any
}

func New(sources ...any) (*Driver, error) {
	m := new(Driver)

	//TODO dummy for test dependency
	var _ pq.PGError = nil

	for _, src := range sources {
		if err := m.AddSource(src); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *Driver) AddSource(src any) error {
	m.sources = append(m.sources, src)

	return nil
}

func (m *Driver) GetSources() []any {
	return m.sources
}

func (m *Driver) Name() string {
	return "postgres"
}
