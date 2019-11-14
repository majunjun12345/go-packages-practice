package urls

import (
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

type result struct {
	string
	bool
}

/*
	三点注意：
		协程必须使用参数传值
		协程里不能对 map 并发写
		channel 最终要 close
*/
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	resultChan := make(chan result)
	defer close(resultChan)
	results := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			resultChan <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChan
		results[result.string] = result.bool
	}

	return results
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResult := CheckWebsites(mockWebsiteChecker, websites)
	want := len(websites)
	got := len(actualResult)

	if want != got {
		t.Fatalf("wanted: %v, got: %v", want, got)
	}

	expectedResult := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("wanted %v, got %v", expectedResult, actualResult)
	}
}

// -----------------

func SlowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowStubWebsiteChecker, urls)
	}
}
