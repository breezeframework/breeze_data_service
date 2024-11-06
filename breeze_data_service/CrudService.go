package breeze_data_service

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/breezeframework/breeze_data/breeze_data"
)

type CrudService[T any] interface {
	GetById(ctx context.Context, id int64) (*T, error)
	GetBy(ctx context.Context, where squirrel.Eq) (*[]T, error)
	GetAll(ctx context.Context) (*[]T, error)
	Create(ctx context.Context, entity T) (int64, error)
	Update(ctx context.Context, id int64, entity T) error
	Delete(ctx context.Context, id int64) error
}

type CrudServiceImpl[T any] struct {
	breeze_data.CrudRepository[T]
}

func NewCrudService[T any](repo breeze_data.CrudRepository[T]) CrudServiceImpl[T] {
	return CrudServiceImpl[T]{repo}
}
