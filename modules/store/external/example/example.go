package example

import "github.com/dhanielsales/golang-scaffold/internal/gql"

type Example struct {
	client *gql.Client
}

func New(client *gql.Client) *Example {
	return &Example{
		client: client,
	}
}
