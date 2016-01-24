package series_now

import (
	"net/http"
	"golang.org/x/net/html"
)
/*
type Series struct {
	ObjectID string
	Title string
	Img string
	Bio string
	Discuss []string
}*/

func GetSeries(url string) ([]string, error) {
	var s []string
	var new []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()	
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	forEachNode(doc, startElement, nil, &s)
	// remove blanks
	for _, value := range s {
		if value == "" {
			continue
		}
		new = append(new, value)
	}
		
	return new, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) string, s *[]string) {
	if pre != nil {
		*s = append(*s, pre(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, s)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) string {
	var attr []html.Attribute
	if n.Type == html.ElementNode && n.Data == "h3" {
		n = n.FirstChild
		attr = n.Attr
		return attr[0].Val
	}
	return ""
}

/*func endElement(n *html.Node) s{
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}*/

/*
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
}*/
