package gqlcheck

import "encoding/json"

type Query struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

func (q Query) String() string {
	b, _ := json.MarshalIndent(q, "", "  ")
	return string(b)
}

func (tt *Tester) Request(q Query) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query":     q.Query,
		"variables": q.Variables,
	})}
}

func (tt *Tester) Query(q string) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query": q,
	})}
}

func (tt *Tester) QueryWithVariables(q string, variables map[string]any) *Tester {
	return &Tester{client: tt.client.WithJSON(map[string]any{
		"query":     q,
		"variables": variables,
	})}
}
