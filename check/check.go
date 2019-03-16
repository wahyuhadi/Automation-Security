// @author : Wahyuhadi

package check

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	. "github.com/logrusorgru/aurora"

	"github.com/gocolly/colly"
)

func IsCheck(isURL string) {
	URL := isURL
	if URL == "" {
		log.Println("missing URL argument")
		return
	}

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "" && link != "#" {
			hostname, err := url.Parse(link)
			validLink := "?"
			if err != nil {
				log.Fatal(err)
			}

			if hostname.Hostname() != "" {
				validLink = link
			} else {
				validLink = URL + link
			}
			fmt.Println(Bold(Cyan("[+] Is Parent ")), Bold(Cyan(validLink)))
			methodForm := e.Attr("method")
			inputName := e.ChildAttrs("input", "name")
			urlAttach := "?"
			for i := 0; i < len(inputName); i++ {
				urlAttach = urlAttach + "&" + inputName[i] + "=attack"
			}
			if methodForm == "" {
				methodForm = "GET"
			}
			fmt.Print("[+] FULLY URL ATTACK : ", methodForm)
			fmt.Println(" ", validLink+urlAttach)
			urlAttack := validLink + urlAttach
			IsCheckXSS(urlAttack)
			IsCheckParentURL(validLink)

		}

	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(URL)
}

func IsCheckParentURL(isURL string) {
	URL := isURL
	if URL == "" {
		log.Println("missing URL argument")
		return
	}

	c := colly.NewCollector()
	c.OnHTML("form[action]", func(e *colly.HTMLElement) {
		link := e.Attr("action")
		method := strings.ToUpper(e.Attr("method"))

		if link != "" && link != "#" && method == "GET" {
			hostname, err := url.Parse(link)
			validLink := "?"
			if err != nil {
				log.Fatal(err)
			}

			if hostname.Hostname() != "" {
				validLink = link
			} else {
				validLink = URL + link
			}
			fmt.Println(Bold(Green("[+] Is Child ")), Bold(Cyan(validLink)))

			methodForm := e.Attr("method")
			inputName := e.ChildAttrs("input", "name")
			urlAttach := "?"
			for i := 0; i < len(inputName); i++ {
				urlAttach = urlAttach + "&" + inputName[i] + "=attack"
			}
			if methodForm == "" {
				methodForm = "GET"
			}
			fmt.Print("[+] FULLY URL ATTACK : ", methodForm)
			fmt.Println(" ", validLink+urlAttach)
			urlAttack := validLink + urlAttach
			IsCheckXSS(urlAttack)

		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(URL)
}
