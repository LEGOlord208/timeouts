# Timeouts
This is the most pointless package you've **ever** seen.  
All it does is literally just timeouts.  

[GoDoc](https://godoc.org/github.com/legOlord208/timeouts)

## Example
```Go
package main;

import (
	"fmt"
	"github.com/jD91mZM2/timeouts"
	"time"
)

func main(){
	timeout := timeouts.NewTimeout();
	timeout.SetTimeout("lol", time.Second);
	fmt.Println(timeout.InTimeout("lol"));
	time.Sleep(time.Second * 2);
	fmt.Println(timeout.InTimeout("lol"));
}
```
