package srvul

import "testing"

var expectations = map[string]string{
	"http://www.starjuice.net/":           "http://www.starjuice.net/",
	"http://_test._tcp.starjuice.net/":    "http://localhost.localdomain:3000/",
	"http://_test._tcp.starjuice.net:80/": "http://_test._tcp.starjuice.net:80/",
}

func TestUrlResolver(t *testing.T) {
	for url, expected := range expectations {
		actual := ResolveUrl(url)
		if actual != expected {
			t.Errorf("!! %s -> %s != %s", url, actual, expected)
		} else {
			t.Logf("OK %s -> %s", url, actual)
		}
	}
}
