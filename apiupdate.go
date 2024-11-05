package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-admin-areas/ent"
)

type UpdateAdminArea201JSONResponse AdminAreaCreate

func (response UpdateAdminArea201JSONResponse) VisitUpdateAdminAreaResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	return json.NewEncoder(w).Encode(response)
}

// UpdateAdminArea Updates a AdminArea
// (PATCH /admin-areas/{id})
func (s Server) UpdateAdminArea(
	ctx context.Context, request UpdateAdminAreaRequestObject,
) (UpdateAdminAreaResponseObject, error) {
	tx, err := s.EC.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	ac := tx.AdminArea.UpdateOneID(uint32(request.Id))
	if request.Body.Name != nil {
		ac.SetName(*request.Body.Name)
	}
	if request.Body.Abbr != nil {
		val, _ := request.Body.Abbr.Get()
		ac.SetAbbr(val)
	}
	if request.Body.ParentId != nil {
		if *request.Body.ParentId == request.Id {
			return nil, ent.NewValidationError(
				"parent_id", fmt.Errorf("ParentId cannot be equal to self"),
			)
		}
		ac.SetParentID(uint32(*request.Body.ParentId))
	}
	var aa *ent.AdminArea
	aa, err = ac.Save(ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var abbr nullable.Nullable[string]
	if nil == aa.Abbr {
		abbr = nullable.NewNullNullable[string]()
	} else {
		abbr = nullable.NewNullableWithValue(*aa.Abbr)
	}
	var pid *int
	if nil == aa.ParentID {
		pid = nil
	} else {
		val := int(*aa.ParentID)
		pid = &val
	}
	return UpdateAdminArea201JSONResponse{
		Id:        int(aa.ID),
		ParentId:  pid,
		Name:      aa.Name,
		Abbr:      abbr,
		CreatedAt: aa.CreatedAt,
		UpdatedAt: aa.UpdatedAt,
	}, nil
}
