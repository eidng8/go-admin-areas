package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"

	"github.com/eidng8/go-admin-areas/ent"
)

// DeleteAdminArea Deletes a AdminArea by ID
// (DELETE /admin-areas/{id})
func (s Server) DeleteAdminArea(
	ctx context.Context, request DeleteAdminAreaRequestObject,
) (DeleteAdminAreaResponseObject, error) {
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	tx, err := s.EC.Tx(qc)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	if err = tx.AdminArea.DeleteOneID(uint32(request.Id)).Exec(qc); err != nil {
		if ent.IsNotFound(err) {
			return DeleteAdminArea404JSONResponse{}, nil
		}
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return DeleteAdminArea204Response{}, nil
}
