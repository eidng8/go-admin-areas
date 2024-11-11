package main

import (
	"context"

	"github.com/eidng8/go-ent/softdelete"

	"github.com/eidng8/go-admin-areas/ent"
)

func (s Server) RestoreAdminArea(
	ctx context.Context, request RestoreAdminAreaRequestObject,
) (RestoreAdminAreaResponseObject, error) {
	qc := softdelete.IncludeTrashed(ctx)
	id := uint32(request.Id)
	tx, err := s.EC.Tx(qc)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	err = tx.AdminArea.UpdateOneID(id).ClearDeletedAt().Exec(qc)
	if err != nil {
		if ent.IsNotFound(err) {
			return RestoreAdminArea404JSONResponse{}, nil
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return RestoreAdminArea204Response{}, nil
}
