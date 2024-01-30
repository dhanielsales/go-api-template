package example

import (
	"context"

	"github.com/dhanielsales/golang-scaffold/internal/gql"
)

const getImage = `
	query GetImage($extId: String!) {
		allImages(filter: { extId: $extId }) {
			id
			extId
			url
		}
	}
`

type GetImageResponse struct {
	AllImages []struct {
		ID    string `json:"id"`
		ExtId string `json:"extId"`
		Url   string `json:"url"`
	} `mapstructure:"allImages"`
}

func (i *Example) GetImage(ctx context.Context, id string) (*GetImageResponse, error) {
	request := gql.NewRequest(ctx, getImage, map[string]any{
		"extId": gql.NewID(id),
	})

	var response GetImageResponse

	_, err := i.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

const createImage = `
	mutation CreateImage($extId: String!, $url: String!) {
		createImage(extId: $extId, url: $url) {
			id
			extId
			url
		}
	}
`

type CreateImageResponse struct {
	CreateImage struct {
		ID    string `json:"id"`
		ExtId string `json:"extId"`
		Url   string `json:"url"`
	} `mapstructure:"CreateImage"`
}

func (i *Example) CreateImage(ctx context.Context, id, url string) (*CreateImageResponse, error) {
	request := gql.NewRequest(ctx, createImage, map[string]any{
		"extId": id,
		"url":   url,
	})

	var response CreateImageResponse

	_, err := i.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
