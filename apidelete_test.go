package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
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
