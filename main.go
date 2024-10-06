package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"
)

const maxSleepSec int64 = 10

var (
	isDebugMode bool
	// peek time (JST): 10:00-11:00, 13:00-14:00, 18:00-20:00
	peekStartHours = []int{1, 4, 9}
	peekEndHours   = []int{2, 5, 11}
)

func main() {
	isDebugMode = os.Getenv("DEBUG") == "true"

	go gatling(http.MethodGet)
	go gatling(http.MethodGet)
	go gatling(http.MethodGet)
	go gatling(http.MethodGet)
	go gatling(http.MethodGet)

	go gatling(http.MethodPost)
	go gatling(http.MethodPost)

	select {}
}

func gatling(method string) {
	for i := 0; ; i++ {
		request(method)

		sleepSec := maxSleepSec
		nowHour := time.Now().Hour()
		for j := 0; j < len(peekStartHours); j++ {
			if peekStartHours[j] <= nowHour && nowHour < peekEndHours[j] {
				log("in peek time!")
				sleepSec /= 2
			}
		}
		sleep(sleepSec)
	}
}

func request(method string) {
	targetURL := os.Getenv("TARGET_URL")
	req, _ := http.NewRequest(method, targetURL, nil)
	client := http.DefaultClient
	_, _ = client.Do(req)
	log(fmt.Sprintf("send %s request", method))
}

func sleep(maxSleepSec int64) {
	sec, _ := rand.Int(rand.Reader, big.NewInt(maxSleepSec))
	log(fmt.Sprintf("sleep %d sec", sec))
	time.Sleep(time.Duration(sec.Int64()) * time.Second)
}

func log(message string) {
	if isDebugMode {
		fmt.Println(message)
	}
}
