package lifecycle

import (
	"context"
	"syscall"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

const (
	// DefaultKillSignal default kill signal
	DefaultKillSignal = "SIGKILL"
)

// valid Linux signal table
// http://www.comptechdoc.org/os/linux/programming/linux_pgsignals.html
var linuxSignals = map[string]syscall.Signal{
	"SIGHUP":    syscall.SIGHUP,
	"SIGINT":    syscall.SIGINT,
	"SIGQUIT":   syscall.SIGQUIT,
	"SIGILL":    syscall.SIGILL,
	"SIGTRAP":   syscall.SIGTRAP,
	"SIGIOT":    syscall.Signal(0x6), //nolint:mnd // SIGIOT (use number since this signal is not defined for Windows)
	"SIGBUS":    syscall.SIGBUS,
	"SIGFPE":    syscall.SIGFPE,
	"SIGKILL":   syscall.SIGKILL,
	"SIGUSR1":   syscall.Signal(0x1e), //nolint:mnd // SIGUSR1 (use number since this signal is not defined for Windows)
	"SIGSEGV":   syscall.SIGSEGV,
	"SIGUSR2":   syscall.Signal(0x1f), //nolint:mnd // SIGUSR2 (use number since this signal is not defined for Windows)
	"SIGPIPE":   syscall.SIGPIPE,
	"SIGALRM":   syscall.SIGALRM,
	"SIGTERM":   syscall.SIGTERM,
	"SIGCHLD":   syscall.Signal(0x14), //nolint:mnd // SIGCHLD (use number since this signal is not defined for Windows)
	"SIGCONT":   syscall.Signal(0x13), //nolint:mnd // SIGCONT (use number since this signal is not defined for Windows)
	"SIGSTOP":   syscall.Signal(0x11), //nolint:mnd // SIGSTOP (use number since this signal is not defined for Windows)
	"SIGTSTP":   syscall.Signal(0x12), //nolint:mnd // SIGTSTP (use number since this signal is not defined for Windows)
	"SIGTTIN":   syscall.Signal(0x15), //nolint:mnd // SIGTTIN (use number since this signal is not defined for Windows)
	"SIGTTOU":   syscall.Signal(0x16), //nolint:mnd // SIGTTOU (use number since this signal is not defined for Windows)
	"SIGURG":    syscall.Signal(0x10), //nolint:mnd // SIGURG (use number since this signal is not defined for Windows)
	"SIGXCPU":   syscall.Signal(0x18), //nolint:mnd // SIGXCPU (use number since this signal is not defined for Windows)
	"SIGXFSZ":   syscall.Signal(0x19), //nolint:mnd // SIGXFSZ (use number since this signal is not defined for Windows)
	"SIGVTALRM": syscall.Signal(0x1a), //nolint:mnd // SIGVTALRM (use number since this signal is not defined for Windows)
	"SIGPROF":   syscall.Signal(0x1b), //nolint:mnd // SIGPROF (use number since this signal is not defined for Windows)
	"SIGWINCH":  syscall.Signal(0x1c), //nolint:mnd // SIGWINCH (use number since this signal is not defined for Windows)
	"SIGIO":     syscall.Signal(0x17), //nolint:mnd // SIGIO (use number since this signal is not defined for Windows)
}

// killClient is the narrow interface needed by the kill command.
type killClient interface {
	container.Lister
	KillContainer(context.Context, *container.Container, string, bool) error
}

// `docker kill` command
type killCommand struct {
	client  killClient
	names   []string
	pattern string
	labels  []string
	signal  string
	limit   int
	dryRun  bool
}

// NewKillCommand create new Kill Command instance
func NewKillCommand(client killClient, params *chaos.GlobalParams, signal string, limit int) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}

// Run kill command
func (k *killCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}
