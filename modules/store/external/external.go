package external

import (
	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/dhanielsales/golang-scaffold/modules/store/external/ideal"
)

type External struct {
	Ideal *ideal.Ideal
}

func New(client *gql.Client) *External {
	return &External{
		Ideal: ideal.New(client),
	}
}
