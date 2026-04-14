package grinex

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	depthURL = "https://grinex.io/api/v1/spot/depth"
	symbol   = "usdta7a5"
)

// Book stores market depth prices extracted from the exchange response.
type Book struct {
	Asks []string
	Bids []string
}

// Client fetches order book data from the Grinex API.
type Client struct {
	httpClient *resty.Client
}

// NewClient constructs a new Grinex API client.
func NewClient(timeout time.Duration, transport http.RoundTripper) *Client {
	stdlibClient := &http.Client{Timeout: timeout}
	if transport != nil {
		stdlibClient.Transport = transport
	}

	client := resty.NewWithClient(stdlibClient)
	client.SetTimeout(timeout)

	return &Client{
		httpClient: client,
	}
}

// FetchBook retrieves and parses the current order book.
func (c *Client) FetchBook(ctx context.Context) (Book, error) {
	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetQueryParam("symbol", symbol).
		Get(depthURL)
	if err != nil {
		return Book{}, fmt.Errorf("fetch grinex order book: %w", err)
	}
	if resp.IsError() {
		return Book{}, fmt.Errorf("fetch grinex order book: unexpected status %s", resp.Status())
	}

	book, err := parseBook(resp.Body())
	if err != nil {
		return Book{}, fmt.Errorf("parse grinex order book from %s: %w", endpointURL(), err)
	}
	return book, nil
}

func endpointURL() string {
	u, err := url.Parse(depthURL)
	if err != nil {
		return depthURL
	}
	query := u.Query()
	query.Set("symbol", symbol)
	u.RawQuery = query.Encode()
	return u.String()
}

type depthResponse struct {
	Asks []depthLevel `json:"asks"`
	Bids []depthLevel `json:"bids"`
}

type depthLevel struct {
	Price string `json:"price"`
}

func parseBook(body []byte) (Book, error) {
	var payload depthResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return Book{}, fmt.Errorf("decode response: %w", err)
	}
	if len(payload.Asks) == 0 || len(payload.Bids) == 0 {
		return Book{}, fmt.Errorf("missing asks or bids")
	}

	asks, err := pricesFromLevels(payload.Asks)
	if err != nil {
		return Book{}, fmt.Errorf("parse asks: %w", err)
	}
	bids, err := pricesFromLevels(payload.Bids)
	if err != nil {
		return Book{}, fmt.Errorf("parse bids: %w", err)
	}

	return Book{Asks: asks, Bids: bids}, nil
}

func pricesFromLevels(levels []depthLevel) ([]string, error) {
	result := make([]string, 0, len(levels))
	for _, level := range levels {
		if level.Price == "" {
			return nil, fmt.Errorf("missing price")
		}
		price, err := parsePrice(level.Price)
		if err != nil {
			return nil, err
		}
		result = append(result, price)
	}
	return result, nil
}

func parsePrice(value string) (string, error) {
	if _, ok := new(big.Rat).SetString(value); !ok {
		return "", fmt.Errorf("invalid price %q", value)
	}
	return value, nil
}
