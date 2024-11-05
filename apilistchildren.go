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
)

// ListAdminAreaChildren List attached Children
// (GET /admin-areas/{id}/children)
func (s Server) ListAdminAreaChildren(
	ctx context.Context, request ListAdminAreaChildrenRequestObject,
) (ListAdminAreaChildrenResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(c)
	query := s.EC.AdminArea.Query().Order(adminarea.ByID()).Where(
		adminarea.HasParentWith(adminarea.ID(uint32(request.Id))),
	)
	qc := newQueryContext(request.Params.Trashed, ctx)
	applyChildrenNameFilter(request, query)
	applyChildrenAbbrFilter(request, query)
	areas, err := paginate.GetPage[ent.AdminArea](c, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return ListAdminAreaChildren200JSONResponse{
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

func applyChildrenNameFilter(
	request ListAdminAreaChildrenRequestObject, query *ent.AdminAreaQuery,
) {
	name := request.Params.Name
	if name != nil && utf8.RuneCountInString(*name) > 1 {
		query.Where(adminarea.NameHasPrefix(*name))
	}
}

func applyChildrenAbbrFilter(
	request ListAdminAreaChildrenRequestObject, query *ent.AdminAreaQuery,
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
