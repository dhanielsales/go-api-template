package external

import (
	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/dhanielsales/golang-scaffold/modules/store/external/example"
)

type External struct {
	Example *example.Example
}

func New(client *gql.Client) *External {
	return &External{
		Example: example.New(client),
	}
}
