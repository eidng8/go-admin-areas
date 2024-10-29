package main

import (
	"context"

	"eidng8.cc/microservices/admin-areas/ent/adminarea"
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
	return NewReadAdminArea200JSONResponseFromEnt(area), nil
}
