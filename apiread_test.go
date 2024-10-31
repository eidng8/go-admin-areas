package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oapi-codegen/nullable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_ReadAdminArea_should_return_one_record(t *testing.T) {
	engine, entClient, _ := setupGinTest(t)
	rec := entClient.AdminArea.Query().Where(adminarea.ID(1)).
		OnlyX(context.Background())
	eaa := ReadAdminArea200JSONResponse{
		Id:        int(rec.ID),
		Name:      rec.Name,
		Abbr:      nullable.NewNullableWithValue(*rec.Abbr),
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
	bytes, err := jsoniter.Marshal(eaa)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas/1", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}
