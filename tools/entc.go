//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/eidng8/go-ent/simpletree"
	"github.com/eidng8/go-ent/softdelete"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ogen-go/ogen"
)

func main() {
	oas, err := newOasExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ext := entc.Extensions(oas, &simpletree.Extension{})
	err = entc.Generate("./ent/schema", genConfig(), ext)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func newOasExtension() (*entoas.Extension, error) {
	return entoas.NewExtension(
		entoas.Mutations(
			func(g *gen.Graph, s *ogen.Spec) error {
				genSpec(s)
				constraintRequestBody(s.Paths)
				softdelete.AttachTo(s, "/admin-areas", "AdminAreaRead")
				ep := s.Paths["/admin-areas"]
				fixPerPageParamName(ep.Get.Parameters)
				ep.Get.AddParameters(nameParam(), abbrParam())
				simpletree.RemoveEdges(ep.Post)
				setPaginateResponse(ep.Get)
				ep = s.Paths["/admin-areas/{id}"]
				simpletree.RemoveEdges(ep.Patch)
				ep = s.Paths["/admin-areas/{id}/parent"]
				ep = s.Paths["/admin-areas/{id}/children"]
				fixPerPageParamName(ep.Get.Parameters)
				setPaginateResponse(ep.Get)
				ep.Get.AddParameters(nameParam(), abbrParam(), recurseParam())
				return nil
			},
		),
	)
}

func genConfig() *gen.Config {
	return &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureExecQuery,
			gen.FeatureVersionedMigration,
		},
		// Templates: []*gen.Template{
		// 	gen.MustParse(gen.NewTemplate("query_cte").ParseFiles("query_cte.tmpl")),
		// },
	}
}

func genSpec(s *ogen.Spec) {
	s.Info.SetTitle("Administrative areas listing API")
	s.Info.SetDescription("This is an API listing administrative areas")
	s.Info.SetVersion("0.0.1")
}

func removeFields(props []ogen.Property, fields ...string) []ogen.Property {
	for _, field := range fields {
		for i, prop := range props {
			if prop.Name == field {
				props = append(props[:i], props[i+1:]...)
				break
			}
		}
	}
	return props

}

func constraintRequestBody(paths ogen.Paths) {
	b := false
	for _, path := range paths {
		for _, op := range []*ogen.Operation{path.Put, path.Post, path.Patch} {
			if nil == op || nil == op.RequestBody || nil == op.RequestBody.Content {
				continue
			}
			for _, param := range op.RequestBody.Content {
				if nil == param.Schema {
					continue
				}
				param.Schema.AdditionalProperties = &ogen.AdditionalProperties{Bool: &b}
			}
		}
	}
}

func fixPerPageParamName(params []*ogen.Parameter) {
	for _, param := range params {
		if "itemsPerPage" == param.Name {
			param.Name = "per_page"
		}
	}
}

func removeEdges(op *ogen.Operation) {
	schema := op.RequestBody.Content["application/json"].Schema
	schema.Properties = removeFields(schema.Properties, "parent", "children")
}

func nameParam() *ogen.Parameter {
	u2 := uint64(2)
	u255 := uint64(255)
	return &ogen.Parameter{
		Name:        "name",
		In:          "query",
		Description: "Name of the administrative area",
		Required:    false,
		Schema: &ogen.Schema{
			Type:      "string",
			MinLength: &u2,
			MaxLength: &u255,
		},
	}
}

func abbrParam() *ogen.Parameter {
	u2 := uint64(2)
	u255 := uint64(255)
	return &ogen.Parameter{
		Name:        "abbr",
		In:          "query",
		Description: "Abbreviation of the administrative area, can be a CSV list",
		Required:    false,
		Schema: &ogen.Schema{
			Type:      "string",
			MinLength: &u2,
			MaxLength: &u255,
		},
	}
}

func trashedParam() *ogen.Parameter {
	return &ogen.Parameter{
		Name:        "trashed",
		In:          "query",
		Description: "Whether to include trashed items",
		Required:    false,
		Schema:      &ogen.Schema{Type: "boolean"},
	}
}

func recurseParam() *ogen.Parameter {
	return &ogen.Parameter{
		Name:        "recurse",
		In:          "query",
		Description: "Whether to return all descendants (recurse to last leaf)",
		Required:    false,
		Schema:      &ogen.Schema{Type: "boolean"},
	}
}

func setPaginateResponse(op *ogen.Operation) {
	op.Responses["200"] = &ogen.Response{
		Description: "Paginated list of administrative areas",
		Content: map[string]ogen.Media{
			"application/json": {
				Schema: &ogen.Schema{
					Type: "object",
					Properties: []ogen.Property{
						{
							Name: "current_page",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Page number (1-based)",
								Minimum:     ogen.Num("1"),
							},
						},
						{
							Name: "total",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Total number of items",
								Minimum:     ogen.Num("0"),
							},
						},
						{
							Name: "per_page",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Number of items per page",
								Minimum:     ogen.Num("1"),
							},
						},
						{
							Name: "last_page",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Last page number",
								Minimum:     ogen.Num("1"),
							},
						},
						{
							Name: "from",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Index (1-based) of the first item in the current page",
								Minimum:     ogen.Num("0"),
							},
						},
						{
							Name: "to",
							Schema: &ogen.Schema{
								Type:        "integer",
								Description: "Index (1-based) of the last item in the current page",
								Minimum:     ogen.Num("0"),
							},
						},
						{
							Name: "first_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the first page",
							},
						},
						{
							Name: "last_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the last page",
							},
						},
						{
							Name: "next_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the next page",
							},
						},
						{
							Name: "prev_page_url",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "URL to the previous page",
							},
						},
						{
							Name: "path",
							Schema: &ogen.Schema{
								Type:        "string",
								Description: "Base path of the request",
							},
						},
						{
							Name: "data",
							Schema: &ogen.Schema{
								Type:        "array",
								Description: "List of administrative areas",
								Items: &ogen.Items{
									Item: &ogen.Schema{Ref: "#/components/schemas/AdminAreaList"},
								},
							},
						},
					},
					Required: []string{
						"current_page",
						"total",
						"per_page",
						"last_page",
						"from",
						"to",
						"first_page_url",
						"last_page_url",
						"next_page_url",
						"prev_page_url",
						"path",
						"data",
					},
				},
			},
		},
	}
}

func addRestoreEndpoint() *ogen.PathItem {
	return &ogen.PathItem{
		Post: &ogen.Operation{
			Tags:        []string{"admin-areas"},
			Summary:     "Restore a trashed administrative area",
			Description: "Restore a trashed administrative area",
			OperationID: "restoreAdminArea",
			Parameters: []*ogen.Parameter{
				&ogen.Parameter{
					Name:        "id",
					In:          openapi3.ParameterInPath,
					Description: "ID of the AdminArea",
					Required:    true,
					Schema: &ogen.Schema{
						Type:    "integer",
						Minimum: ogen.Num("1"),
					},
				},
			},
			Responses: map[string]*ogen.Response{
				"204": {Description: "AdminArea with requested ID was restored"},
				"400": {Ref: "#/components/responses/400"},
				"404": {Ref: "#/components/responses/404"},
				"409": {Ref: "#/components/responses/409"},
				"500": {Ref: "#/components/responses/500"},
			},
		},
	}
}

func addDeletedAtField(schema *ogen.Schema) {
	schema.Properties = append(
		schema.Properties, ogen.Property{
			Name: "deleted_at",
			Schema: &ogen.Schema{
				Type:        "string",
				Format:      "date-time",
				Nullable:    true,
				Description: "Date and time when the record was deleted",
			},
		},
	)
}
