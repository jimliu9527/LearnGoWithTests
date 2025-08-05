package concurrency

type WebsiteChecker func(string) bool
type result struct {
	url string
	ok  bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// go routine pitfall
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url: u, ok: wc(u)}
		}(url)
	}

	for range urls {
		result := <-resultChannel
		results[result.url] = result.ok
	}
	return results
}
