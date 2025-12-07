package _select

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) (channel chan struct{}) {
	channel = make(chan struct{})
	go func() {
		if _, err := http.Get(url); err != nil {
			panic("Unreachable.")
		}
		close(channel)
	}()

	return
}

func Racer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}
