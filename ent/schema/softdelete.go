package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	gen "github.com/eidng8/go-admin-areas/ent"
	"github.com/eidng8/go-admin-areas/ent/hook"
	"github.com/eidng8/go-admin-areas/ent/intercept"
)

// SoftDeleteMixin implements the soft delete pattern for schemas.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
// Once you declare "deleted_at" in here, you MUST DELETE IT from the entity that will use that Mixin
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").Optional().Nillable().
			Annotations(entoas.Skip(true)),
	}
}

type softDeleteKey struct{}

// IncludeTrashed returns a new context that skips the soft-delete interceptor/mutators.
func IncludeTrashed(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(
			func(ctx context.Context, q intercept.Query) error {
				if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
					return nil
				}
				d.P(q)
				return nil
			},
		),
	}
}

// Hooks of the SoftDeleteMixin.
func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(
					func(ctx context.Context, m ent.Mutation) (
						ent.Value, error,
					) {
						if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
							return next.Mutate(ctx, m)
						}
						mx, ok := m.(interface {
							SetOp(ent.Op)
							Client() *gen.Client
							SetDeletedAt(time.Time)
							WhereP(...func(*sql.Selector))
						})
						if !ok {
							return nil, fmt.Errorf(
								"unexpected mutation type %T %+v", m, m,
							)
						}
						d.P(mx)
						mx.SetOp(ent.OpUpdate)
						mx.SetDeletedAt(time.Now())
						return mx.Client().Mutate(ctx, m)
					},
				)
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(
					func(ctx context.Context, m ent.Mutation) (
						ent.Value, error,
					) {
						if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
							return next.Mutate(ctx, m)
						}
						mx, ok := m.(interface{ WhereP(...func(*sql.Selector)) })
						if !ok {
							return nil, fmt.Errorf(
								"unexpected mutation type %T %+v", m, m,
							)
						}
						d.P(mx)
						return next.Mutate(ctx, m)
					},
				)
			},
			ent.OpUpdateOne|ent.OpUpdate,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[0].Descriptor().Name),
	)
}
