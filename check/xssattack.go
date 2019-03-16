// Author : Rahmat Wahyu Hadi
package check

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	IsURLParse(isURL)

}

func IsURLParse(isURL string) {

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

	params := u.Query()
	values, _ := url.ParseQuery(tempUrl.RawQuery)
	if len(values) == 0 {
		fmt.Println(Bold(Red("[!] Opps Url Params Not Found")))
		return
	}

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
}

// Function for Xss fro parsing html responses
func GetHtmlData(isURL string, isPayload string) {
	response, err := http.Get(isURL)
	if err != nil {
		fmt.Println(Bold(Red("[!] Cannot Parse Page")))
		return
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(Bold(Red("[!] Cannot Parse response body")))
		return
	}
	responseString := string(responseData)
	ParsingHtml(responseString, isPayload, isURL)
}

// Final Function for checking is Xss true
func ParsingHtml(isHtml string, isPayload string, isURL string) {
	if strings.Contains(isHtml, isPayload) {
		fmt.Println(Bold(Green("[+] XSS Found at : ")), Bold(Green(isURL)))
	}
}
