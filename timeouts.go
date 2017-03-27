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

// Checks if 'handle' is (still) in a timeout.
// Also removes the handle from the timeout to save resources.
func (to *Timeout) InTimeout(handle string) bool{
	now := time.Now().UTC();

	timeout, ok := to.timeouts[handle];
	val := ok && (now.Before(timeout) || now.Equal(timeout));

	if(val){
		to.RemoveTimeout(handle);
	}
	return val;
}

// Forcefully remove 'handle' from a timeout, if any.
func (to *Timeout) RemoveTimeout(handle string){
	delete(to.timeouts, handle);
}

// List handles in timeout.
// Also removes unnecessary handles.
// No order guaranteed.
func (to *Timeout) ListTimeouts() []string{
	now := time.Now().UTC();

	var handles []string;
	for handle, timeout := range to.timeouts{
		if(now.Before(timeout) || now.Equal(timeout)){
			handles = append(handles, handle);
		}
	}
	return handles;
}

// Prune passed timeouts.
// Return removed handles.
func (to *Timeout) PruneTimeouts() []string{
	now := time.Now().UTC();

	var handles []string;
	for handle, timeout := range to.timeouts{
		if(now.After(timeout)){
			handles = append(handles, handle);
			to.RemoveTimeout(handle);
		}
	}
	return handles;
}
