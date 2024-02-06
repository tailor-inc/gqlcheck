package gqlcheck

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Response is a struct that represents a GraphQL response.
type Response struct {
	Data   map[string]any   `json:"data,omitempty"`
	Errors []map[string]any `json:"errors,omitempty"`
}

// String returns the string representation of the response.
func (r Response) String() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return string(b)
}

// Cb set a callback function to evaluate the response.
func (tt *Tester) Cb(callback func(*http.Response)) {
	tt.client.Cb(callback)
}

// Response sets the response to the provided variable.
func (tt *Tester) Response(out any) {
	t := tt.client.T()
	tt.client.Cb(func(resp *http.Response) {
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NoError(t, resp.Body.Close())
		require.NoError(t, json.Unmarshal(b, &out))
	})
}

// HasErrors checks if the response contains errors.
func (tt *Tester) HasErrors() *Tester {
	return &Tester{client: tt.client.MatchesJSONQuery(`.errors`)}
}

// HasNoErrors checks if the response does not contain errors.
func (tt *Tester) HasNoErrors() *Tester {
	return &Tester{client: tt.client.NotMatchesJSONQuery(`.errors`)}
}

// HasJSON checks if the response body has the provided value.
func (tt *Tester) HasJSON(expected any) *Tester {
	return &Tester{client: tt.client.HasJSON(expected)}
}

// HasData checks if the response body has the provided data.
func (tt *Tester) HasData(expected any) *Tester {
	return &Tester{client: tt.client.HasJSON(map[string]any{"data": expected})}
}

// ContainsString checks if the response body contains the provided string.
func (tt *Tester) ContainsString(s string) *Tester {
	return &Tester{client: tt.client.ContainsString(s)}
}
