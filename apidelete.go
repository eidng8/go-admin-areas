package main

import (
	"context"

	"github.com/eidng8/go-admin-areas/ent"
)

// Deletes a AdminArea by ID
// (DELETE /admin-areas/{id})
func (s Server) DeleteAdminArea(
	ctx context.Context, request DeleteAdminAreaRequestObject,
) (DeleteAdminAreaResponseObject, error) {
	err := s.EC.AdminArea.DeleteOneID(uint32(request.Id)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return DeleteAdminArea404JSONResponse{}, nil
		}
		return DeleteAdminArea500JSONResponse{}, err
	}
	return DeleteAdminArea204Response{}, nil
}
