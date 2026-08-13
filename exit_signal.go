//go:build !js

package generic

import (
	"os"
	"os/signal"
	"syscall"
)

// SetupSigTermCleanup runs ExitWithCleanup on SIGINT/SIGTERM/SIGQUIT/SIGABRT.
// No-op on platforms without these signals (see exit_signal_js.go).
func SetupSigTermCleanup() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	go func() {
		<-sigChan
		ExitWithCleanup(1)
	}()
}
