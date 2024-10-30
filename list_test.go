package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"testing"

	"github.com/eidng8/go-paginate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_ListAdminArea_should_return_first_page(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	count := entClient.AdminArea.Query().CountX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	last := int(math.Ceil(float64(count) / float64(10)))
	page := paginate.PaginatedList[AdminArea]{
		Total:        count,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     last,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=10",
		LastPageUrl: fmt.Sprintf(
			"http://127.0.0.1/admin-areas?page=%d&per_page=10", last,
		),
		NextPageUrl: "http://127.0.0.1/admin-areas?page=2&per_page=10",
		PrevPageUrl: "",
		Path:        "http://127.0.0.1/admin-areas",
		From:        1,
		To:          10,
		Data:        list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_fourth_page(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	count := entClient.AdminArea.Query().CountX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		Offset(30).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	last := int(math.Ceil(float64(count) / float64(10)))
	page := paginate.PaginatedList[AdminArea]{
		Total:        count,
		PerPage:      10,
		CurrentPage:  4,
		LastPage:     last,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=10",
		LastPageUrl: fmt.Sprintf(
			"http://127.0.0.1/admin-areas?page=%d&per_page=10", last,
		),
		NextPageUrl: "http://127.0.0.1/admin-areas?page=5&per_page=10",
		PrevPageUrl: "http://127.0.0.1/admin-areas?page=3&per_page=10",
		Path:        "http://127.0.0.1/admin-areas",
		From:        31,
		To:          40,
		Data:        list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?page=4", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_fourth_page_5_per_page(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	count := entClient.AdminArea.Query().CountX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(5).
		Offset(15).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	last := int(math.Ceil(float64(count) / float64(5)))
	page := paginate.PaginatedList[AdminArea]{
		Total:        count,
		PerPage:      5,
		CurrentPage:  4,
		LastPage:     last,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=5",
		LastPageUrl: fmt.Sprintf(
			"http://127.0.0.1/admin-areas?page=%d&per_page=5", last,
		),
		NextPageUrl: "http://127.0.0.1/admin-areas?page=5&per_page=5",
		PrevPageUrl: "http://127.0.0.1/admin-areas?page=3&per_page=5",
		Path:        "http://127.0.0.1/admin-areas",
		From:        16,
		To:          20,
		Data:        list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?page=4&per_page=5", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_report_400_for_invalid_page(t *testing.T) {
	engine, _, response := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas?page=a", nil)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func Test_ListAdminArea_should_report_400_for_invalid_perPage(t *testing.T) {
	engine, _, response := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas?per_page=a", nil)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
