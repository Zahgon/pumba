package cliflags

import (
	"time"

	"github.com/urfave/cli"
)

// V1 adapts urfave/cli v1's *cli.Context to the Flags interface. The Context
// pointer is exposed so call sites can still drop down to cli-specific helpers
// when the abstraction would be lossy (e.g. ParseGlobalParams reads --random
// from the root via GlobalBool).
type V1 struct {
	Ctx *cli.Context
}

// NewV1 wraps a *cli.Context as Flags. Returning Flags (interface) lets
// callers swap adapters without touching parser signatures.
func NewV1(ctx *cli.Context) Flags {
	_ = "STUB: not implemented"

	// NewV1FromApp wraps the application-level (root) *cli.Context as Flags.
	// Use it from app.Before / app.After callbacks and any other call site that
	// already holds the root context and reads global flags (e.g. cmd/main.go).
	// The constructor walks to the root via Global() so that a subcommand context
	// passed by mistake still yields correct global-flag reads instead of silently
	// returning subcommand-scoped values.
	return *new(Flags)
}

func NewV1FromApp(ctx *cli.Context) Flags { _ = "STUB: not implemented"; return *new(Flags) }

// String returns the value of the named string flag.
func (f V1) String(name string) string { _ = "STUB: not implemented"; return "" }

// Bool returns the value of the named bool flag (defaults false).
func (f V1) Bool(name string) bool { _ = "STUB: not implemented"; return false }

// BoolT returns the value of the named "bool-true" flag (defaults true).
func (f V1) BoolT(name string) bool { _ = "STUB: not implemented"; return false }

// Duration returns the value of the named duration flag.
func (f V1) Duration(name string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Int returns the value of the named int flag.
func (f V1) Int(name string) int { _ = "STUB: not implemented"; return 0 }

// Float64 returns the value of the named float64 flag.
func (f V1) Float64(name string) float64 { _ = "STUB: not implemented"; return 0 }

// StringSlice returns the value of the named string-slice flag.
func (f V1) StringSlice(name string) []string { _ = "STUB: not implemented"; return nil }

// Args returns positional arguments as a plain []string.
func (f V1) Args() []string { _ = "STUB: not implemented"; return nil }

// Parent returns the parent subcommand's flags, or nil at the root.
func (f V1) Parent() Flags { _ = "STUB: not implemented"; return *new(Flags) }

// Global walks up the parent chain and returns the root context's flags.
// Mirrors urfave/cli v1's GlobalX semantics enough for parser use.
func (f V1) Global() Flags { _ = "STUB: not implemented"; return *new(Flags) }
