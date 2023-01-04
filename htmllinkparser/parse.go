package link

import (
	"io"

	"golang.org/x/net/html"
)

type link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) link {
	var ret link
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Val
			break
		}
	}
	ret.Text = "TODO: parse text......"
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
