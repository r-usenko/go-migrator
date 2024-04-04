package processor

type Driver struct {
	sources []any
}

func New(sources ...any) (*Driver, error) {
	m := new(Driver)

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
