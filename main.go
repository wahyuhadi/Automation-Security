// Author : Rahmat Wahyu Hadi
package main

import (
	"flag"
	"fmt"
	"secu/urls"
	"secu/xss"
)

func main() {
	isEnum := flag.String("enum", "not", "Type Vuln")
	isURL := flag.String("url", "", "A url/endpoint will check")
	flag.Parse()
	if *isEnum == "xss" {
		xss.IsCheckXSS(*isURL)
	} else if *isEnum == "links" {
		urls.FindURLS(*isURL)
	} else if *isEnum == "not" {
		fmt.Println("Enum:", *isEnum)
		fmt.Println("URI:", *isURL)
		fmt.Println("tail:", flag.Args())
	} else {
		fmt.Println("[+] Opps Something Errors")
	}
}
