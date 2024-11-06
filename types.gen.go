// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
// @formatter:off
package main

import (
	"time"

	"github.com/oapi-codegen/nullable"
)

// AdminArea defines model for AdminArea.
type AdminArea struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr      nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`
	Children  *[]AdminArea              `json:"children,omitempty" yaml:"children,omitempty" xml:"children,omitempty" bson:"children,omitempty"`
	CreatedAt *time.Time                `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        int                       `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Administrative area name
	Name      string     `json:"name" yaml:"name" xml:"name" bson:"name"`
	Parent    *AdminArea `json:"parent,omitempty" yaml:"parent,omitempty" xml:"parent,omitempty" bson:"parent,omitempty"`
	ParentId  *int       `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// AdminAreaCreate defines model for AdminAreaCreate.
type AdminAreaCreate struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr      nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`
	CreatedAt *time.Time                `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        int                       `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Administrative area name
	Name      string     `json:"name" yaml:"name" xml:"name" bson:"name"`
	ParentId  *int       `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// AdminAreaList defines model for AdminAreaList.
type AdminAreaList struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr      nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`
	CreatedAt *time.Time                `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        int                       `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Administrative area name
	Name      string     `json:"name" yaml:"name" xml:"name" bson:"name"`
	ParentId  *int       `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// AdminAreaRead defines model for AdminAreaRead.
type AdminAreaRead struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr      nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`
	CreatedAt *time.Time                `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`

	// DeletedAt Date and time when the record was deleted
	DeletedAt nullable.Nullable[time.Time] `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty" xml:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Id        int                          `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Administrative area name
	Name      string     `json:"name" yaml:"name" xml:"name" bson:"name"`
	ParentId  *int       `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// AdminAreaUpdate defines model for AdminAreaUpdate.
type AdminAreaUpdate struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr      nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`
	CreatedAt *time.Time                `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        int                       `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Administrative area name
	Name      string     `json:"name" yaml:"name" xml:"name" bson:"name"`
	ParentId  *int       `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// AdminAreaParentRead defines model for AdminArea_ParentRead.
type AdminAreaParentRead struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr      nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`
	CreatedAt *time.Time                `json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at,omitempty"`
	Id        int                       `json:"id" yaml:"id" xml:"id" bson:"id"`

	// Name Administrative area name
	Name      string     `json:"name" yaml:"name" xml:"name" bson:"name"`
	ParentId  *int       `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// N400 defines model for 400.
type N400 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// N404 defines model for 404.
type N404 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// N409 defines model for 409.
type N409 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// N500 defines model for 500.
type N500 struct {
	Code   int          `json:"code" yaml:"code" xml:"code" bson:"code"`
	Errors *interface{} `json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty" bson:"errors,omitempty"`
	Status string       `json:"status" yaml:"status" xml:"status" bson:"status"`
}

// ListAdminAreaParams defines parameters for ListAdminArea.
type ListAdminAreaParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty" bson:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty" yaml:"per_page,omitempty" xml:"per_page,omitempty" bson:"per_page,omitempty"`

	// Name Name of the administrative area
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// Abbr Abbreviation of the administrative area, can be a CSV list
	Abbr *string `form:"abbr,omitempty" json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`

	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// CreateAdminAreaJSONBody defines parameters for CreateAdminArea.
type CreateAdminAreaJSONBody struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`

	// Name Administrative area name
	Name     string `json:"name" yaml:"name" xml:"name" bson:"name"`
	ParentId *int   `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
}

// DeleteAdminAreaParams defines parameters for DeleteAdminArea.
type DeleteAdminAreaParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// ReadAdminAreaParams defines parameters for ReadAdminArea.
type ReadAdminAreaParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// UpdateAdminAreaJSONBody defines parameters for UpdateAdminArea.
type UpdateAdminAreaJSONBody struct {
	// Abbr Administrative area abbreviations, CSV values
	Abbr nullable.Nullable[string] `json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`

	// Name Administrative area name
	Name     *string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`
	ParentId *int    `json:"parent_id,omitempty" yaml:"parent_id,omitempty" xml:"parent_id,omitempty" bson:"parent_id,omitempty"`
}

// ListAdminAreaChildrenParams defines parameters for ListAdminAreaChildren.
type ListAdminAreaChildrenParams struct {
	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty" bson:"page,omitempty"`

	// PerPage item count to render per page
	PerPage *int `form:"per_page,omitempty" json:"per_page,omitempty" yaml:"per_page,omitempty" xml:"per_page,omitempty" bson:"per_page,omitempty"`

	// Name Name of the administrative area
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty"`

	// Abbr Abbreviation of the administrative area, can be a CSV list
	Abbr *string `form:"abbr,omitempty" json:"abbr,omitempty" yaml:"abbr,omitempty" xml:"abbr,omitempty" bson:"abbr,omitempty"`

	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// ReadAdminAreaParentParams defines parameters for ReadAdminAreaParent.
type ReadAdminAreaParentParams struct {
	// Trashed Whether to include trashed items
	Trashed *bool `form:"trashed,omitempty" json:"trashed,omitempty" yaml:"trashed,omitempty" xml:"trashed,omitempty" bson:"trashed,omitempty"`
}

// CreateAdminAreaJSONRequestBody defines body for CreateAdminArea for application/json ContentType.
type CreateAdminAreaJSONRequestBody CreateAdminAreaJSONBody

// UpdateAdminAreaJSONRequestBody defines body for UpdateAdminArea for application/json ContentType.
type UpdateAdminAreaJSONRequestBody UpdateAdminAreaJSONBody
