package breeze_data_servise

import (
	"context"
	"github.com/breezeframework/breeze_data/repository"
)

type CrudService[T any] interface {
	CreateEntity(ctx context.Context, entity T) (int64, error)
	GetById(ctx context.Context, id int64) (T, error)
	GetAll(ctx context.Context) (*[]T, error)
	Update(ctx context.Context, id int64, entity T) error
	Delete(ctx context.Context, id int64) error
}

type GenericService[T any] struct {
	Repo repository.CrudRepository[T]
}

func NewGenericService[T any](repo repository.CrudRepository[T]) *GenericService[T] {
	return &GenericService[T]{Repo: repo}
}

func (s *GenericService[T]) Create(ctx context.Context, entity T) (int64, error) {
	return s.Repo.Create(ctx, entity)
}

func (s *GenericService[T]) GetById(ctx context.Context, id int64) (*T, error) {
	return s.Repo.GetById(ctx, id)
}

func (s *GenericService[T]) GetAll(ctx context.Context) (*[]T, error) {
	return s.Repo.GetAll(ctx)
}

func (s *GenericService[T]) Update(ctx context.Context, id int64, entity T) error {
	return s.Repo.Update(ctx, id, entity)
}

func (s *GenericService[T]) Delete(ctx context.Context, id int64) error {
	return s.Repo.Delete(ctx, id)
}
