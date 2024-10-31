package main

import (
	"context"

	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

// ReadAdminArea Find a AdminArea by ID
// (GET /admin-areas/{id})
func (s Server) ReadAdminArea(
	ctx context.Context, request ReadAdminAreaRequestObject,
) (ReadAdminAreaResponseObject, error) {
	area, err := s.EC.AdminArea.Query().
		Where(adminarea.ID(uint32(request.Id))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return newReadAdminArea200JSONResponseFromEnt(area), nil
}

func newReadAdminArea200JSONResponseFromEnt(eaa *ent.AdminArea) ReadAdminArea200JSONResponse {
	aa := ReadAdminArea200JSONResponse{}
	aa.Id = int(eaa.ID)
	aa.Name = eaa.Name
	if eaa.Abbr != nil {
		aa.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.Memo != nil {
		aa.Memo = nullable.NewNullableWithValue(*eaa.Memo)
	}
	if eaa.ParentID != nil {
		val := int(*eaa.ParentID)
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	return aa
}
