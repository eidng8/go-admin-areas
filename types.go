package main

import (
	"github.com/eidng8/go-ent/paginate"
	"github.com/oapi-codegen/nullable"

	"github.com/eidng8/go-admin-areas/ent"
)

func newAdminAreaFromEnt(eaa *ent.AdminArea) *AdminArea {
	aa := AdminArea{}
	aa.Id = int(eaa.ID)
	aa.Name = eaa.Name
	if eaa.Abbr != nil {
		aa.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.ParentID != nil {
		val := int(*eaa.ParentID)
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	if eaa.Edges.Parent != nil {
		aa.Parent = newAdminAreaFromEnt(eaa.Edges.Parent)
	}
	if eaa.Edges.Children != nil {
		children := make([]AdminArea, len(eaa.Edges.Children))
		for i, child := range eaa.Edges.Children {
			children[i] = *newAdminAreaFromEnt(child)
		}
		aa.Children = &children
	}
	return &aa
}

func newAdminAreaListFromEnt(eaa *ent.AdminArea) AdminAreaList {
	aa := AdminAreaList{}
	aa.Id = int(eaa.ID)
	aa.Name = eaa.Name
	if eaa.Abbr != nil {
		aa.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.ParentID != nil {
		val := int(*eaa.ParentID)
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	return aa
}

func mapAdminAreaListFromEnt(array []*ent.AdminArea) []AdminAreaList {
	data := make([]AdminAreaList, len(array))
	for i, row := range array {
		data[i] = newAdminAreaListFromEnt(row)
	}
	return data
}

func mapPage[T ListAdminArea200JSONResponse | ListAdminAreaChildren200JSONResponse](
	page *paginate.PaginatedList[ent.AdminArea],
) T {
	return T{
		CurrentPage:  page.CurrentPage,
		FirstPageUrl: page.FirstPageUrl,
		From:         page.From,
		LastPage:     page.LastPage,
		LastPageUrl:  page.LastPageUrl,
		NextPageUrl:  page.NextPageUrl,
		Path:         page.Path,
		PerPage:      page.PerPage,
		PrevPageUrl:  page.PrevPageUrl,
		To:           page.To,
		Total:        page.Total,
		Data:         mapAdminAreaListFromEnt(page.Data),
	}
}
