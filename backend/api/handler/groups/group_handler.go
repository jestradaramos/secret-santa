package groups

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"secret-santa/api/models"
	"secret-santa/api/restapi/operations/groups"
	"secret-santa/domain"
	"secret-santa/services"
)

// TODO prevent leaking

// PostGroup ...
func PostGroup(params groups.PostGroupParams, service services.GroupService) middleware.Responder {
	fmt.Println("We are here")
	group := modelGroupToDomainGroup(params.Body)
	_, err := service.AddGroup(group)
	if err != nil {
		return &groups.PostGroupMethodNotAllowed{}
	}
	return &groups.PostGroupOK{}
}

// GetGroupID ...
func GetGroupID(params groups.GetGroupIDParams, service services.GroupService) middleware.Responder {
	g, err := service.GetGroupByID(params.ID)
	if err != nil {
		return &groups.GetGroupIDNotFound{}
	}
	return &groups.GetGroupIDOK{
		Payload: domainGroupToModel(g),
	}
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

func domainGroupToModel(d *domain.Group) *models.Group {
	modelGroup := &models.Group{
		ID:           &d.ID,
		MoneyLimit:   &d.MoneyLimit,
		Deadline:     &d.Deadline,
		ExchangeDate: &d.ExchangeDate,
	}
	return modelGroup
}
