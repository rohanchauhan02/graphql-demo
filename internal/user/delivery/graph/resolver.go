package graph

import "github.com/rohanchauhan02/graphql-demo/internal/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase user.Usecase
}

func NewResolver(usecase user.Usecase) *Resolver {
	return &Resolver{usecase: usecase}
}
