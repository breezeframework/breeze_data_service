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
	Repo breeze_data.CrudRepository[T]
}

func NewGenericService[T any](repo breeze_data.CrudRepository[T]) CrudService[T] {
	return &CrudServiceImpl[T]{Repo: repo}
}

func (s *CrudServiceImpl[T]) Create(ctx context.Context, entity T) (int64, error) {
	return s.Repo.Create(ctx, entity)
}

func (s *CrudServiceImpl[T]) GetById(ctx context.Context, id int64) (*T, error) {
	return s.Repo.GetById(ctx, id)
}

func (s *CrudServiceImpl[T]) GetAll(ctx context.Context) (*[]T, error) {
	return s.Repo.GetAll(ctx)
}
func (s *CrudServiceImpl[T]) GetBy(ctx context.Context, where squirrel.Eq) (*[]T, error) {
	return s.Repo.GetBy(ctx, where)
}

func (s *CrudServiceImpl[T]) Update(ctx context.Context, id int64, entity T) error {
	return s.Repo.Update(ctx, id, entity)
}

func (s *CrudServiceImpl[T]) Delete(ctx context.Context, id int64) error {
	return s.Repo.Delete(ctx, id)
}
