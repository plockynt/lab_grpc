package application

import (
	"context"

	"local/internal/domain"
)

type PersistencePort interface {
	Close(context.Context) error
	Insert(context.Context, domain.Project) error
	Update(context.Context, domain.Project) error
	GetByID(context.Context, int64) (domain.Project, error)
	GetAll(context.Context) ([]domain.Project, error)
}
