package parse

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) (links []Link, err error) {
	node, err := html.Parse(r)
	if err != nil {
		return links, err
	}
	links = dfs(node)
	return links, nil
}

func dfs(node *html.Node) []Link {
	if node.Type == html.ElementNode && node.Data == "a" {
		return fetchLink(node)
	}
	var link []Link
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		link = append(link, dfs(node)...)
	}
	return link
}

func fetchLink(node *html.Node) []Link {
	var link Link
	fetchHref(node, &link)
	next := node.FirstChild
	if next.Type == html.TextNode {
		fetchText(next, &link)
	}
	var links []Link
	links = append(links, link)
	return links
}

func fetchHref(node *html.Node, link *Link) {
	for _, value := range node.Attr {
		if value.Key == "href" {
			link.Href = value.Val
		}
	}
}

func fetchText(node *html.Node, link *Link) {
	link.Text = node.Data
}
