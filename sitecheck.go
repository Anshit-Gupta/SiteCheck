package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type SiteResults struct {
	URL        string
	StatusCode int
	LatencyMs  time.Duration
}

func checkURL(url string) (SiteResults, error) {

	startTime := time.Now()
	res, err := http.Get(url) //it returns respone,error
	latency := time.Since(startTime)
	if err != nil {
		fmt.Println("something is wrong")

		return SiteResults{}, errors.New("error opening the url")
	} else {
		defer res.Body.Close() //closing the connection
		results := SiteResults{URL: url, StatusCode: res.StatusCode, LatencyMs: latency}
		return results, nil
	}

}

func (s SiteResults) Summary()  {
	fmt.Printf(" Results for %s \n StatusCode : %d \n Lantency %v", s.URL, s.StatusCode, s.LatencyMs)
}

func main() {
	 var url string
	 fmt.Println("Welcome to siteCheck!!!!")
	 fmt.Printf("Enter the url : ")
	 fmt.Scanln(&url)
	result, err := checkURL(url)

	if err != nil {
		fmt.Println("somthing went wrong", err)
	} else {
		result.Summary()
	}
}
