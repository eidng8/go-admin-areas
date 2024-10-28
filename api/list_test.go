package api

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"eidng8.cc/microservices/admin-areas/ent/adminarea"
)

func Test_ListAdminArea_should_return_first_page(t *testing.T) {
	engine, entClient := setupGinTest(t)
	count := entClient.AdminArea.Query().CountX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		aa := AdminArea{}
		aa.FromEnt(row)
		list[i] = &aa
	}
	last := int(math.Ceil(float64(count) / float64(10)))
	page := PaginatedList[AdminArea]{
		Total:        count,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     last,
		FirstPageUrl: "/admin-areas?page=1&per_page=10",
		LastPageUrl:  fmt.Sprintf("/admin-areas?page=%d&per_page=10", last),
		NextPageUrl:  "/admin-areas?page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "/admin-areas",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest("GET", "/admin-areas", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_fourth_page(t *testing.T) {
	engine, entClient := setupGinTest(t)
	count := entClient.AdminArea.Query().CountX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		Offset(30).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		aa := AdminArea{}
		aa.FromEnt(row)
		list[i] = &aa
	}
	last := int(math.Ceil(float64(count) / float64(10)))
	page := PaginatedList[AdminArea]{
		Total:        count,
		PerPage:      10,
		CurrentPage:  4,
		LastPage:     last,
		FirstPageUrl: "/admin-areas?page=1&per_page=10",
		LastPageUrl:  fmt.Sprintf("/admin-areas?page=%d&per_page=10", last),
		NextPageUrl:  "/admin-areas?page=5&per_page=10",
		PrevPageUrl:  "/admin-areas?page=3&per_page=10",
		Path:         "/admin-areas",
		From:         31,
		To:           40,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest("GET", "/admin-areas?page=4", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_fourth_page_5_per_page(t *testing.T) {
	engine, entClient := setupGinTest(t)
	count := entClient.AdminArea.Query().CountX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(5).
		Offset(15).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		aa := AdminArea{}
		aa.FromEnt(row)
		list[i] = &aa
	}
	last := int(math.Ceil(float64(count) / float64(5)))
	page := PaginatedList[AdminArea]{
		Total:        count,
		PerPage:      5,
		CurrentPage:  4,
		LastPage:     last,
		FirstPageUrl: "/admin-areas?page=1&per_page=5",
		LastPageUrl:  fmt.Sprintf("/admin-areas?page=%d&per_page=5", last),
		NextPageUrl:  "/admin-areas?page=5&per_page=5",
		PrevPageUrl:  "/admin-areas?page=3&per_page=5",
		Path:         "/admin-areas",
		From:         16,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest("GET", "/admin-areas?page=4&per_page=5", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}
