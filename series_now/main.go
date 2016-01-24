package series_now

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var stack []string

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "series_now: %v\n", err)
		os.Exit(1)
	}
	series(doc)
	fmt.Println(stack)
}

func series(n *html.Node) {
	var attr []html.Attribute
	if n.Type == html.ElementNode && n.Data == "h3" {
		n = n.FirstChild
		attr = n.Attr
//		fmt.Println(attr[0].Val)
		stack = append(stack, attr[0].Val) // push tag
		//fmt.Println(stack)
	}
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		series(c)
	}
}
