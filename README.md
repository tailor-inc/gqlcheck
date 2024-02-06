gqlcheck: GraphQL testing library
===

gqlcheck is a library for testing GraphQL servers.   


# Example

You can test the target API by passing the handler of the GraphQL server or the URL of the running GraphQL server.

The following is a test sample of the server included in testdata/gqlgen-todo. gqlgen-todo is the server described in gqlgen's [Getting Started](https://gqlgen.com/getting-started/).

```go
func TestServer(t *testing.T) {
	h := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	h.AddTransport(transport.POST{})

	checker := gqlcheck.New(h, gqlcheck.Debug())
	checker.Test(t).
		WithHeader("Content-Type", "application/json").
		Query(`query {todos {text}}`).
		Check().
		HasStatusOK().
		HasNoError().
		HasData(map[string]any{
			"todos": []any{},
		})
}
```

OUTPUT:
```console
=== RUN   TestServer
2024/02/05 11:31:12 == POST http://127.0.0.1:61506
2024/02/05 11:31:12 >> header map[Content-Type:[application/json]]
2024/02/05 11:31:12 >> body: {"query":"query {todos {text}}"}
2024/02/05 11:31:12 << status: 200 OK
2024/02/05 11:31:12 << body: {"data":{"todos":[]}}
--- PASS: TestServer (0.00s)
PASS
```

---

MIT