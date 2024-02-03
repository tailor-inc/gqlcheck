package gqlcheck

import "net/http"

func (tt *Tester) HasStatus(status int) *Tester {
	return &Tester{client: tt.client.HasStatus(status)}
}

func (tt *Tester) HasStatusOK() *Tester {
	return &Tester{client: tt.client.HasStatus(http.StatusOK)}
}
