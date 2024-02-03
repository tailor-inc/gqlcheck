package gqlcheck

import (
	"net/http"
	"time"

	"github.com/ikawaha/httpcheck"
)

type Checker struct {
	client *httpcheck.Checker
}

// Option represents the option for the checker.
type Option = httpcheck.Option

// ClientTimeout sets the client timeout.
func ClientTimeout(d time.Duration) Option {
	return httpcheck.ClientTimeout(d)
}

func Debug() Option {
	return httpcheck.Debug()
}

// New returns a new Checker for the given handler.
func New(h http.Handler, opts ...Option) *Checker {
	return &Checker{client: httpcheck.New(h, opts...)}
}

// NewExternal returns a new Checker for external server.
func NewExternal(endpoint string, opts ...Option) *Checker {
	return &Checker{client: httpcheck.NewExternal(endpoint, opts...)}
}
