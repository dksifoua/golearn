package _select

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type httpHandlerMock struct {
	delay time.Duration
}

func (h *httpHandlerMock) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(h.delay)
	w.WriteHeader(http.StatusOK)
}

func TestRacer(t *testing.T) {
	slowHttpServer := httptest.NewServer(&httpHandlerMock{20 * time.Millisecond})
	defer slowHttpServer.Close()

	fastHttpServer := httptest.NewServer(&httpHandlerMock{0 * time.Millisecond})
	defer fastHttpServer.Close()

	slowUrl, fastUrl := slowHttpServer.URL, fastHttpServer.URL

	t.Run("Test slow vs fast urls", func(t *testing.T) {
		expected := fastUrl
		actual, err := Racer(slowUrl, fastUrl, 30*time.Millisecond)
		if err != nil {
			t.Errorf("Expected: %s - Actual: %s", expected, err)
		}
		if actual != expected {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})

	t.Run("Test slow vs fast urls with timeout", func(t *testing.T) {
		expected := fmt.Errorf("timed out waiting for %s and %s", slowUrl, slowUrl)
		_, actual := Racer(slowUrl, slowUrl, 5*time.Millisecond)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected: %v - Actual: %v", expected, actual)
		}
	})
}
