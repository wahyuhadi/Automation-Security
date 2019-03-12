// Author : Rahmat Wahyu Hadi
package xss

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	. "github.com/logrusorgru/aurora"
)

// IsCheckXSS function checking XssUrl
func IsCheckXSS(isURL string) {
	if isURL == "" {
		fmt.Println(Bold(Red("[!] Opps Url Not Found")))
		return
	}
	fmt.Println(Bold(Green("[+] Checking For Xss ..")))
	IsURLParse(isURL)

}

func IsURLParse(isURL string) {

	file, err := os.Open("xss/payload.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	u, err := url.Parse(isURL)
	if err != nil {
		fmt.Println(Bold(Red("[!] Opps Url Not Found")))
		return
	}

	tempUrl, err := url.Parse(isURL)
	if err != nil {
		fmt.Println(Bold(Red("[!] Opps Url Not Found")))
		return
	}

	domain := u.Hostname()
	params := u.Query()
	values, _ := url.ParseQuery(tempUrl.RawQuery)
	fmt.Println(Bold(Green("[+] Parsing Query: ")), Bold(Green(domain)), "..")

	for key := range params {
		for _, isPayload := range lines {
			values.Set(key, isPayload)
			tempUrl.RawQuery = values.Encode()

			GetHtmlData(tempUrl.String(), isPayload)
		}

	}
}

// Function for Xss
func GetHtmlData(isURL string, isPayload string) {
	response, err := http.Get(isURL)
	if err != nil {
		fmt.Println("Cannot parse page")
		return
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	ParsingHtml(responseString, isPayload, isURL)
}

func ParsingHtml(isHtml string, isPayload string, isURL string) {
	if strings.Contains(isHtml, isPayload) {
		fmt.Println("Found", isURL)
	}

}
