// Author : Rahmat Wahyu Hadi
package xss

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

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
			fmt.Println("new url:", tempUrl)
		}

	}
}
