package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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
:smile:

## Looks Good To Me :joy:
`, `<p>&#x1f604;</p>
<h2>Looks Good To Me &#x1f602;</h2>
`},
}

func Test_Render(t *testing.T) {
	for _, tt := range rendertests {
		t.Run(tt.in, func(t *testing.T) {
			html, err := Render(context.Background(), tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, html)
		})
	}
}
