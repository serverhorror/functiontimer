package functiontimer

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// l := log.New(ioutil.Discard, "Test", log.LstdFlags)
	l := log.New(os.Stderr, "Test functiontimer", log.LstdFlags)
	New(l)
}

func sleep1() {
	time.Sleep(1 * time.Second)
}

func sleep1WithTimer(tmr *Timer) {
	defer tmr.RecordTime(time.Now())
	time.Sleep(1 * time.Second)
}

func TestElapsed(t *testing.T) {
	var elapsed time.Duration
	expected := 1.0
	defer func(start time.Time) {
		elapsed = Elapsed(start)
		if math.Floor(elapsed.Seconds()) != expected {
			t.Fatalf("Got %2f, expected %2f", elapsed.Seconds(), expected)
		} else {
			t.Logf("Got %2f, expected %2f", elapsed.Seconds(), expected)
		}
	}(time.Now())
	sleep1()
}

func TestLogger(t *testing.T) {
	l := log.New(ioutil.Discard, "Test", log.LstdFlags)
	tmr := New(l)
	sleep1WithTimer(tmr)

}

// vim: set ts=4 sts=4 fenc=utf-8 syn=go noexpandtab :

