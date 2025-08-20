package generic

import (
	"os"
	"os/signal"
	"syscall"
)

var cleanupFns []func()

func AddExitCleanup(fn func()) {
	Append(&cleanupFns, fn)
}

func safeCall(fn func()) {
	defer func() {
		recover()
	}()
	fn()
}

func ExitWithCleanup(code int) {
	Cleanup()
	os.Exit(code)
}

func Cleanup() {
	for i := len(cleanupFns) - 1; i >= 0; i-- {
		safeCall(cleanupFns[i])
	}
}

func SetupSigTermCleanup() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	go func() {
		<-sigChan
		// os.WriteFile("out.txt", []byte(fmt.Sprintf("Signal! %v", sig)), 0644)

		ExitWithCleanup(1)
	}()
}
