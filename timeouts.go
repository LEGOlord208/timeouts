package timeouts;

import (
	"time"
)

type Timeout struct{
	timeouts map[string]time.Time;
}

func NewTimeout() Timeout{
	return Timeout{
		timeouts: make(map[string]time.Time),
	};
}

func (to *Timeout) SetTimeout(handle string, duration time.Duration){
	to.timeouts[handle] = time.Now().UTC().Add(duration);
}
func (to *Timeout) InTimeout(handle string) bool{
	timeout, ok := to.timeouts[handle];
	return !ok || time.Now().UTC().After(timeout);
}
