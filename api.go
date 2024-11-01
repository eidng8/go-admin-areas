package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/schema"
)

type Server struct {
	EC *ent.Client
}

func (s Server) RestoreAdminArea(
	ctx context.Context, request RestoreAdminAreaRequestObject,
) (RestoreAdminAreaResponseObject, error) {
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

var _ StrictServerInterface = (*Server)(nil)

func newServer(entClient *ent.Client) Server {
	return Server{EC: entClient}
}

func newEngine(mode string, entClient *ent.Client) (*gin.Engine, error) {
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
	server := newServer(entClient)
	handler := NewStrictHandler(server, []StrictMiddlewareFunc{})
	RegisterHandlers(engine, handler)
	return engine, nil
}

func newQueryContext(withTrashed *bool, ctx context.Context) context.Context {
	var qc context.Context
	if nil == ctx {
		qc = context.Background()
	} else {
		qc = ctx
	}
	if nil != withTrashed && *withTrashed {
		qc = schema.IncludeTrashed(qc)
	}
	return qc
}
