package example

import "github.com/dhanielsales/golang-scaffold/internal/gql"

type Example struct {
	client *gql.Client
}

type Image struct {
	ID    string `json:"id"`
	ExtId string `json:"extId"`
	Url   string `json:"url"`
}

func New(client *gql.Client) *Example {
	return &Example{
		client: client,
	}
}
