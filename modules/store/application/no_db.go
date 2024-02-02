package application

import (
	"github.com/dhanielsales/golang-scaffold/entity"
	"golang.org/x/net/context"
)

func (s *StoreService) GetManyNoDb(ctx context.Context) (*[]entity.NoDb, error) {
	return &[]entity.NoDb{
		{
			ID:   "1",
			Name: "NoDb 1",
		},
		{
			ID:   "2",
			Name: "NoDb 2",
		},
		{
			ID:   "3",
			Name: "NoDb 3",
		},
		{
			ID:   "4",
			Name: "NoDb 4",
		},
		{
			ID:   "5",
			Name: "NoDb 5",
		},
	}, nil
}
