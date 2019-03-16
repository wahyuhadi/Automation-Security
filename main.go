// Author : Rahmat Wahyu Hadi
package main

import (
	"Automation-Security/check"
	"Automation-Security/cred"
	"Automation-Security/urls"
	"Automation-Security/xss"
	"flag"
	"fmt"
)

type Input struct {
	IsURL  string
	IsEnum string
}

// Value Receiver
func (i Input) call() {
	if i.IsEnum == "xss" {
		xss.IsCheckXSS(i.IsURL)
	} else if i.IsEnum == "spider" {
		urls.FindURLS(i.IsURL)
	} else if i.IsEnum == "check" {
		check.IsCheck(i.IsURL)
	} else if i.IsEnum == "cred" {
		cred.CheckCred(i.IsURL)
	} else {
		fmt.Println("[+] Opps Something Errors")
	}
}

func main() {
	isEnum := flag.String("enum", "not", "Type Vuln")
	isURL := flag.String("url", "", "A url/endpoint will check")
	flag.Parse()
	isInput := Input{*isURL, *isEnum}
	isInput.call()
}
