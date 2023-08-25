package grpcapi

import (
	"context"
	"local/internal/domain"
)

type ProjectsServicePort interface {
	GetByID(context.Context, int64) (domain.Project, error)
	GetAll(context.Context) ([]domain.Project, error)
	Save(context.Context, domain.Project) error
}
