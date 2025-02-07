// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	stdsql "database/sql"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...any)
	// hooks to execute on mutations.
	hooks *hooks
	// interceptors to execute on queries.
	inters *inters
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ApiToken         []ent.Hook
		Credit           []ent.Hook
		CreditType       []ent.Hook
		DeviceInfo       []ent.Hook
		DisposableEmail  []ent.Hook
		Generation       []ent.Hook
		GenerationModel  []ent.Hook
		GenerationOutput []ent.Hook
		NegativePrompt   []ent.Hook
		Prompt           []ent.Hook
		Scheduler        []ent.Hook
		Upscale          []ent.Hook
		UpscaleModel     []ent.Hook
		UpscaleOutput    []ent.Hook
		User             []ent.Hook
		UserRole         []ent.Hook
	}
	inters struct {
		ApiToken         []ent.Interceptor
		Credit           []ent.Interceptor
		CreditType       []ent.Interceptor
		DeviceInfo       []ent.Interceptor
		DisposableEmail  []ent.Interceptor
		Generation       []ent.Interceptor
		GenerationModel  []ent.Interceptor
		GenerationOutput []ent.Interceptor
		NegativePrompt   []ent.Interceptor
		Prompt           []ent.Interceptor
		Scheduler        []ent.Interceptor
		Upscale          []ent.Interceptor
		UpscaleModel     []ent.Interceptor
		UpscaleOutput    []ent.Interceptor
		User             []ent.Interceptor
		UserRole         []ent.Interceptor
	}
)

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
