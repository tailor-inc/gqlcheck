package gqlcheck

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ikawaha/httpcheck"
	"github.com/stretchr/testify/require"
)

type Checker struct {
	client *httpcheck.Checker
}

// Option represents the option for the checker.
type Option = httpcheck.Option

func ClientTimeout(d time.Duration) Option {
	return httpcheck.ClientTimeout(d)
}

// NewExternal returns a new Checker for external server.
func NewExternal(url string, opts ...Option) *Checker {
	return &Checker{client: httpcheck.NewExternal(url, opts...)}
}

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Errorf(format string, args ...any)
	FailNow()
}

type Tester struct {
	client *httpcheck.Tester
}

func (c *Checker) Test(t TestingT) *Tester {
	return &Tester{
		client: c.client.Test(t, http.MethodPost, c.client.GetURL()).
			WithHeader("Content-Type", "application/graphql"),
	}
}

func (tt *Tester) Check() *Tester {
	return &Tester{client: tt.client.Check()}
}

// WithBasicAuth is an alias to set basic auth in the request header.
func (tt *Tester) WithBasicAuth(user, pass string) *Tester {
	return &Tester{client: tt.client.WithBasicAuth(user, pass)}
}

// WithBearerAuth is an alias to set bearer auth in the request header.
func (tt *Tester) WithBearerAuth(token string) *Tester {
	return &Tester{client: tt.client.WithHeader("Authorization", "Bearer: "+token)}
}

func (tt *Tester) Query(q string) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query": q,
	})}
}

func (tt *Tester) QueryWithVariables(q, v map[string]any) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query":     q,
		"variables": v,
	})}
}

func (tt *Tester) StatusOK() *Tester {
	return &Tester{client: tt.client.HasStatus(http.StatusOK)}
}

func (tt *Tester) Response(t TestingT, gqlResponse map[string]any) {
	tt.client.Cb(func(resp *http.Response) {
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(b, &gqlResponse))
	})
}
