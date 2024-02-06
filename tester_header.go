package gqlcheck

// WithHeader set header in the request.
func (tt *Tester) WithHeader(key, value string) *Tester {
	return &Tester{client: tt.client.WithHeader(key, value)}
}

// WithHeaders sets header in the request.
func (tt *Tester) WithHeaders(headers map[string]string) *Tester {
	return &Tester{client: tt.client.WithHeaders(headers)}
}
