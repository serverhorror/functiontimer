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
func RecordTime(start time.Time) {
	pc, _, _, ok := runtime.Caller(0)
	caller := runtime.FuncForPC(pc)
	elapsed := time.Since(start)
	if ok {
		log.Printf("Function `%s` took `%s`",
			caller.Name(), elapsed)
	} else {
		log.Printf("UNKNOWN Function took `%s`", elapsed)
	}

}