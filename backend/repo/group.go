package repo

import (
	// "errors"
	// "github.com/jinzhu/gorm"
	"secret-santa/domain"
)

// InitGroupTable ...
func (r *Repository) InitGroupTable() {
	if !r.db.HasTable(&domain.Group{}) {
		r.db.CreateTable(&domain.Group{})
	}
}

// AddGroup ...
func (r *Repository) AddGroup(group *domain.Group) (string, error) {
	g := &domain.Group{}
	if result := r.db.FirstOrCreate(&g, group); result.Error != nil {
		return "", result.Error
	}
	return "Group Made", nil
}

// GetGroupByID ...
func (r *Repository) GetGroupByID(id string) (*domain.Group, error) {
	g := &domain.Group{}
	if result := r.db.Where("id = ?", id).Find(g); result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}
