package postgres

import (
	"context"
	"errors"
)

type CtxKey struct{}

var (
	ErrNotFound = errors.New("instance not found")
)

func (m *Driver) WithCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKey{}, m)
}

func FromCtx(ctx context.Context) (*Driver, error) {
	t, ok := ctx.Value(CtxKey{}).(*Driver)
	if !ok {
		return nil, ErrNotFound
	}

	return t, nil
}
