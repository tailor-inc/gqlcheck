package gqlcheck

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Response struct {
	Data   map[string]any
	Errors []map[string]any
}

func (tt *Tester) Cb(callback func(*http.Response)) {
	tt.client.Cb(callback)
}

func (tt *Tester) Response(out *Response) {
	t := tt.client.T()
	tt.client.Cb(func(resp *http.Response) {
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NoError(t, resp.Body.Close())
		require.NoError(t, json.Unmarshal(b, &out))
	})
}

func (tt *Tester) HasError() *Tester {
	return &Tester{client: tt.client.MatchesJSONQuery(`.errors`)}
}

func (tt *Tester) HasNoError() *Tester {
	return &Tester{client: tt.client.NotMatchesJSONQuery(`.errors`)}
}
