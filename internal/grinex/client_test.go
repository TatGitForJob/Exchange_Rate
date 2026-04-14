package grinex

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestClientFetchBook(t *testing.T) {
	t.Parallel()

	client := NewClient(2*time.Second, roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if got := r.URL.String(); got != "https://grinex.io/api/v1/spot/depth?symbol=usdta7a5" {
			t.Fatalf("url = %q, want %q", got, "https://grinex.io/api/v1/spot/depth?symbol=usdta7a5")
		}
		if got := r.URL.Query().Get("symbol"); got != "usdta7a5" {
			t.Fatalf("symbol = %q, want usdta7a5", got)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"asks":[{"price":"95.10","volume":"1"},{"price":"95.20","volume":"2"}],"bids":[{"price":"94.90","volume":"3"},{"price":"94.80","volume":"4"}]}`)),
		}, nil
	}))

	book, err := client.FetchBook(context.Background())
	if err != nil {
		t.Fatalf("FetchBook() error = %v", err)
	}
	if len(book.Asks) != 2 || len(book.Bids) != 2 {
		t.Fatalf("unexpected book lengths: %+v", book)
	}
	if book.Asks[0] != "95.10" || book.Bids[0] != "94.90" {
		t.Fatalf("unexpected prices: %+v", book)
	}
}

func TestEndpointURL(t *testing.T) {
	t.Parallel()

	if got := endpointURL(); got != "https://grinex.io/api/v1/spot/depth?symbol=usdta7a5" {
		t.Fatalf("endpointURL() = %q", got)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestParseBookErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		body string
	}{
		{
			name: "missing bids",
			body: `{"asks":[{"price":"95.10","volume":"1"}]}`,
		},
		{
			name: "missing level price",
			body: `{"asks":[{"volume":"1"}],"bids":[{"price":"94.90","volume":"1"}]}`,
		},
		{
			name: "non numeric price",
			body: `{"asks":[{"price":"oops","volume":"1"}],"bids":[{"price":"94.90","volume":"1"}]}`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if _, err := parseBook([]byte(tc.body)); err == nil {
				t.Fatal("parseBook() error = nil, want non-nil")
			}
		})
	}
}
