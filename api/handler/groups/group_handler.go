package groups

import (
	"github.com/go-openapi/runtime/middleware"
	"secret-santa/api/models"
	"secret-santa/api/restapi/operations/groups"
	"secret-santa/domain"
	"secret-santa/services"
)

// TODO Fix correct responses

// PostGroup ...
func PostGroup(params groups.PostGroupParams, service services.GroupService) middleware.Responder {
	group := modelGroupToDomainGroup(params.Body)
	_, err := service.AddGroup(group)
	if err != nil {
		return &groups.GetGroupIDNotFound{}
	}
	return &groups.PostGroupOK{}
}

// GetGroupID ...
func GetGroupID(params groups.GetGroupIDParams, service services.GroupService) middleware.Responder {
	return &groups.GetGroupIDOK{}
}

func modelGroupToDomainGroup(model *models.Group) *domain.Group {
	domainGroup := &domain.Group{
		ID:           *model.ID,
		MoneyLimit:   *model.MoneyLimit,
		Deadline:     *model.Deadline,
		ExchangeDate: *model.ExchangeDate,
	}
	return domainGroup
}
