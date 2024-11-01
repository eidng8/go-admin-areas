package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
	"github.com/eidng8/go-admin-areas/ent/schema"
)

func Test_DeleteAdminArea_should_delete_by_id(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	assert.Nil(t, rec.DeletedAt)
	req, _ := http.NewRequest(http.MethodDelete, "/admin-areas/1", nil)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
	_, err := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		Only(context.Background())
	assert.True(t, ent.IsNotFound(err))
}

func Test_DeleteAdminArea_should_physically_delete_if_requested(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	req, _ := http.NewRequest(
		http.MethodDelete, "/admin-areas/1?trashed=1", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
	_, err := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		Only(schema.IncludeTrashed(context.Background()))
	assert.True(t, ent.IsNotFound(err))
}

func Test_DeleteAdminArea_should_returns_404_if_not_found(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodDelete, "/admin-areas/987654321", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
