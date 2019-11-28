package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type requestStruct struct{
	 wg sync.WaitGroup
	 mutex sync.Mutex
	 start, end time.Time
	 wrongReq, noTimeout int
	 req []time.Duration
	 client http.Client
}

func request(address *string, reqStruct *requestStruct) {
	var startTime = time.Now()
	reqStruct.mutex.Lock()
	if (reqStruct.start.IsZero()) {
		reqStruct.start = startTime
	}
	reqStruct.mutex.Unlock()
	resp, err := reqStruct.client.Get(*address)
	var endTime = time.Now()
	reqStruct.mutex.Lock()
	reqStruct.end = endTime
	reqStruct.mutex.Unlock()
	reqStruct.mutex.Lock()
	reqStruct.req = append(reqStruct.req, endTime.Sub(startTime))
	reqStruct.mutex.Unlock()
	if e, ok := err.(net.Error); ok && e.Timeout() {
		reqStruct.mutex.Lock()
		reqStruct.noTimeout++
		reqStruct.mutex.Unlock()
	} else if err != nil {
		reqStruct.mutex.Lock()
		reqStruct.wrongReq++
		reqStruct.mutex.Unlock()
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		reqStruct.mutex.Lock()
		reqStruct.wrongReq++
		reqStruct.mutex.Unlock()
		_ = resp.Body.Close()
	}
	defer reqStruct.wg.Done()
}

func reqAnalysis(req []time.Duration) (time.Duration, time.Duration, time.Duration) {
	minTime := req[1]
	maxTime := req[1]
	sum := 0
	for i := 0; i < len(req); i++ {
		sum = sum + int(req[i])
		if maxTime < req[i] {
			maxTime = req[i]
		} else if minTime > req[i] {
			minTime = req[i]
		}
	}
	wTime := time.Duration(sum / len(req))
	return minTime, maxTime, wTime
}

func main() {
	var reqStruct *requestStruct = new(requestStruct)
	var address = flag.String("address", "", "address")
	var reqCount = flag.Int("reqCount", 0, "a reqCount")
	var timeOut = flag.Int("timeOut", 0, "nanosecond timeout")
	flag.Parse()
	if *address == "" || *reqCount == 0 {
		fmt.Println("No parameters")
		os.Exit(1)
	}
	reqStruct.wg.Add(*reqCount)
	reqStruct.client = http.Client{
		Timeout: time.Duration(*timeOut),
	}
	for i := 0; i < *reqCount; i++ {
		go request(address, reqStruct)
	}
	reqStruct.wg.Wait()
	var minTime, maxTime, averageTime = reqAnalysis(reqStruct.req)
	fmt.Println("min response time = ", minTime)
	fmt.Println("max response time = ", maxTime)
	fmt.Println("average working time = ", averageTime)
	fmt.Println("all requests time = ", reqStruct.end.Sub(reqStruct.start))
	fmt.Println("unsuccessful requests = ", reqStruct.wrongReq)
	fmt.Println("did not meet timeout = ", reqStruct.noTimeout)
}
