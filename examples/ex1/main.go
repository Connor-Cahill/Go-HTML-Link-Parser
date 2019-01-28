package main

import (
	"fmt"
	link "html-link-parser"
	"strings"
)

var exampleHTML = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">
	A Link To Another Page
	<span>some span </span>
	</a>
	<a href="/page-two">A Link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
