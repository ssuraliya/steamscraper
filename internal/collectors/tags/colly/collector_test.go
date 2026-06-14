package colly

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	gocolly "github.com/gocolly/colly"
)

var steamTagPage = []byte(`
<!DOCTYPE html>
<html>
<head>
<title>Steam Tag Page</title>
</head>
<body>
<div class="tag_browse_tag">First</div>
<div class="tag_browse_tag">Second</div>
<div class="tag_browse_tag">Third</div>
</body>
</html>
`)

var steamTagPageWithNoTags = []byte(`
<!DOCTYPE html>
<html>
<head>
<title>Steam Tag Page</title>
</head>
<body>
</body>
</html>
`)

func newServer(htmlPage []byte, url string) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(htmlPage)
	})

	return httptest.NewServer(mux)
}

func TestTagCollector(t *testing.T) {
	expectedTags := []string{"First", "Second", "Third"}

	server := newServer(steamTagPage, TAG_PAGE_PATH)
	defer server.Close()

	collector := gocolly.NewCollector()

	collyTagCollector := CollyTagCollector{
		steamBaseUrl: server.URL,
		collector:    *collector,
	}

	tags, err := collyTagCollector.All()

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if !slices.Equal(tags, expectedTags) {
		t.Errorf("got %v want %v", tags, expectedTags)
	}

}

func TestTagCollectorWithPageNotFound(t *testing.T) {

	url := "/unknown-page"
	server := newServer(steamTagPage, url)
	defer server.Close()

	collector := gocolly.NewCollector()

	collyTagCollector := CollyTagCollector{
		steamBaseUrl: server.URL,
		collector:    *collector,
	}

	_, err := collyTagCollector.All()

	if err == nil {
		t.Fatal("expected error got no errors")
	}

}

func TestTagCollectorWithNoTagsFound(t *testing.T) {

	expectedTags := []string{}

	server := newServer(steamTagPageWithNoTags, TAG_PAGE_PATH)
	defer server.Close()

	collector := gocolly.NewCollector()

	collyTagCollector := CollyTagCollector{
		steamBaseUrl: server.URL,
		collector:    *collector,
	}

	tags, err := collyTagCollector.All()

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if !slices.Equal(tags, expectedTags) {
		t.Errorf("got %v want %v", tags, expectedTags)
	}

}
