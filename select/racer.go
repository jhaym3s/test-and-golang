package racer

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) string {
// 	startA := time.Now()
// 	http.Get(a)
// 	aDuration := time.Since(startA)
// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b

// }

// func Racer1(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

///

func Racer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	//!Why struct{} and not another type like a bool? Well, a 
	//!"chan struct{}" is the smallest data type available from a 
	//!memory perspective so we get no allocation versus a bool.
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

//!What select lets you do is wait on multiple channels. 
//!The first one to send a value "wins" and the code underneath the case is executed.


///Notice how we have to use make when creating a channel; 
//rather than say var ch chan struct{}. When you use var the variable will be 
//initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.