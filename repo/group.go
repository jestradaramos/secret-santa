package repo
import (
	"errors"
	"github.com/jinzhu/gorm"
	"secret-santa/domain"
)

// InitGroupTable makes sure that the table exists on the DB we working on
func (r *Repository) InitGroupTable() {
	if !r.db.HasTable(&domain.Group{}) {
		r.db.CreateTable(&domain.Group{})
	}
}

// AddGroup used for adding groups to our db
func (r *Repository) AddGroup(group *domain.Group) (string, error){
	g := &domain.Group{}
	if result := r.db.FirstOrCreate(&g, group); result.Error != nil {
		return "", result.Error
	}
	return "Group Made", nil
}