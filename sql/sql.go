// Author : @wahyuhadi
// For Check Sql Injection in EndPoint 
// Base On rest API or params Link

package sql

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"bufio"
)

func IsChekSQL(IsURL string) {
	fmt.Println(IsURL)
	IsURLParse(IsURL)
}

func IsURLParse(isURL string) {

	u, err := url.Parse(isURL)
	if err != nil {
		fmt.Println(Bold(Red("[!] Opps Url Not Found")))
		return
	}

	tempUrl, err := url.Parse(isURL)
	if err != nil {
		fmt.Println(Bold(Red("[!] Opps Url Parse Erorr")))
		return
	}

	params := u.Query()
	values, _ := url.ParseQuery(tempUrl.RawQuery)
	if len(values) == 0 {
		fmt.Println(Bold(Red("[!] Opps Url Params Not Found")))
		fmt.Println(Bold(Green("[+] Example -url='https://google.com' Or check with -h")))
		return
	}
	fmt.Println(Bold(Blue("[+] Open Payload from file ..")))

	file, err := os.Open("xss/payload.txt")
	if err != nil {
		fmt.Println(Bold(Red("[!] Error When Open xss/payload.txt ")))
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}


	for key := range params {
		for _, isPayload := range lines {
			values.Set(key, isPayload)
			tempUrl.RawQuery = values.Encode()

			GetHtmlData(tempUrl.String(), isPayload)
		}

	}
	fmt.Println(Bold(Cyan("[+] Finished !!!")))
}