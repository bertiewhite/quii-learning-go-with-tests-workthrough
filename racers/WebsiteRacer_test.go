package racers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusTeapot)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	faster, err := Racer(slowUrl, fastUrl)

	assert.Nil(t, err)
	assert.Equal(t, fastUrl, faster)
}

func TestTimeOutRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := makeDelayedServer(10 * time.Millisecond)
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	_, err := ConfigurableRacer(slowUrl, fastUrl, 5*time.Millisecond)
	expErr := fmt.Sprintf(TimeOutErrorFmt, slowUrl, fastUrl)

	assert.Equal(t, expErr, err.Error())
}
