package mem

import (
	"context"
	"errors"

	"golang.org/x/exp/maps"

	"local/internal/application"
	"local/internal/domain"
)

func New(ctx context.Context) (application.PersistencePort, error) {
	client := make(map[int64]domain.Project)

	return &adapter{
		client: client,
	}, nil
}

type adapter struct {
	client map[int64]domain.Project
}

func (obj *adapter) Close(context.Context) error {
	return nil
}

func (obj *adapter) Insert(ctx context.Context, project domain.Project) error {
	obj.client[project.ID] = project
	return nil
}

func (obj *adapter) Update(ctx context.Context, project domain.Project) error {
	_, found := obj.client[project.ID]
	if found {
		obj.client[project.ID] = project
		return nil
	}
	return errors.New("Not found")
}

func (obj *adapter) GetByID(ctx context.Context, ID int64) (domain.Project, error) {
	project, ok := obj.client[ID]
	if !ok {
		return domain.Project{}, errors.New("Not found")
	}
	return project, nil
}

func (obj *adapter) GetAll(ctx context.Context) ([]domain.Project, error) {
	return maps.Values(obj.client), nil
}
