package gql

import (
	"errors"

	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mitchellh/mapstructure"
)

type Response struct {
	Data       any                        `json:"data"`
	Errors     []gqlerrors.FormattedError `json:"errors,omitempty"`
	Extensions map[string]any             `json:"extensions,omitempty"`
}

func (r *Response) To(v any) error {
	return mapstructure.Decode(r.Data, v)
}

func (r *Response) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *Response) Err() error {
	if r.HasErrors() {
		var errs []error = []error{ErrResponseWithErrors}

		for _, err := range r.Errors {
			errs = append(errs, errors.New(err.Message))
		}

		return errors.Join(errs...)
	}

	return nil
}
