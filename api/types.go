package api

import (
	"github.com/oapi-codegen/nullable"

	"eidng8.cc/microservices/admin-areas/ent"
)

func NewReadAdminArea200JSONResponseFromEnt(eaa *ent.AdminArea) ReadAdminArea200JSONResponse {
	aa := ReadAdminArea200JSONResponse{}
	aa.Id = int(eaa.ID)
	aa.Name = eaa.Name
	if eaa.Abbr != nil {
		aa.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.Memo != nil {
		aa.Memo = nullable.NewNullableWithValue(*eaa.Memo)
	}
	if eaa.ParentID != nil {
		val := int(*eaa.ParentID)
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	return aa
}

func (aa *AdminArea) FromEnt(eaa *ent.AdminArea) {
	aa.Id = int(eaa.ID)
	aa.Name = eaa.Name
	if eaa.Abbr != nil {
		aa.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.Memo != nil {
		aa.Memo = nullable.NewNullableWithValue(*eaa.Memo)
	}
	if eaa.ParentID != nil {
		val := int(*eaa.ParentID)
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
	if eaa.Edges.Parent != nil {
		aa.Parent = &AdminArea{}
		aa.Parent.FromEnt(eaa.Edges.Parent)
	}
	if eaa.Edges.Children != nil {
		children := make([]AdminArea, len(eaa.Edges.Children))
		for i, child := range eaa.Edges.Children {
			children[i] = AdminArea{}
			children[i].FromEnt(child)
		}
		aa.Children = &children
	}
}

func (aa *AdminAreaList) FromEnt(eaa *ent.AdminArea) {
	aa.Id = int(eaa.ID)
	aa.Name = eaa.Name
	if eaa.Abbr != nil {
		aa.Abbr = nullable.NewNullableWithValue(*eaa.Abbr)
	}
	if eaa.Memo != nil {
		aa.Memo = nullable.NewNullableWithValue(*eaa.Memo)
	}
	if eaa.ParentID != nil {
		val := int(*eaa.ParentID)
		aa.ParentId = &val
	}
	aa.CreatedAt = eaa.CreatedAt
	aa.UpdatedAt = eaa.UpdatedAt
}

func mapAdminArea(row *ent.AdminArea) *AdminArea {
	aa := AdminArea{}
	aa.FromEnt(row)
	return &aa
}

func mapAdminAreaIdx(row *ent.AdminArea, idx int) *AdminArea {
	return mapAdminArea(row)
}
