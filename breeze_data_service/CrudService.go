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
	repo breeze_data.CrudRepository[T]
}

func NewCrudService[T any](repo breeze_data.CrudRepository[T]) CrudServiceImpl[T] {
	return CrudServiceImpl[T]{repo: repo}
}

func (s *CrudServiceImpl[T]) Create(ctx context.Context, entity T) (int64, error) {
	return s.repo.Create(ctx, entity)
}

func (s *CrudServiceImpl[T]) GetById(ctx context.Context, id int64) (*T, error) {
	return s.repo.GetById(ctx, id)
}

func (s *CrudServiceImpl[T]) GetAll(ctx context.Context) (*[]T, error) {
	return s.repo.GetAll(ctx)
}
func (s *CrudServiceImpl[T]) GetBy(ctx context.Context, where squirrel.Eq) (*[]T, error) {
	return s.repo.GetBy(ctx, where)
}

func (s *CrudServiceImpl[T]) Update(ctx context.Context, id int64, entity T) error {
	return s.repo.Update(ctx, id, entity)
}

func (s *CrudServiceImpl[T]) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
