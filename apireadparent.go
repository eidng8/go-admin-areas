package main

import (
	"context"

	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

// ReadAdminAreaParent Find a AdminArea by ID
// (GET /admin-areas/{id}/parent)
func (s Server) ReadAdminAreaParent(
	ctx context.Context, request ReadAdminAreaParentRequestObject,
) (ReadAdminAreaParentResponseObject, error) {
	area, err := s.EC.AdminArea.Query().Where(adminarea.ID(uint32(request.Id))).
		WithParent().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ReadAdminAreaParent404JSONResponse{}, nil
		}
		return nil, err
	}
	if nil == area || nil == area.Edges.Parent {
		return ReadAdminAreaParent404JSONResponse{}, nil
	}
	return newReadAdminAreaParent200JSONResponseFromEnt(area.Edges.Parent), nil
}

func newReadAdminAreaParent200JSONResponseFromEnt(
	eaa *ent.AdminArea,
) ReadAdminAreaParent200JSONResponse {
	aar := ReadAdminAreaParent200JSONResponse{}
	aar.Id = eaa.ID
	aar.Name = eaa.Name
	if eaa.Abbr != nil {
		aar.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.ParentID != nil {
		val := *eaa.ParentID
		aar.ParentId = &val
	}
	aar.CreatedAt = eaa.CreatedAt
	aar.UpdatedAt = eaa.UpdatedAt
	return aar
}
