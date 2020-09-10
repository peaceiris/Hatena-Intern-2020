package renderer

import (
	"bytes"
	"context"
	"fmt"

	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"mvdan.cc/xurls/v2"

	pb_fetcher "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string, fetcherClient pb_fetcher.FetcherClient) (string, error) {
	markdown := goldmark.New(
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
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&autoTitleLinker{fetcherCli: fetcherClient, context: ctx}, 999),
			),
		),
	)

	var html bytes.Buffer
	if err := markdown.Convert([]byte(src), &html); err != nil {
		fmt.Errorf("failed to render: %+v", err)
		return src, err
	}
	return html.String(), nil
}

type autoTitleLinker struct {
	fetcherCli pb_fetcher.FetcherClient
	context    context.Context
}

func fetchTitle(ctx context.Context, fetcherCli pb_fetcher.FetcherClient, url string) string {
	reply, err := fetcherCli.Fetch(ctx, &pb_fetcher.FetchRequest{Src: url})
	if err != nil {
		fmt.Errorf("failed to fetch: %+v", err)
		return ""
	}
	return reply.Title
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			node.AppendChild(node, ast.NewString([]byte(fetchTitle(l.context, l.fetcherCli, string(node.Destination)))))
		}
		return ast.WalkContinue, nil
	})
}
