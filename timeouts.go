package timeouts;

import (
	"time"
)

// A timeout object.
// Please use the NewTimeout() function.
type Timeout struct{
	timeouts map[string]time.Time;
}

// Creates a new timeout object.
func NewTimeout() Timeout{
	return Timeout{
		timeouts: make(map[string]time.Time),
	};
}

// Sets a timeout for 'handle' for 'duration'.
func (to *Timeout) SetTimeout(handle string, duration time.Duration){
	to.timeouts[handle] = time.Now().UTC().Add(duration);
}

// Checks if 'handle' is (still) in a timeout
func (to *Timeout) InTimeout(handle string) bool{
	now := time.Now().UTC();

	timeout, ok := to.timeouts[handle];
	return ok && (now.Equal(timeout) || now.Before(timeout));
}

// Forcefully remove 'handle' from a timeout, if any.
func (to *Timeout) RemoveTimeout(handle string){
	delete(to.timeouts, handle);
}
