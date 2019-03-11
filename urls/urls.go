package urls

import (
	"fmt"
	"net/http"

	. "github.com/logrusorgru/aurora"
	"golang.org/x/net/html"
)

// FindURLS function
func FindURLS(isUrl string) {
	page, err := Parse(isUrl)
	if err != nil {
		fmt.Println("Error getting page %s %s\n", isUrl, err)
		return
	}
	links := pageLinks(nil, page)
	for _, link := range links {
		fmt.Println("[+] Link Found = ", Bold(Green(link)))
	}

	fmt.Println("\n")
	linksJs := pageJs(nil, page)
	for _, js := range linksJs {
		fmt.Println("[+] JS Found = ", Bold(Cyan(js)))
	}
}

// Parse function
func Parse(isUrl string) (*html.Node, error) {

	fmt.Println("[+] Checking URL : ", isUrl)
	response, err := http.Get(isUrl)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}
	b, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse page")
	}
	return b, err
}

func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

func pageJs(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "script" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageJs(links, c)
	}
	return links
}
