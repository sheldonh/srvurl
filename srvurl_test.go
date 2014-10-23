package srvul

import "testing"

var expectations = map[string]string{
	"":                                    "",
	"file:///":                            "file:///",
	"http://127.0.0.1:8080":               "http://127.0.0.1:8080",
	"https://www.starjuice.net/":          "https://www.starjuice.net/",
	"http://_test._tcp.starjuice.net:80/": "http://_test._tcp.starjuice.net:80/",
	"http://_test._tcp.starjuice.net:/":   "http://_test._tcp.starjuice.net:/",
	"http://_test._tcp.starjuice.net/":    "http://localhost.localdomain:3000/",
	"http://www.starjuice.net:8080/":      "http://www.starjuice.net:8080/",
	"http://www.starjuice.net/":           "http://www.starjuice.net/",
	":/:nonsense/:!":                      ":/:nonsense/:!",
}

func TestUrlResolver(t *testing.T) {
	r := NewResolver()
	for url, expected := range expectations {
		actual := r.ResolveUrl(url)
		if actual != expected {
			t.Errorf("!! %s -> %s != %s", url, actual, expected)
		}
	}
}
