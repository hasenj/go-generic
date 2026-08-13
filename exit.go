package generic

import "os"

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
