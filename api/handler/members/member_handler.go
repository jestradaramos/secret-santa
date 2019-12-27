package members

import (
	"github.com/go-openapi/runtime/middleware"
	"secret-santa/api/restapi/operations/members"
	"secret-santa/domain"
	"secret-santa/services"
)

// TODO fix responses

// PostMember ...
func PostMember(params members.PostMemberParams, service services.MemberService) middleware.Responder {
	member := &domain.Member{
		Name:    params.Member.Name,
		Email:   params.Member.Email,
		Spouse:  params.Member.Spouse,
		GroupID: params.Member.Group,
	}
	err := service.AddMember(member)
	if err != nil {
		return &members.PostMemberBadRequest{}
	}
	return &members.PostMemberOK{}
}
