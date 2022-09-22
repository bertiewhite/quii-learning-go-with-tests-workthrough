package websites

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func stubWebsiteChecker(url string) bool {
	if url == "http://fails.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://pass.com",
		"http://passes.com",
		"http://fails.com",
		"http://passing.com",
	}

	expectedResults := map[string]bool{
		"http://pass.com":    true,
		"http://passes.com":  true,
		"http://fails.com":   false,
		"http://passing.com": true,
	}

	results := CheckWebsites(stubWebsiteChecker, websites)

	assert.Equal(t, expectedResults, results)
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(10 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("http://www.a_url_%v.com", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
