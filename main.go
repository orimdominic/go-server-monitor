package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ival, url := getVals()
	fmt.Printf("Monitoring %s every %ds\n", url, ival/time.Second)
	for {
		ping(url, ival)
	}
}

func getVals() (time.Duration, string) {
	var ival time.Duration
	var url string
	flag.DurationVar(&ival, "ival", 5*time.Second, "interval for monitoring in seconds")
	flag.StringVar(&url, "url", "http://localhost:12345", "url of the server to monitor")
	flag.Parse()
	return ival, url
}

func ping(url string, ival time.Duration) {
	client := &http.Client{
		Timeout: ival - time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making http request %s\n", err) // *url.Error.
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		fmt.Println("Success", res.StatusCode)
	} else {
		fmt.Println("Failure", res.StatusCode)
	}
	time.Sleep(ival)
}

/*
	Algorithm

	Implemented
	==============
*	print help/instructions
-	Get interval in secs from user or set to default of 5s
-	Get server url
-	run a sample func every interval
-	make http request
-	obtain status code
-	determine if it is success or not

	Not implemented yet
	====================
-	set values accordingly (pings, up, down)
-	push output to array (ping time, status)
-	log output to console (pings, ups, downs, currentStatus)
-	on exit or interrupt, write array to disk, with details
- create Go, Python and Node.js servers to test
*/
