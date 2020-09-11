package grpc

import (
	"context"
	"testing"

	pb "github.com/peaceiris/Hatena-Intern-2020/services/image-fetcher/pb/image-fetcher"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var expectedMetaOGImage = "<meta property=\"og:image\" content=\"https://example.com/images/ogp.jpg\">"
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

// var HTMLWithoutMetaOGImage = `<!doctype html><html><head>` + baseHTML
var expectedHTML = `<!doctype html><html><head>` + expectedMetaOGImage + baseHTML

func Test_Server_Fetch_OGImageURL(t *testing.T) {
	defer gock.Off()

	url := "https://example.com"
	gock.New(url).
		Get("/").
		Reply(200).
		BodyString(expectedHTML)

	s := NewServer()
	extected := "https://example.com/images/ogp.jpg"
	reply, err := s.Fetch(context.Background(), &pb.FetchRequest{Url: url})
	assert.NoError(t, err)
	assert.Equal(t, extected, reply.Url)
}
