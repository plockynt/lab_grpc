package application

import (
	"context"

	"local/internal/domain"
)

type ProjectsServicePort interface {
	Save(context.Context, domain.Project) error
	GetByID(context.Context, int64) (domain.Project, error)
	GetAll(context.Context) ([]domain.Project, error)
}

func NewProjectsSvc(persistencePort PersistencePort) ProjectsServicePort {
	return &projectsSvc{
		persistence: persistencePort,
	}
}

type projectsSvc struct {
	persistence PersistencePort
}

func (obj *projectsSvc) Save(ctx context.Context, project domain.Project) error {
	err := obj.persistence.Update(ctx, project)
	if err != nil {
		return obj.persistence.Insert(ctx, project)
	}
	return err
}

func (obj *projectsSvc) GetByID(ctx context.Context, ID int64) (domain.Project, error) {
	return obj.persistence.GetByID(ctx, ID)
}

func (obj *projectsSvc) GetAll(ctx context.Context) ([]domain.Project, error) {
	return obj.persistence.GetAll(ctx)
}
