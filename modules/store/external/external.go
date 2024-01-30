package external

import (
	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/dhanielsales/golang-scaffold/modules/store/external/example"
)

type External struct {
	Ideal *example.Example
}

func New(client *gql.Client) *External {
	return &External{
		Ideal: example.New(client),
	}
}
