package gql

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dhanielsales/golang-scaffold/internal/utils"
)

type Client struct {
	URL        string
	httpClient *http.Client
}

func NewClient(url string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		URL:        url,
		httpClient: httpClient,
	}
}

func (c *Client) Do(ctx context.Context, req *Request, target any) (*Response, error) {
	if !utils.IsPointer(target) {
		return nil, ErrTargetIsNotPointer
	}

	buff, err := req.Buffer()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", c.URL, buff)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responsePayload Response
	if err := json.Unmarshal(bytes, &responsePayload); err != nil {
		return nil, err
	}

	if err := responsePayload.Err(); err != nil {
		return nil, err
	}

	if err := responsePayload.To(target); err != nil {
		return nil, err
	}

	return &responsePayload, nil
}
