package renderer

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func getTitleElementFromHTMLNode(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := getTitleElementFromHTMLNode(c)
		if ok {
			return result, ok
		}
	}
	return "", false
}

func getTitleElementFromHTML(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Errorf("Fail to parse html")
		return "", false
	}
	return getTitleElementFromHTMLNode(doc)
}

// Fetch は受け取った URL から title を取得する
func Fetch(ctx context.Context, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("failed to fetch %s: %+v", url, err)
		return "", err
	}
	defer resp.Body.Close()

	if title, ok := getTitleElementFromHTML(resp.Body); ok {
		return title, nil
	}
	return "", nil
}
