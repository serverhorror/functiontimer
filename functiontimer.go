// functiontimer is a naive approach to record the duration of your function
// calls to a specified logger
package functiontimer

import (
	"log"
	"runtime"
	"time"
)

type Timer struct {
	l *log.Logger
}

func New(l *log.Logger) *Timer {

	return &Timer{
		l: l,
	}
}

func Elapsed(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed
}

// RecordTime should be used as a `defer`ed function at the beginning of a
// function that you want to time.
//
// It is a very naive approach and does nothing more than Identify the
// function and measure the elapsed time.
//
// Use it like this:
// 
//	import (
//		"githtub.com/serverhorror/functiontimer"
//	)
//	var (
//		mytimer = functiontimer.New(mylogger)
//	)
//	func sleep3() {
//		defer mytimer.RecordTime(time.Now())
//		time.sleep(3 * Second)
//	}
func (t *Timer) RecordTime(start time.Time) time.Duration {
	pc, _, _, ok := runtime.Caller(1)
	caller := runtime.FuncForPC(pc)
	elapsed := Elapsed(start)
	if ok {
		t.l.Printf("Function `%s` took `%s`",
			caller.Name(), elapsed)
	} else {
		t.l.Printf("UNKNOWN Function took `%s`", elapsed)
	}
	return elapsed

}

// vim: set ts=4 sts=4 fenc=utf-8 syn=go noexpandtab :

