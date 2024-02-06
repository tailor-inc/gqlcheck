package gqlcheck

import "net/http"

// HasStatus checks if the response status code is the given status code.
func (tt *Tester) HasStatus(status int) *Tester {
	return &Tester{client: tt.client.HasStatus(status)}
}

// HasStatusOK checks if the response status code is 200.
func (tt *Tester) HasStatusOK() *Tester {
	return &Tester{client: tt.client.HasStatus(http.StatusOK)}
}
