package octogo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const endpointBasePath = "v1/electricity-meter-points"

type MeterService interface {
	GetElectricityMeter(ctx context.Context, mpan string) (*Meter, *Response, error)
	GetConsumption(ctx context.Context, mpan, serialNumber string, opts *ConsumptionOptions) (*ConsumptionResponse, *Response, error)
}

type Meter struct {
	GSP          string `json:"gsp"`
	MPAN         string `json:"mpan"`
	ProfileClass int8   `json:"profile_class"`
}

type ConsumptionReading struct {
	Consumption   float64   `json:"consumption"`
	IntervalStart time.Time `json:"interval_start"`
	IntervalEnd   time.Time `json:"interval_end"`
}

type ConsumptionResponse struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []ConsumptionReading `json:"results"`
}

type ConsumptionOptions struct {
	PeriodFrom *time.Time
	PeriodTo   *time.Time
	PageSize   *int
	Page       *int
	OrderBy    *string
	GroupBy    *string
}

// MeterServiceOp handles operations on electricity meters
type MeterServiceOp struct {
	client *Client
}

func NewMeterService(client *Client) MeterService {
	return &MeterServiceOp{client: client}
}

func (s *MeterServiceOp) GetElectricityMeter(ctx context.Context, mpan string) (*Meter, *Response, error) {
	return s.get(ctx, mpan)
}

func (s *MeterServiceOp) GetConsumption(ctx context.Context, mpan, serialNumber string, opts *ConsumptionOptions) (*ConsumptionResponse, *Response, error) {
	return s.getConsumption(ctx, mpan, serialNumber, opts)
}

func (s *MeterServiceOp) get(ctx context.Context, id string) (*Meter, *Response, error) {
	path := fmt.Sprintf("%s/%s", endpointBasePath, id)

	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	fullURL := s.client.BaseUrl.ResolveReference(u)

	req, err := http.NewRequest(http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	meter := &Meter{}
	resp, err := s.client.Do(ctx, req, meter)
	if err != nil {
		return nil, resp, err
	}
	return meter, resp, nil
}

func (s *MeterServiceOp) getConsumption(ctx context.Context, mpan, serialNumber string, opts *ConsumptionOptions) (*ConsumptionResponse, *Response, error) {
	path := fmt.Sprintf("%s/%s/meters/%s/consumption/", endpointBasePath, mpan, serialNumber)

	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// Add query parameters if options are provided
	if opts != nil {
		query := u.Query()

		if opts.PeriodFrom != nil {
			query.Set("period_from", opts.PeriodFrom.Format(time.RFC3339))
		}
		if opts.PeriodTo != nil {
			query.Set("period_to", opts.PeriodTo.Format(time.RFC3339))
		}
		if opts.PageSize != nil {
			query.Set("page_size", fmt.Sprintf("%d", *opts.PageSize))
		}
		if opts.Page != nil {
			query.Set("page", fmt.Sprintf("%d", *opts.Page))
		}
		if opts.OrderBy != nil {
			query.Set("order_by", *opts.OrderBy)
		}
		if opts.GroupBy != nil {
			query.Set("group_by", *opts.GroupBy)
		}

		u.RawQuery = query.Encode()
	}

	fullURL := s.client.BaseUrl.ResolveReference(u)

	req, err := http.NewRequest(http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	consumption := &ConsumptionResponse{}
	resp, err := s.client.Do(ctx, req, consumption)
	if err != nil {
		return nil, resp, err
	}
	return consumption, resp, nil
}
