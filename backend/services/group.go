package services

import (
	"secret-santa/domain"
)

// GroupService struct
type GroupService struct {
	repo domain.Repo
}

// NewGroupService constructor here
func NewGroupService(repo domain.Repo) GroupService {
	return GroupService{repo}
}

// AddGroup adds groups
func (s GroupService) AddGroup(g *domain.Group) (string, error) {
	id, err := s.repo.AddGroup(g)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetGroupByID ...
func (s GroupService) GetGroupByID(id string) (*domain.Group, error) {
	g, err := s.repo.GetGroupByID(id)
	if err != nil {
		return nil, err
	}
	return g, nil
}
