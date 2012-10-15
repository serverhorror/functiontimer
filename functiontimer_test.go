package functiontimer

import (
	"io/ioutil"
	"log"
	"testing"
"time"
"os"
)

func TestTimer(t *testing.T) {
	// l := log.New(ioutil.Discard, "Test", log.LstdFlags)
	l := log.New(os.Stderr, "Test functiontimer", log.LstdFlags)
	New(l)
}

func sleep3() {
defer RecordTime(time.Now())
	time.Sleep(3 * time.Second)
}

func TestLogger(t *testing.T) {
	l := log.New(ioutil.Discard, "Test", log.LstdFlags)
	New(l)
	sleep3()

	
}
