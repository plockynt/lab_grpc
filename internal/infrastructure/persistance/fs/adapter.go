package fs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/exp/maps"

	"local/internal/application"
	"local/internal/domain"
)

func New(ctx context.Context, path string) (application.PersistencePort, error) {
	client := make(map[int64]domain.Project)
	adapter := adapter{
		client: client,
		path:   path,
	}
	if err := adapter.readFile(); err != nil {
		adapter.writeFile()
	}
	return &adapter, nil
}

type adapter struct {
	client map[int64]domain.Project `json:"projects"`
	path   string
}

func (obj *adapter) writeFile() error {
	content, err := json.Marshal(obj.client)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(obj.path, content, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (obj *adapter) readFile() error {
	content, err := ioutil.ReadFile(obj.path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(content, &obj.client)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (obj *adapter) Close(ctx context.Context) error {
	return nil
}

func (obj *adapter) Insert(ctx context.Context, project domain.Project) error {
	obj.client[project.ID] = project
	return obj.writeFile()
}

func (obj *adapter) Update(ctx context.Context, project domain.Project) error {
	_, found := obj.client[project.ID]
	if found {
		obj.client[project.ID] = project
		return obj.writeFile()
	}
	return errors.New("Not found")
}

func (obj *adapter) GetByID(ctx context.Context, ID int64) (domain.Project, error) {
	err := obj.readFile()
	if err != nil {
		return domain.Project{}, err
	}
	project, ok := obj.client[ID]
	if !ok {
		return domain.Project{}, errors.New("Not found")
	}
	return project, nil
}

func (obj *adapter) GetAll(ctx context.Context) ([]domain.Project, error) {
	err := obj.readFile()
	if err != nil {
		return []domain.Project{}, err
	}
	return maps.Values(obj.client), nil
}
