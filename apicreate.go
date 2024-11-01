package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-admin-areas/ent"
)

type CreateAdminArea201JSONResponse AdminAreaCreate

func (response CreateAdminArea201JSONResponse) VisitCreateAdminAreaResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	return json.NewEncoder(w).Encode(response)
}

// CreateAdminArea Create a new AdminArea
// (POST /admin-areas)
func (s Server) CreateAdminArea(
	ctx context.Context, request CreateAdminAreaRequestObject,
) (CreateAdminAreaResponseObject, error) {
	tx, err := s.EC.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if nil != err {
			_ = tx.Rollback()
		}
	}()
	ac := tx.AdminArea.Create()
	ac.SetName(request.Body.Name)
	if request.Body.Abbr != nil {
		val, _ := request.Body.Abbr.Get()
		ac.SetAbbr(val)
	}
	if request.Body.ParentId != nil {
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
	return CreateAdminArea201JSONResponse{
		Id:        int(aa.ID),
		ParentId:  pid,
		Name:      aa.Name,
		Abbr:      abbr,
		CreatedAt: aa.CreatedAt,
		UpdatedAt: aa.UpdatedAt,
	}, nil
}
