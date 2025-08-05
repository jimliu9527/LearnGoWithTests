package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://example.com" {
		return true
	}
	return false
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://baidu.com",
		"http://123.com",
		"https://example.com",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)
	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", got, want)
	}

	expectedResults := map[string]bool{
		"http://google.com":   false,
		"http://baidu.com":    false,
		"http://123.com":      false,
		"https://example.com": true,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func slowStubWebsitesChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsitesChecker, urls)
	}
}
