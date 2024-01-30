package example

import (
	"context"
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
	AllImages []Image `mapstructure:"allImages"`
}

func (i *Example) GetImage(ctx context.Context, id string) (*GetImageResponse, error) {
	// request := gql.NewRequest(ctx, getImage, map[string]any{
	// 	"extId": gql.NewID(id),
	// })

	// var response GetImageResponse

	// _, err := i.client.Do(request, &response)
	// if err != nil {
	// 	return nil, err
	// }

	// return &response, nil

	return &GetImageResponse{
		AllImages: []Image{},
	}, nil
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
	CreateImage Image `mapstructure:"CreateImage"`
}

func (i *Example) CreateImage(ctx context.Context, id, url string) (*CreateImageResponse, error) {
	// request := gql.NewRequest(ctx, createImage, map[string]any{
	// 	"extId": id,
	// 	"url":   url,
	// })

	// var response CreateImageResponse

	// _, err := i.client.Do(request, &response)
	// if err != nil {
	// 	return nil, err
	// }

	// return &response, nil

	return &CreateImageResponse{CreateImage: Image{}}, nil
}
