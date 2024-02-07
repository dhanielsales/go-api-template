package gql

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
)

type Request struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

func NewRequest(query string, variables map[string]any) *Request {
	return &Request{
		Variables: variables,
		Query:     Sanitize(query),
	}
}

func (r *Request) Buffer() (*bytes.Buffer, error) {
	reqBodyByte, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(reqBodyByte), nil
}

func Sanitize(query string) string {
	trimmedString := strings.Trim(query, " \t\n\r")
	pattern := regexp.MustCompile(`\s+`)
	return pattern.ReplaceAllString(trimmedString, " ")
}
