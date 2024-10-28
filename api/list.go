package api

import (
	"context"

	"github.com/gin-gonic/gin"

	"eidng8.cc/microservices/admin-areas/ent"
	"eidng8.cc/microservices/admin-areas/ent/adminarea"
)

// ListAdminArea List all AdminAreas
// (GET /admin-areas)
func (s Server) ListAdminArea(
	ctx context.Context, request ListAdminAreaRequestObject,
) (ListAdminAreaResponseObject, error) {
	c := ctx.(*gin.Context)
	pageParams := GetPaginationParams(c)
	query := s.EntClient.AdminArea.Query().Order(adminarea.ByID())
	areas, err := GetPage[ent.AdminArea](
		c, context.Background(), query, pageParams,
	)
	if err != nil {
		return nil, err
	}
	data := make([]AdminAreaList, len(areas.Data))
	for i, row := range areas.Data {
		data[i].FromEnt(row)
	}
	return ListAdminArea200JSONResponse{
		CurrentPage:  &areas.CurrentPage,
		FirstPageUrl: &areas.FirstPageUrl,
		From:         areas.From,
		LastPage:     &areas.LastPage,
		LastPageUrl:  &areas.LastPageUrl,
		NextPageUrl:  &areas.NextPageUrl,
		Path:         areas.Path,
		PerPage:      &areas.PerPage,
		PrevPageUrl:  &areas.PrevPageUrl,
		To:           areas.To,
		Total:        areas.Total,
		Data:         data,
	}, nil
}
