package grpcapi

import (
	"context"
	"errors"
	"fmt"
	"log"

	"local/internal/domain"
)

type Server struct {
	projectsSvc ProjectsServicePort
}

func NewServer(projectsSvc ProjectsServicePort) Server {
	return Server{projectsSvc: projectsSvc}
}

func toDTO(project domain.Project) Project {
	return Project{
		Id:          project.ID,
		Label:       project.Label,
		Description: project.Description,
		Status:      project.Status,
		HoursSpent:  project.HoursSpent,
		Subprojects: project.Subprojects,
	}
}

func fromDTO(project Project) domain.Project {
	return domain.Project{
		ID:          project.Id,
		Label:       project.Label,
		Description: project.Description,
		Status:      project.Status,
		HoursSpent:  project.HoursSpent,
		Subprojects: project.Subprojects,
	}
}

func (s *Server) GetByID(ctx context.Context, in *ID) (*Project, error) {
	log.Printf("GetByID request, ID : %s\n", in.Id)
	project, err := s.projectsSvc.GetByID(ctx, in.Id)
	if err != nil {
		return &Project{}, errors.New(fmt.Sprintf("error while getting by id: %s", err))
	}
	projectRet := toDTO(project)
	return &projectRet, nil
}

func (s *Server) GetAll(in *GetAllParams, srv ProjectTracker_GetAllServer) error {
	log.Printf("GetAll request\n")
	projects, err := s.projectsSvc.GetAll(context.Background())
	if err != nil {
		return errors.New(fmt.Sprintf("error while getting all: %s", err))
	}
	for _, project := range projects {
		projectRet := toDTO(project)
		if err := srv.Send(&projectRet); err != nil {
			log.Println("error generating response: %s", err)
			return err
		}
	}

	return nil
}

func (s *Server) Save(ctx context.Context, in *Project) (*SaveReturn, error) {
	log.Printf("Save request, ProjectID : %s\n", in.Id)
	project := fromDTO(*in)
	err := s.projectsSvc.Save(ctx, project)
	if err != nil {
		return &SaveReturn{}, errors.New(fmt.Sprintf("error while saving: %s", err))
	}
	return &SaveReturn{}, nil
}

func (s *Server) mustEmbedUnimplementedProjectTrackerServer() {

}
