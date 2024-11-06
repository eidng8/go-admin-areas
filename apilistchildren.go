package main

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"entgo.io/ent/dialect/sql"
	"github.com/eidng8/go-paginate"
	"github.com/eidng8/go-softdelete"
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
	query := s.EC.Debug().AdminArea.Query().Order(adminarea.ByID()).Where(
		adminarea.HasParentWith(adminarea.ID(uint32(request.Id))),
	)
	qc := softdelete.NewSoftDeleteQueryContext(request.Params.Trashed, ctx)
	applyChildrenNameFilter(request, query)
	applyChildrenAbbrFilter(request, query)
	return getPage(ctx, qc, query)
}

func getPage(
	ctx context.Context, qc context.Context,
	query *ent.AdminAreaQuery,
) (ListAdminAreaChildrenResponseObject, error) {
	gc := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(gc)
	areas, err := paginate.GetPage[ent.AdminArea](gc, qc, query, pageParams)
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

func descendantsQuery(query *ent.AdminAreaQuery) *ent.AdminAreaQuery {
	return query.Where(
		func(stmt *sql.Selector) {
			child := sql.Table(adminarea.Table)
			parent := sql.Table(adminarea.Table)
			view := fmt.Sprintf("%s_tree", adminarea.Table)
			keys := []string{adminarea.FieldID, adminarea.FieldParentID}
			cte := sql.WithRecursive(view, keys...)
			pid := cte.C(adminarea.FieldID)
			cte.As(
				sql.Select(parent.Columns(keys...)...).From(child).
					Where(sql.IsNull(parent.C(adminarea.FieldParentID))).
					UnionAll(
						sql.Select(child.Columns(keys...)...).From(child).
							Join(cte).On(child.C(adminarea.FieldParentID), pid),
					),
			)
			stmt.Prefix(cte).Join(cte).On(stmt.C(adminarea.FieldID), pid)
		},
	)
}
