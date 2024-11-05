package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_RestoreAdminArea_should_restore_by_id(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	_, err := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		Only(context.Background())
	assert.True(t, ent.IsNotFound(err))
	req, _ := http.NewRequest(http.MethodPost, "/admin-areas/1/restore", nil)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	assert.Nil(t, rec.DeletedAt)
}

func Test_RestoreAdminArea_reports_404_if_not_found(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodPost, "/admin-areas/987654321/restore", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotFound, res.Code)
}
