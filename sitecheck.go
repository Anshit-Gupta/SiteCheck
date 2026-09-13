package main

//here's a list of sites, tell me which ones are up and how fast

import (
	 // "errors"   //error handling
	"fmt"      //formating , input/output
	"net/http" //for http call
	"sync" 
	"time" //to cal latency
)

type SiteResults struct {
	URL        string
	StatusCode int
	LatencyMs  time.Duration //since time.since(startTime) returns this way
	error      string
}

func checkURL(url string , wg *sync.WaitGroup , results chan SiteResults) {
    defer wg.Done() 
	startTime := time.Now()
	res, err := http.Get(url) //it returns respone,error in error type and not string 
	latency := time.Since(startTime)
	if err != nil {

		 results <- SiteResults{ URL: url , error : err.Error()}
	} else {
		defer res.Body.Close() //closing the connection
		
		results <- SiteResults{URL: url, StatusCode: res.StatusCode, LatencyMs: latency , error: "none"} //send value into the channel 
		
	}

}

func (s SiteResults) Summary() {
	fmt.Printf(" Results for %s \n StatusCode : %d \n Lantency : %v \n error : %v \n", s.URL, s.StatusCode, s.LatencyMs , s.error)
}

func main() {
    var wg sync.WaitGroup
	results := make(chan SiteResults) //create the channel- a pipe carriying ints
	fmt.Println("Welcome to siteCheck!!!!")
	fmt.Printf("Enter the number of urls you want to check : ")
	var n int
	urls := []string{}
	fmt.Scanln(&n)
	fmt.Printf("Enter the urls :")
	var urlinput string
	for i:=0;i<n;i++{
       fmt.Scanln(&urlinput)
	   urls = append(urls, urlinput)

	}
	
	start := time.Now() 
	wg.Add(len(urls))
	for _, url := range urls {
		 go checkURL(url , &wg ,results)
	}

	for i:=1;i<=len(urls);i++{
		value := <- results //receive one value from the channel 
		value.Summary()
	}


   
	totalTime := time.Since(start)

	fmt.Printf("Total time : %v ", totalTime)

}
