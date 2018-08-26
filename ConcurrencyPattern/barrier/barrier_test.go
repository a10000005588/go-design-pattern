package barrier

import (
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	// Makes two calls to the correct endpoints.
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)
		// if response has "Accept-Encoding" or "User-Agent"
		// then passed
		if !strings.Contains(result, "Accept-Encoding") || strings.Contains(result, "User-Agent") {
			t.Fail()
		}

		t.Log(result)
	})
	// Have an incorrect endpoints, it must return an error.
	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}

		t.Log(result)
	})
	// Return maximum timeout time so that we can force a timeout error.
	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		timeoutMilliseconds = 1

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}

		t.Log(result)
	})
}
