package util

import (
	"context"

	"github.com/samber/mo"
)

func ConvertMapperWithCtx[T any](ctx context.Context, fn func(context.Context, T) mo.Result[T]) func(v T) mo.Result[T] {
	return func(v T) mo.Result[T] {
		return fn(ctx, v)
	}
}
