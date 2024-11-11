package main

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/eidng8/go-ent/softdelete"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-ent/paginate"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
	"github.com/eidng8/go-admin-areas/ent/predicate"
)

// ListAdminArea List all AdminAreas
// (GET /admin-areas)
func (s Server) ListAdminArea(
	ctx context.Context, request ListAdminAreaRequestObject,
) (ListAdminAreaResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(c)
	query := s.EC.AdminArea.Query().Order(adminarea.ByID())
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	applyNameFilter(request, query)
	applyAbbrFilter(request, query)
	areas, err := paginate.GetPage[ent.AdminArea](c, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return mapPage[ListAdminArea200JSONResponse](areas), nil
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
