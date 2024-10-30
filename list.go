package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-paginate"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/adminarea"
)

// ListAdminArea List all AdminAreas
// (GET /admin-areas)
func (s Server) ListAdminArea(
	ctx context.Context, request ListAdminAreaRequestObject,
) (ListAdminAreaResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := paginate.GetPaginationParams(c)
	query := s.EC.AdminArea.Query().Order(adminarea.ByID())
	areas, err := paginate.GetPage[ent.AdminArea](
		c, context.Background(), query, pageParams,
	)
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
