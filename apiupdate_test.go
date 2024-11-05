package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/oapi-codegen/nullable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_UpdateAdminArea_creates_new_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	body := `{"name":"test name","abbr":"test abbr","parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, "/admin-areas/2",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
	actual := res.Body.String()
	aa := entClient.AdminArea.Query().Where(adminarea.NameEQ("test name")).
		Where(adminarea.AbbrEQ("test abbr")).Where(adminarea.IDEQ(2)).
		OnlyX(context.Background())
	pid := 1
	b, err := jsoniter.Marshal(
		CreateAdminArea201JSONResponse{
			Id:        2,
			ParentId:  &pid,
			Name:      "test name",
			Abbr:      nullable.NewNullableWithValue("test abbr"),
			CreatedAt: aa.CreatedAt,
			UpdatedAt: aa.UpdatedAt,
		},
	)
	assert.Nil(t, err)
	expected := string(b)
	require.JSONEq(t, expected, actual)
}

func Test_UpdateAdminArea_reports_422(t *testing.T) {
	engine, _, res := setupGinTest(t)
	body := `{"name":"a","parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, "/admin-areas/2",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}

func Test_UpdateAdminArea_reports_422_if_parentId_equals_self(t *testing.T) {
	engine, _, res := setupGinTest(t)
	body := `{"parent_id":1}`
	req, _ := http.NewRequest(
		http.MethodPatch, "/admin-areas/1",
		io.NopCloser(strings.NewReader(body)),
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
}
