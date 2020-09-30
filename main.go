package main

import (
  "flag"
  "fmt"
  "log"
  "math/rand"
  "net/http"
  "time"
)

func main() {
  ival, useTstSrv, url := getVals()
  if useTstSrv {
    go strtTstSrvr()
  }
  fmt.Printf("Monitoring %s every %ds\n", url, ival/time.Second)
  for {
    ping(url, ival)
  }
}

func getVals() (time.Duration, bool, string) {
  var ival time.Duration
  var testSrv bool
  var url string
  flag.DurationVar(&ival, "ival", 5*time.Second, "interval for monitoring in seconds")
  flag.BoolVar(&testSrv, "tstsrv", true, "use test server")
  flag.StringVar(&url, "url", "http://localhost:8080", "url of the server to monitor")
  flag.Parse()
  flag.Usage()
  fmt.Print("Use CTRL + C to stop the application\n\n")
  if testSrv {
    return ival, testSrv, "http://localhost:8080"
  }
  return ival, testSrv, url
}

func ping(url string, ival time.Duration) {
  client := &http.Client{
    Timeout: ival - time.Second,
  }
  res, err := client.Get(url)
  if err != nil {
    fmt.Printf("Error making http request %s\n", err) //! handle *url.Error.
  }
  defer res.Body.Close()
  if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
    fmt.Println("Success", res.StatusCode)
  } else {
    fmt.Println("Failure", res.StatusCode)
  }
  time.Sleep(ival)
}

func strtTstSrvr() {
  pingHandler := func(w http.ResponseWriter, req *http.Request) {
    rCodes := []int{http.StatusOK, http.StatusInternalServerError}
    rand.Seed(time.Now().UnixNano())
    respCode := rCodes[rand.Intn(len(rCodes))]
    w.WriteHeader(respCode)
  }
  http.HandleFunc("/", pingHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
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
- create Go simulation server

  Not implemented yet
  ====================
-	set values accordingly (pings, up, down)
-	push output to array (ping time, status)
-	log output to console (pings, ups, downs, currentStatus)
-	on exit or interrupt, write array to disk, with details
*/
