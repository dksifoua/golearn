package concurrency

type WebsiteCheckerFn func(string) bool

type _Channel struct {
	url      string
	response bool
}

func CheckWebsites(wc WebsiteCheckerFn, urls []string, concurrent bool) (result map[string]bool) {
	result = make(map[string]bool)
	channel := make(chan _Channel)

	for _, url := range urls {
		if concurrent {
			go func() {
				channel <- _Channel{url, wc(url)}
			}()
		} else {
			result[url] = wc(url)
		}
	}

	if concurrent {
		for i := 0; i < len(urls); i++ {
			item := <-channel
			result[item.url] = item.response
		}
	}

	return
}
