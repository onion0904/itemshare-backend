package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/onion0904/CarShareSystem/app/middleware"
)

var Directive DirectiveRoot = DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

//ここではログイン済みかどうかを判断するだけ。認証ではない。
func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if _, ok := middleware.GetUserID(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}