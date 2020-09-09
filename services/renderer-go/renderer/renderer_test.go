package renderer

import (
	"context"
	"testing"

	pb_fetcher "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var rendertests = []struct {
	in  string
	out string
}{
	{"foo https://google.com/ bar", "<p>foo <a href=\"https://google.com/\">https://google.com/</a> bar</p>\n"},
	{`
# This is a h1 tag

## This is a h2 tag

### This is a h3 tag

#### This is a h4 tag

##### This is a h5 tag

###### This is a h6 tag
`, `<h1>This is a h1 tag</h1>
<h2>This is a h2 tag</h2>
<h3>This is a h3 tag</h3>
<h4>This is a h4 tag</h4>
<h5>This is a h5 tag</h5>
<h6>This is a h6 tag</h6>
`}, {`
- foo
- bar
+ baz

1. foo
2. bar
3) baz
`, `<ul>
<li>foo</li>
<li>bar</li>
</ul>
<ul>
<li>baz</li>
</ul>
<ol>
<li>foo</li>
<li>bar</li>
</ol>
<ol start="3">
<li>baz</li>
</ol>
`}, {`
:+1:
:smile::star:

## Looks Good To Me :joy:
`, `<p>&#x1f44d;
&#x1f604;&#x2b50;</p>
<h2>Looks Good To Me &#x1f602;</h2>
`}, {`## link samples
[normal link](https://example.com)
[](https://example.com)
`, `<h2>link samples</h2>
<p><a href="https://example.com">normal link</a>
<a href="https://example.com">Example Domain</a></p>
`},
}

type fakeFecherClient struct {
	FakeFetch func(ctx context.Context, req *pb_fetcher.FetchRequest, opt ...grpc.CallOption) (*pb_fetcher.FetchReply, error)
}

func (c *fakeFecherClient) Fetch(ctx context.Context, req *pb_fetcher.FetchRequest, opt ...grpc.CallOption) (*pb_fetcher.FetchReply, error) {
	return c.FakeFetch(ctx, req)
}

func Test_Render(t *testing.T) {
	fecherCli := &fakeFecherClient{
		FakeFetch: func(ctx context.Context, req *pb_fetcher.FetchRequest, opt ...grpc.CallOption) (*pb_fetcher.FetchReply, error) {
			return &pb_fetcher.FetchReply{Title: "Example Domain"}, nil
		},
	}

	for _, tt := range rendertests {
		t.Run(tt.in, func(t *testing.T) {
			html, err := Render(context.Background(), tt.in, fecherCli)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, html)
		})
	}
}
