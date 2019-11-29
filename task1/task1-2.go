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

type T2requestStruct struct{
	wg sync.WaitGroup
	startChan chan time.Time
	endChan chan time.Time
	wgChan chan time.Duration
	noTimeoutChan chan bool
	wrongReqChan chan bool
	start, end time.Time
	wrongReq, noTimeout int
	req []time.Duration
	client http.Client
}

func T2request(address *string, reqStruct *T2requestStruct) {
	var startTime = time.Now()
	reqStruct.startChan <- startTime
	resp, err := reqStruct.client.Get(*address)
	var endTime = time.Now()
	reqStruct.endChan <- endTime
	reqStruct.end = endTime
	reqStruct.wgChan <- endTime.Sub(startTime)
	reqStruct.req = append(reqStruct.req, endTime.Sub(startTime))

	if e, ok := err.(net.Error); ok && e.Timeout() {

		reqStruct.noTimeoutChan<-true

	} else if err != nil {

		reqStruct.wrongReqChan<-true

	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		reqStruct.wrongReqChan<-true
		_ = resp.Body.Close()
	}
	defer reqStruct.wg.Done()
}

func T2reqAnalysis(req []time.Duration) (time.Duration, time.Duration, time.Duration) {
	minTime := req[0]
	maxTime := req[0]
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
	var reqStruct *T2requestStruct = new(T2requestStruct)
	var address = flag.String("address", "", "address")
	var reqCount = flag.Int("reqCount", 0, "a reqCount")
	var timeOut = flag.Int("timeOut", 0, "nanosecond timeout")
	flag.Parse()
	reqStruct.startChan = make(chan time.Time, *reqCount)
	reqStruct.endChan = make(chan time.Time, *reqCount)
	reqStruct.noTimeoutChan = make(chan bool, *reqCount)
	reqStruct.wrongReqChan = make(chan bool, *reqCount)
	reqStruct.wgChan = make(chan time.Duration, *reqCount)
	if *address == "" || *reqCount == 0 {
		fmt.Println("No parameters")
		os.Exit(1)
	}
	reqStruct.wg.Add(*reqCount)
	reqStruct.client = http.Client{
		Timeout: time.Duration(*timeOut),
	}
	for i := 0; i < *reqCount; i++ {
		go T2request(address, reqStruct)
	}
	reqStruct.wg.Wait()
	var minTime, maxTime, averageTime = T2reqAnalysis(reqStruct.req)
	fmt.Println("min response time = ", minTime)
	fmt.Println("max response time = ", maxTime)
	fmt.Println("average working time = ", averageTime)
	fmt.Println("all requests time = ", reqStruct.end.Sub(reqStruct.start))
	fmt.Println("unsuccessful requests = ", reqStruct.wrongReq)
	fmt.Println("did not meet timeout = ", reqStruct.noTimeout)
}

