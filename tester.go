package gqlcheck

import (
	"net/http"

	"github.com/ikawaha/httpcheck"
)

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Errorf(format string, args ...any)
	FailNow()
}

// Tester represents the GraphQL tester.
type Tester struct {
	client *httpcheck.Tester
}

// Test starts a new test with the given *testing.T.
func (c *Checker) Test(t TestingT) *Tester {
	return &Tester{
		client: c.client.Test(t, http.MethodPost, "").
			WithHeader("Content-Type", "application/graphql"),
	}
}

// Check makes request to built request object.
// After request is made, it saves response object for future assertions.
func (tt *Tester) Check() *Tester {
	return &Tester{client: tt.client.Check()}
}
