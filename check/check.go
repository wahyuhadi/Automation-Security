// @author : Wahyuhadi

package check

import (
	"fmt"
	"log"

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
			fmt.Println("[+] Is Parent ", link)
			IsCheckParentURL(link)
			IsCheckFormHtml(link)
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
		if link != "" && link != "#" {
			fmt.Println("[+] Is Child ", link)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(URL)
}

func IsCheckFormHtml(isURL string) {
	URL := isURL
	if URL == "" {
		log.Println("missing URL argument")
		return
	}

	c := colly.NewCollector()
	c.OnHTML("form[action]", func(e *colly.HTMLElement) {
		link := e.Attr("action")
		if link != "" && link != "#" {
			fmt.Println("[+] Is Child ", link)

			methodForm := e.Attr("method")
			inputName := e.ChildAttrs("input", "name")
			fmt.Println(inputName)
			urlAttach := "?"
			for i := 0; i < len(inputName); i++ {
				urlAttach = urlAttach+"&" + inputName[i] + "=<script>alert('payload');<script>"
			}
			if methodForm == ""{
				methodForm = "GET"
			}
			fmt.Print("[+] FULLY URL ATTACK : ", methodForm)
			fmt.Println(" ", link+urlAttach)

		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(URL)
}
