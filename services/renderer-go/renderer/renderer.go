package renderer

import (
	"bytes"
	"context"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
    "github.com/yuin/goldmark/renderer/html"
    "github.com/yuin/goldmark-emoji"
	"mvdan.cc/xurls/v2"
)

var markdown = goldmark.New(
    goldmark.WithRendererOptions(
        html.WithXHTML(),
        html.WithUnsafe(),
    ),
    goldmark.WithExtensions(
        extension.NewLinkify(
            extension.WithLinkifyAllowedProtocols([][]byte{
                []byte("http:"),
                []byte("https:"),
            }),
            extension.WithLinkifyURLRegexp(
                xurls.Strict(),
            ),
        ),
    ),
    goldmark.WithExtensions(
        emoji.Emoji,
    ),
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error)
{
	var html bytes.Buffer
	if err := markdown.Convert([]byte(src), &html); err != nil {
		fmt.Errorf("failed to render: %+v", err)
		return src, err
	}
	return html.String(), nil
}
