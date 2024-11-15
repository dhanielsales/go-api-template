package db

import "database/sql"

//go:generate mockgen -source ./wrapper.go -destination ./wrapper_mock.go -package $GOPACKAGE

type QueryWrapper interface {
	Querier
	WithTx(tx *sql.Tx) QueryWrapper
}

type wrapper struct {
	*Queries
}

func (q *wrapper) WithTx(tx *sql.Tx) QueryWrapper {
	return &wrapper{
		Queries: q.Queries.WithTx(tx),
	}
}

func NewQueryWrapper(db DBTX) QueryWrapper {
	return &wrapper{
		Queries: New(db),
	}
}
