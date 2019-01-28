package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

//Link represents a link in an HTML doc (<a href="/example"></a>)
type Link struct {
	Href string
	Text string
}

//Parse takes in a html document and returns slice of Links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
