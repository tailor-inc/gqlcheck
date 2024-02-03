package gqlcheck

// WithBasicAuth is an alias to set basic auth in the request header.
func (tt *Tester) WithBasicAuth(user, pass string) *Tester {
	return &Tester{client: tt.client.WithBasicAuth(user, pass)}
}

// WithBearerAuth is an alias to set bearer auth in the request header.
func (tt *Tester) WithBearerAuth(token string) *Tester {
	return &Tester{client: tt.client.WithHeader("Authorization", "Bearer: "+token)}
}
