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
	now := time.Now().UTC();

	timeout, ok := to.timeouts[handle];
	return ok && (now.Equal(timeout) || now.Before(timeout));
}
func (to *Timeout) RemoveTimeout(handle string){
	delete(to.timeouts, handle);
}
