package main

import (
	"net/http"
	"golang.org/x/net/html"
	"strings"
	)

func GetPipo() (string) {
	// Get the pipo a la con (theme it) page into resp
	resp, err := http.Get("http://pipo.alacon.org/?pipo=flanc&theme=it")
	if err != nil {
		return "Could not fetch HTML page"
		}
	// Parse Page body and place into doc
	doc, err := html.Parse(resp.Body)
	if err != nil {
		 return "Could not parse HTML page"
		 }

	slice := []string{}

	var f func(*html.Node, bool)
	// Look for for type ElementNode get Text if data is h2 and append it in slice
	f = func(n *html.Node, IsText bool) {
		if IsText && n.Type == html.TextNode {
				 slice = append(slice, string(n.Data))
				    }
		IsText = IsText || (n.Type == html.ElementNode && n.Data == "h2")
		    for c := n.FirstChild; c != nil; c = c.NextSibling {
			            f(c, IsText)
				        }
		}

	f(doc, false)

	var pipo string
	// do some dirty manipulations to just have the sentence
	pipo = strings.Join(slice[1:], " ")
	pipo = strings.Replace(pipo, "\n", "", -1)
	return pipo
}
