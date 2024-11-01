package main

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-paginate"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
	"github.com/eidng8/go-admin-areas/ent/predicate"
	"github.com/eidng8/go-admin-areas/ent/schema"
)

// ListAdminArea List all AdminAreas
// (GET /admin-areas)
func (s Server) ListAdminArea(
	ctx context.Context, request ListAdminAreaRequestObject,
) (ListAdminAreaResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(c)
	query := s.EC.AdminArea.Query().Order(adminarea.ByID())
	qc := newQueryContext(request)
	applyNameFilter(request, query)
	applyAbbrFilter(request, query)
	areas, err := paginate.GetPage[ent.AdminArea](c, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return ListAdminArea200JSONResponse{
		CurrentPage:  areas.CurrentPage,
		FirstPageUrl: areas.FirstPageUrl,
		From:         areas.From,
		LastPage:     areas.LastPage,
		LastPageUrl:  areas.LastPageUrl,
		NextPageUrl:  areas.NextPageUrl,
		Path:         areas.Path,
		PerPage:      areas.PerPage,
		PrevPageUrl:  areas.PrevPageUrl,
		To:           areas.To,
		Total:        areas.Total,
		Data:         mapAdminAreaListFromEnt(areas.Data),
	}, nil
}

func applyNameFilter(
	request ListAdminAreaRequestObject, query *ent.AdminAreaQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(adminarea.NameHasPrefix(*name))
	}
}

func applyAbbrFilter(
	request ListAdminAreaRequestObject, query *ent.AdminAreaQuery,
) {
	abbr := request.Params.Abbr
	if nil == abbr || "" == *abbr {
		return
	}
	vs := strings.Split(*abbr, ",")
	cr := make([]predicate.AdminArea, len(vs))
	for i, a := range vs {
		cr[i] = adminarea.AbbrContains(a)
	}
	query.Where(adminarea.Or(cr...))
}

func newQueryContext(request ListAdminAreaRequestObject) context.Context {
	qc := context.Background()
	if nil != request.Params.Trashed && *request.Params.Trashed {
		qc = schema.IncludeTrashed(qc)
	}
	return qc
}
