package timeouts

import (
	"time"
)

// Timeout is a timeout object.
// Use the NewTimeout() function.
type Timeout struct {
	timeouts map[string]time.Time
}

// NewTimeout creates a new timeout object.
func NewTimeout() Timeout {
	return Timeout{
		timeouts: make(map[string]time.Time),
	}
}

// SetTimeout sets a timeout for 'handle' with 'duration'.
func (to *Timeout) SetTimeout(handle string, duration time.Duration) {
	now := time.Now()
	to.SetTimeoutAt(handle, now.Add(duration))
}

// SetTimeoutAt sets a timeout for 'handle' at 'time'
func (to *Timeout) SetTimeoutAt(handle string, at time.Time) {
	to.timeouts[handle] = at
}

// InTimeout checks if 'handle' is (still) in a timeout.
// Also removes the handle from the timeout to save resources.
func (to *Timeout) InTimeout(handle string) bool {
	now := time.Now()

	timeout, ok := to.timeouts[handle]
	val := ok && (now.Before(timeout) || now.Equal(timeout))

	if !val {
		to.RemoveTimeout(handle)
	}
	return val
}

// RemoveTimeout forcefully removes 'handle' from a timeout, if any.
func (to *Timeout) RemoveTimeout(handle string) {
	delete(to.timeouts, handle)
}

// Timeouts returns a copy of timeouts.
// No guarantee on order, or even if they're passed or not.
// Should nearly always be used together with PruneTimeouts().
func (to *Timeout) Timeouts() map[string]time.Time {
	to.PruneTimeouts()

	copy := make(map[string]time.Time, len(to.timeouts))
	for key, val := range to.timeouts {
		copy[key] = val
	}
	return copy
}

// PruneTimeouts removes passed timeouts.
// Returns removed handles.
func (to *Timeout) PruneTimeouts() []string {
	now := time.Now()

	var handles []string
	for handle, timeout := range to.timeouts {
		if now.After(timeout) {
			handles = append(handles, handle)
			to.RemoveTimeout(handle)
		}
	}
	return handles
}
