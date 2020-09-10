package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var expectedTitle = "<title>Example Domain</title>"
var baseHTML = `
<meta charset="utf-8" />
<meta http-equiv="Content-type" content="text/html; charset=utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<style type="text/css">
body {
	background-color: #f0f0f2;
	margin: 0;
	padding: 0;
	font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
}
</style>
</head>
<body>
<div>
<h1>Example Domain</h1>
<p>This domain is for use in illustrative examples in documents. You may use this
domain in literature without prior coordination or asking for permission.</p>
<p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
`
var HTMLWithoutTitle = `<!doctype html><html><head>` + baseHTML
var expectedHTML = `
<!doctype html>
<html>
<head>` + expectedTitle + baseHTML

func Test_Fetch_Title(t *testing.T) {
	defer gock.Off()

	url := "https://example.com"
	gock.New(url).
		Get("/").
		Reply(200).
		BodyString(expectedHTML)

	extected := "Example Domain"
	title, err := Fetch(context.Background(), url)
	assert.NoError(t, err)
	assert.Equal(t, extected, title)
}

func Test_Fetch_Title_Throw_Error_Not_Found_Title(t *testing.T) {
	defer gock.Off()

	url := "https://example.com"
	gock.New(url).
		Get("/").
		Reply(200).
		BodyString(HTMLWithoutTitle)

	title, err := Fetch(context.Background(), url)
	assert.Nil(t, err)
	assert.Equal(t, "", title)
}

func Test_Fetch_Title_Throw_Error_404(t *testing.T) {
	defer gock.Off()

	url := "https://example.com"
	gock.New(url).
		Get("/").
		Reply(404)

	title, err := Fetch(context.Background(), url)
	assert.Nil(t, err)
	assert.Equal(t, "", title)
}
