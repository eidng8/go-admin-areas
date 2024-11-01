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
	"github.com/eidng8/go-admin-areas/ent/schema"
)

func Test_ListAdminArea_should_return_1st_page(t *testing.T) {
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

func Test_ListAdminArea_should_return_4th_page(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		Offset(30).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        50,
		PerPage:      10,
		CurrentPage:  4,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas?page=5&per_page=10",
		PrevPageUrl:  "http://127.0.0.1/admin-areas?page=3&per_page=10",
		Path:         "http://127.0.0.1/admin-areas",
		From:         31,
		To:           40,
		Data:         list,
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

func Test_ListAdminArea_should_return_all_records(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        50,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas",
		From:         1,
		To:           50,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?per_page=12345", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_2nd_page_exclude_deleted(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.AdminArea.Delete().
		Where(adminarea.Or(adminarea.IDIn(5, 3, 21))).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.And(adminarea.IDNotIn(5, 3, 21))).
		Offset(10).Limit(10).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        47,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas?page=3&per_page=10",
		PrevPageUrl:  "http://127.0.0.1/admin-areas?page=1&per_page=10",
		Path:         "http://127.0.0.1/admin-areas",
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?page=2", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_2nd_page_include_deleted(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.AdminArea.Delete().
		Where(adminarea.IDIn(5, 3, 11)).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDLTE(20)).
		Offset(10).Limit(10).
		AllX(schema.IncludeTrashed(context.Background()))
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        50,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=10&trashed=1",
		LastPageUrl:  "http://127.0.0.1/admin-areas?page=5&per_page=10&trashed=1",
		NextPageUrl:  "http://127.0.0.1/admin-areas?page=3&per_page=10&trashed=1",
		PrevPageUrl:  "http://127.0.0.1/admin-areas?page=1&per_page=10&trashed=1",
		Path:         "http://127.0.0.1/admin-areas",
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?page=2&trashed=1", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_all_records_exclude_deleted(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.AdminArea.Delete().
		Where(adminarea.IDIn(5, 3, 21)).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDNotIn(5, 3, 21)).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        47,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas",
		From:         1,
		To:           47,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?per_page=12345", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_4th_page_5_per_page(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(5).
		Offset(15).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	last := 10
	page := paginate.PaginatedList[AdminArea]{
		Total:        50,
		PerPage:      5,
		CurrentPage:  4,
		LastPage:     last,
		FirstPageUrl: "http://127.0.0.1/admin-areas?page=1&per_page=5",
		LastPageUrl:  "http://127.0.0.1/admin-areas?page=10&per_page=5",
		NextPageUrl:  "http://127.0.0.1/admin-areas?page=5&per_page=5",
		PrevPageUrl:  "http://127.0.0.1/admin-areas?page=3&per_page=5",
		Path:         "http://127.0.0.1/admin-areas",
		From:         16,
		To:           20,
		Data:         list,
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

func Test_ListAdminArea_should_return_specified_name_prefix(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.NameHasPrefix("name 1")).Limit(10).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        11,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     2,
		FirstPageUrl: "http://127.0.0.1/admin-areas?name=name+1&page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas?name=name+1&page=2&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas?name=name+1&page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?name=name%201", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_specified_abbr_prefix(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.AbbrContains("abbr 1")).Limit(10).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        11,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     2,
		FirstPageUrl: "http://127.0.0.1/admin-areas?abbr=abbr+1&page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas?abbr=abbr+1&page=2&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas?abbr=abbr+1&page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas?abbr=abbr%201", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_apply_all_filter(t *testing.T) {
	engine, entClient, response := setupGinTest(t)
	entClient.AdminArea.DeleteOneID(1).ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		Where(adminarea.NameHasPrefix("name 1")).
		Where(adminarea.AbbrContains("abbr 1")).
		AllX(schema.IncludeTrashed(context.Background()))
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        11,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     2,
		FirstPageUrl: "http://127.0.0.1/admin-areas?abbr=abbr+1&name=name+1&page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas?abbr=abbr+1&name=name+1&page=2&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas?abbr=abbr+1&name=name+1&page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas?name=name+1&abbr=abbr%201", nil,
	)
	engine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminArea_should_return_no_record(t *testing.T) {
	engine, _, response := setupGinTest(t)
	page := paginate.PaginatedList[AdminArea]{
		Total:        0,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas?name=not+exist&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas",
		From:         0,
		To:           0,
		Data:         []*AdminArea{},
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas?name=not+exist", nil,
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
