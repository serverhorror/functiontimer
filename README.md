functiontimer
=============

functiontimer is a simple package to log the duration of your function
calls.

RecordTime should be used as a `defer`ed function at the beginning of a
function that you want to time.

It is a very naive approach and does nothing more than:

* identify the function,
* measure the elapsed time,
* write a log entry to the specified Logger

Use it like this:

    import (
        "githtub.com/serverhorror/functiontimer"
    )
    var (
        mytimer = functiontimer.New(mylogger)
    )
    func sleep3() {
        defer mytimer.RecordTime(time.Now())
        time.sleep(3 * Second)
    }

