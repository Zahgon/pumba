package main

import (
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
)

// setupLogging configures the global logrus logger from --log-level, --json,
// --slackhook, and --slackchannel global flags. Called once from before().
func setupLogging(f cliflags.Flags) { _ = "STUB: not implemented"; return }
