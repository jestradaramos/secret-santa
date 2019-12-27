package domain

// Repo Interface for any repos
type Repo interface {
	AddGroup(*Group) (string, error)
	GetGroupByID(id string) (*Group, error)
	AddMember(*Member) error
	GetMembersByGroupID(string) ([]Member, error)
}
