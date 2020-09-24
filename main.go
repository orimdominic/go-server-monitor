package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var ival time.Duration
	var url string
	flag.DurationVar(&ival, "ival", 5*time.Second, "interval for monitoring in seconds")
	flag.StringVar(&url, "url", "localhost:12345", "url of the server to monitor")
	flag.Parse()
	fmt.Println(ival, url)
}

/*
	Algorithm
*	print help/instructions
-	Get interval in secs from user or set to default of 5s
-	Get server url
-	start interval timer
-	make http request
-	obtain status code
-	determine if it is error or not
-	set values accordingly (checks, up, down)
-	push output to array (check time, status)
-	log output to console (checks, ups, downs, currentStatus)
-	on exit or interrupt, write array to disk, with details
- create Go, Python and Node.js servers to test
*/
