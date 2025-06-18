package octogo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const endpointBasePath = "v1/electricity-meter-points"

type MeterService interface {
	GetElectricityMeter(ctx context.Context, mpan string) (*Meter, *Response, error)
}

type Meter struct {
	GSP          string `json:"gsp"`
	MPAN         string `json:"mpan"`
	ProfileClass int8   `json:"profile_class"`
}

type MeterServiceOp struct {
	client *Client
}

func NewMeterService(client *Client) MeterService {
	return &MeterServiceOp{client: client}
}

func (s *MeterServiceOp) GetElectricityMeter(ctx context.Context, mpan string) (*Meter, *Response, error) {
	return s.get(ctx, mpan)
}

func (s *MeterServiceOp) get(ctx context.Context, ID interface{}) (*Meter, *Response, error) {
	path := fmt.Sprintf("%s/%v", endpointBasePath, ID)

	// Build the full URL
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	fullURL := s.client.BaseUrl.ResolveReference(u)

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	meter := new(Meter)
	resp, err := s.client.Do(ctx, req, meter)
	if err != nil {
		return nil, resp, err
	}
	return meter, resp, nil
}
