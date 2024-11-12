package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/oapi-codegen/nullable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_ReadAdminArea_should_return_one_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	eaa := ReadAdminArea200JSONResponse{
		Id:        rec.ID,
		Name:      rec.Name,
		Abbr:      nullable.NewNullableWithValue(*rec.Abbr),
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/1", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ReadAdminArea_does_not_returns_deleted_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/1", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadAdminArea_returns_deleted_record_if_requested(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(softdelete.IncludeTrashed(context.Background()))
	eaa := ReadAdminArea200JSONResponse{
		Id:        rec.ID,
		Name:      rec.Name,
		Abbr:      nullable.NewNullableWithValue(*rec.Abbr),
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: nullable.NewNullableWithValue(*rec.DeletedAt),
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/1?trashed=1", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ReadAdminArea_should_return_404_if_not_found(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/987654321", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
