package main

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/eidng8/go-ent/paginate"
	"github.com/eidng8/go-url"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
	"github.com/eidng8/go-admin-areas/ent/predicate"
)

// ListAdminAreaChildren List attached Children
// (GET /admin-areas/{id}/children)
func (s Server) ListAdminAreaChildren(
	ctx context.Context, request ListAdminAreaChildrenRequestObject,
) (ListAdminAreaChildrenResponseObject, error) {
	gc := ctx.(*gin.Context)
	query := s.EC.AdminArea.Query().Order(adminarea.ByID())
	applyChildrenNameFilter(request, query)
	applyChildrenAbbrFilter(request, query)
	id := request.Id
	if nil != request.Params.Recurse && *request.Params.Recurse {
		return getDescendants(gc, ctx, query, id)
	}
	return getPage(gc, ctx, query, id)
}

func getPage(
	gc *gin.Context, qc context.Context, query *ent.AdminAreaQuery, id uint32,
) (ListAdminAreaChildrenResponseObject, error) {
	query.Where(adminarea.HasParentWith(adminarea.ID(id)))
	pageParams := paginate.GetPaginationParams(gc)
	areas, err := paginate.GetPage[ent.AdminArea](gc, qc, query, pageParams)
	if err != nil {
		return nil, err
	}
	return mapPage[ListAdminAreaChildren200JSONResponse](areas), nil
}

func getDescendants(
	gc *gin.Context, qc context.Context, query *ent.AdminAreaQuery, id uint32,
) (ListAdminAreaChildrenResponseObject, error) {
	areas, err := query.QueryChildrenRecursive(id).All(qc)
	if err != nil {
		return nil, err
	}
	count := len(areas)
	req := gc.Request
	u := paginate.UrlWithoutPageParams(req)
	u = url.WithQueryParam(*u, "recurse", "1")
	return ListAdminAreaChildren200JSONResponse{
		CurrentPage:  1,
		FirstPageUrl: u.String(),
		From:         1,
		LastPage:     1,
		LastPageUrl:  "",
		NextPageUrl:  "",
		Path:         url.RequestBaseUrl(req).String(),
		PerPage:      count,
		PrevPageUrl:  "",
		To:           count,
		Total:        count,
		Data:         mapAdminAreaListFromEnt(areas),
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
