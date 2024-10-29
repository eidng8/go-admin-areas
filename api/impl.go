package api

import (
	"context"

	"github.com/gin-gonic/gin"

	"eidng8.cc/microservices/admin-areas/ent"
)

type Server struct {
	EC *ent.Client
}

// Create a new AdminArea
// (POST /admin-areas)
func (s Server) CreateAdminArea(
	ctx context.Context, request CreateAdminAreaRequestObject,
) (CreateAdminAreaResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// Deletes a AdminArea by ID
// (DELETE /admin-areas/{id})
func (s Server) DeleteAdminArea(
	ctx context.Context, request DeleteAdminAreaRequestObject,
) (DeleteAdminAreaResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// Updates a AdminArea
// (PATCH /admin-areas/{id})
func (s Server) UpdateAdminArea(
	ctx context.Context, request UpdateAdminAreaRequestObject,
) (UpdateAdminAreaResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// List attached Childrens
// (GET /admin-areas/{id}/children)
func (s Server) ListAdminAreaChildren(
	ctx context.Context, request ListAdminAreaChildrenRequestObject,
) (ListAdminAreaChildrenResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// Find the attached AdminArea
// (GET /admin-areas/{id}/parent)
func (s Server) ReadAdminAreaParent(
	ctx context.Context, request ReadAdminAreaParentRequestObject,
) (ReadAdminAreaParentResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func NewServer(entClient *ent.Client) Server {
	return Server{EC: entClient}
}

var _ StrictServerInterface = (*Server)(nil)

func NewEngine(mode string, entClient *ent.Client) (*gin.Engine, error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil
	switch mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	server := NewServer(entClient)
	handler := NewStrictHandler(server, []StrictMiddlewareFunc{})
	RegisterHandlers(engine, handler)
	return engine, nil
}
