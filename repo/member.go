package repo
import (
	"errors"
	"github.com/jinzhu/gorm"
	"secret-santa/domain"
)

// InitMemberTable makes sure that the table exists on the DB we working on
func (r *Repository) InitMemberTable() {
	if !r.db.HasTable(&domain.Member{}) {
		r.db.CreateTable(&domain.Member{})
	}
}

// AddMember used for adding members to our db
func (r *Repository) AddMember(member *domain.Member) error{
	m := &domain.Member{}
	if result := r.db.FirstOrCreate(&m, member); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetMembersByGroupID is just that
func (r *Repository) GetMembersByGroupID(id string) ([]domain.Member , error) {
	return nil, nil
}