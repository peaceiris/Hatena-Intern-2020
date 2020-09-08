package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render_URL(t *testing.T) {
	src := "foo https://google.com/ bar"
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, "<p>foo <a href=\"https://google.com/\">https://google.com/</a> bar</p>\n", html)
}

func Test_Render_Heading(t *testing.T) {
	src := `
# This is a h1 tag

## This is a h2 tag

### This is a h3 tag

#### This is a h4 tag

##### This is a h5 tag

###### This is a h6 tag
`
	expected := `<h1>This is a h1 tag</h1>
<h2>This is a h2 tag</h2>
<h3>This is a h3 tag</h3>
<h4>This is a h4 tag</h4>
<h5>This is a h5 tag</h5>
<h6>This is a h6 tag</h6>
`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}

func Test_Render_Lists(t *testing.T) {
	src := `
- foo
- bar
+ baz

1. foo
2. bar
3) baz
`
	expected := `<ul>
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
`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, expected, html)
}
