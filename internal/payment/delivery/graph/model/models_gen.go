// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Payment struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"createdAt"`
}

type PaymentInput struct {
	Amount float64 `json:"amount"`
	UserID string  `json:"userId"`
}

type Query struct {
}
