package gqlcheck

func (tt *Tester) WithHeader(key, value string) *Tester {
	return &Tester{client: tt.client.WithHeader(key, value)}
}

func (tt *Tester) WithHeaders(headers map[string]string) *Tester {
	return &Tester{client: tt.client.WithHeaders(headers)}
}
