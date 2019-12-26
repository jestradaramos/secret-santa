package services

import (
	"secret-santa/domain"
)

// MemberService struct
type MemberService struct {
	repo domain.Repo
}

// NewMemberService constructor here
func NewMemberService(repo domain.Repo) MemberService {
	return MemberService{repo}
}

// AddMember adds Members
func (s MemberService) AddMember(g *domain.Member) error {
	id, err := s.repo.AddMember(g)
	if err != nil {
		return err
	}
	return nil
}

// GetMembersByGroupID does that
func (s MemberService) GetMembersByGroupID(id string) ([]domain.Member, error) {
	members, err := s.repo.GetMembersByGroupID(id)
	if err != nil {
		return nil, err
	}
	return members, nil
}