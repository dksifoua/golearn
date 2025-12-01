package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

var mockWebCheckerFn WebsiteCheckerFn = func(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return url != "www.facebook.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"www.google.com",
		"www.netflix.com",
		"www.facebook.com",
	}

	expected := map[string]bool{
		"www.google.com":   true,
		"www.netflix.com":  true,
		"www.facebook.com": false,
	}
	actual := CheckWebsites(mockWebCheckerFn, websites, true)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v - Actual: %v", expected, actual)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 1000)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("www.site%d.com", i)
	}

	b.Run("Blocking", func(b *testing.B) {
		for b.Loop() {
			CheckWebsites(mockWebCheckerFn, urls, false)
		}
	})

	b.Run("Non blocking", func(b *testing.B) {
		for b.Loop() {
			CheckWebsites(mockWebCheckerFn, urls, true)
		}
	})
}
