//go:build js

package generic

// SetupSigTermCleanup is a no-op under GOOS=js (no POSIX signals).
func SetupSigTermCleanup() {}
