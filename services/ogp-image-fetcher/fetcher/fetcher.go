package fetcher

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Fetch は受け取った URL から OGP Image URL を取得する
func Fetch(ctx context.Context, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("failed to fetch %s: %+v", url, err)
		return "", err
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Errorf("failed to load html %s: %+v", url, err)
		return "", err
	}

	// Find the meta tag that has property="og:image"
	ogImageURL := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); property == "og:image" {
			ogImageURL, _ = s.Attr("content")
		}
	})
	return ogImageURL, nil
}
