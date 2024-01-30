package ideal

import "github.com/dhanielsales/golang-scaffold/internal/gql"

type Ideal struct {
	client *gql.Client
}

func New(client *gql.Client) *Ideal {
	return &Ideal{
		client: client,
	}
}
