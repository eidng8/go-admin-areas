package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/eidng8/go-paginate"
	"github.com/eidng8/go-softdelete"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

func Test_ListAdminAreaChildren_should_return_1st_page(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		Where(adminarea.ParentID(2)).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        48,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=2&per_page=10",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas/2/children", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_4th_page(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).Limit(10).
		Offset(30).Where(adminarea.IDGT(2)).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        48,
		PerPage:      10,
		CurrentPage:  4,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=5&per_page=10",
		PrevPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=3&per_page=10",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         31,
		To:           40,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas/2/children?page=4", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_all_records(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDGT(2)).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        48,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         1,
		To:           48,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas/2/children?per_page=12345", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_2nd_page_exclude_deleted(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.AdminArea.Delete().
		Where(adminarea.Or(adminarea.IDIn(5, 3, 21))).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.And(adminarea.IDNotIn(5, 3, 21))).
		Where(adminarea.IDGT(2)).Offset(10).Limit(10).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        45,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=10",
		LastPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=5&per_page=10",
		NextPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=3&per_page=10",
		PrevPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=1&per_page=10",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas/2/children?page=2", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_2nd_page_include_deleted(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.AdminArea.Delete().Where(adminarea.IDIn(5, 3, 11)).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDLTE(22)).Where(adminarea.ParentIDEQ(2)).
		Offset(10).Limit(10).AllX(softdelete.IncludeTrashed(context.Background()))
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        48,
		PerPage:      10,
		CurrentPage:  2,
		LastPage:     5,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=10&trashed=1",
		LastPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=5&per_page=10&trashed=1",
		NextPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=3&per_page=10&trashed=1",
		PrevPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=1&per_page=10&trashed=1",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         11,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas/2/children?page=2&trashed=1", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_all_records_exclude_deleted(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.AdminArea.Delete().Where(adminarea.IDIn(5, 3, 21)).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDNotIn(5, 3, 21)).Where(adminarea.ParentIDEQ(2)).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        45,
		PerPage:      12345,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=12345",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         1,
		To:           45,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas/2/children?per_page=12345", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_4th_page_5_per_page(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDGT(2)).Limit(5).Offset(15).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        48,
		PerPage:      5,
		CurrentPage:  4,
		LastPage:     10,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?page=1&per_page=5",
		LastPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=10&per_page=5",
		NextPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=5&per_page=5",
		PrevPageUrl:  "http://127.0.0.1/admin-areas/2/children?page=3&per_page=5",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         16,
		To:           20,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas/2/children?page=4&per_page=5", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_specified_name_prefix(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.ParentIDEQ(2)).Where(adminarea.NameHasPrefix("name 1")).
		Limit(10).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        10,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?name=name+1&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas/2/children?name=name%201",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_specified_abbr_prefix(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.AbbrContains("abbr 1")).Where(adminarea.ParentIDEQ(2)).
		Limit(10).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        10,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?abbr=abbr+1&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         1,
		To:           10,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas/2/children?abbr=abbr%201",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_apply_all_filter(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	entClient.AdminArea.DeleteOneID(11).ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.NameHasPrefix("name 1")).
		Where(adminarea.AbbrContains("abbr 1")).
		Where(adminarea.ParentIDEQ(2)).
		Limit(10).AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        9,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?abbr=abbr+1&name=name+1&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         1,
		To:           9,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas/2/children?name=name+1&abbr=abbr%201",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_return_no_record(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().Where(adminarea.IDGT(2)).SetParentID(2).
		SaveX(context.Background())
	page := paginate.PaginatedList[AdminArea]{
		Total:        0,
		PerPage:      10,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/2/children?name=not+exist&page=1&per_page=10",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/2/children",
		From:         0,
		To:           0,
		Data:         []*AdminArea{},
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1/admin-areas/2/children?name=not+exist", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}

func Test_ListAdminAreaChildren_should_report_400_for_invalid_page(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(
		http.MethodGet, "/admin-areas/2/children?page=a", nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func Test_ListAdminAreaChildren_should_report_400_for_invalid_perPage(t *testing.T) {
	engine, _, res := setupGinTest(t)
	req, _ := http.NewRequest(http.MethodGet, "/admin-areas?per_page=a", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func Test_ListAdminAreaChildren_should_return_all_descendants(t *testing.T) {
	engine, entClient, res := setupGinTest(t)
	entClient.AdminArea.Update().SetParentID(1).
		Where(adminarea.IDIn(2, 3)).ExecX(context.Background())
	entClient.AdminArea.Update().SetParentID(2).
		Where(adminarea.IDIn(4, 5, 6)).ExecX(context.Background())
	entClient.AdminArea.Update().SetParentID(3).
		Where(adminarea.IDIn(7, 8)).ExecX(context.Background())
	entClient.AdminArea.Update().SetParentID(4).
		Where(adminarea.IDIn(9, 10, 11, 12)).
		ExecX(context.Background())
	rows := entClient.AdminArea.Query().Order(adminarea.ByID()).
		Where(adminarea.IDGT(1)).Where(adminarea.IDLTE(12)).
		AllX(context.Background())
	list := make([]*AdminArea, len(rows))
	for i, row := range rows {
		list[i] = newAdminAreaFromEnt(row)
	}
	page := paginate.PaginatedList[AdminArea]{
		Total:        11,
		PerPage:      11,
		CurrentPage:  1,
		LastPage:     1,
		FirstPageUrl: "http://127.0.0.1/admin-areas/1/children?recurse=1",
		LastPageUrl:  "",
		NextPageUrl:  "",
		PrevPageUrl:  "",
		Path:         "http://127.0.0.1/admin-areas/1/children",
		From:         1,
		To:           11,
		Data:         list,
	}
	bytes, err := jsoniter.Marshal(page)
	assert.Nil(t, err)
	expected := string(bytes)
	req, _ := http.NewRequest(
		http.MethodGet, "http://127.0.0.1/admin-areas/1/children?recurse=1",
		nil,
	)
	engine.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	actual := res.Body.String()
	require.JSONEq(t, expected, actual)
}
