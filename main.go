package main

import (
	"fmt"
//	"os"
	"github.com/lidozeneli/PennApps_Spr_2016/series_now"
//	"golang.org/x/net/html"
)

var stack []string

func main() {
/*	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "series_now: %v\n", err)
		os.Exit(1)
	}*/
	//series(doc)
	err, stack := series_now.GetSeries("http://www.tvguide.com/new-tonight/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stack)
}

/*func series(n *html.Node) {
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
*/
