package openweather

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var errNilModel = errors.New("model is nil")

type sender struct {
	*http.Client
}

func newSender(client *http.Client) *sender {
	return &sender{
		client,
	}
}

func (s *sender) do(ctx context.Context, method, url string, model interface{}, body io.Reader) error {
	if model == nil {
		return errNilModel
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}

	response, err := s.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return json.NewDecoder(response.Body).Decode(&model)
	default:
		return httpErrorFromJson(response.Body)
	}
}
