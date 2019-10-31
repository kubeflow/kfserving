package logger

import "fmt"

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan LogRequest, 100)

func QueueLogRequest(req LogRequest) error {
	WorkQueue <- req
	fmt.Println("Work request queued")
	return nil
}
