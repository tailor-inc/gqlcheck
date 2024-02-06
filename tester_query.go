package gqlcheck

import "encoding/json"

// Query is a struct to represent a query.
type Query struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// String returns the string representation of the query.
func (q Query) String() string {
	b, _ := json.MarshalIndent(q, "", "  ")
	return string(b)
}

// Request sets the query and variables to the request.
func (tt *Tester) Request(q Query) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query":     q.Query,
		"variables": q.Variables,
	})}
}

// Query sets the query to the request.
func (tt *Tester) Query(q string) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query": q,
	})}
}

// QueryWithVariables sets the query and variables to the request.
func (tt *Tester) QueryWithVariables(q string, variables map[string]any) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query":     q,
		"variables": variables,
	})}
}
