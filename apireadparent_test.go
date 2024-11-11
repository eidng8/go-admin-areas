package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	sd "github.com/eidng8/go-ent/softdelete"
	"github.com/oapi-codegen/nullable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_ReadAdminAreaParent_should_return_one_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.UpdateOneID(uint32(2)).SetParentID(1).
		SaveX(context.Background())
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	eaa := ReadAdminAreaParent200JSONResponse{
		Id:        1,
		Name:      "name 0",
		Abbr:      nullable.NewNullableWithValue("abbr 0"),
		CreatedAt: rec.CreatedAt,
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ReadAdminAreaParent_does_not_returns_deleted_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.DeleteOneID(2).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadAdminAreaParent_does_not_return_deleted_parent(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/2/parent", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_ReadAdminAreaParent_returns_deleted_record_if_requested(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.UpdateOneID(uint32(2)).SetParentID(1).
		SetDeletedAt(time.Now()).
		SaveX(sd.IncludeTrashed(context.Background()))
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	eaa := ReadAdminAreaParent200JSONResponse{
		Id:        1,
		Name:      "name 0",
		Abbr:      nullable.NewNullableWithValue("abbr 0"),
		CreatedAt: rec.CreatedAt,
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "/admin-areas/2/parent?trashed=1", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ReadAdminAreaParent_returns_deleted_parent_if_requested(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.UpdateOneID(uint32(2)).SetParentID(1).
		SaveX(context.Background())
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(sd.IncludeTrashed(context.Background()))
	eaa := ReadAdminAreaParent200JSONResponse{
		Id:        1,
		Name:      "name 0",
		Abbr:      nullable.NewNullableWithValue("abbr 0"),
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "/admin-areas/2/parent?trashed=1", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ReadAdminAreaParent_should_return_404_if_not_found(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodGet, "/admin-areas/987654321/parent", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
