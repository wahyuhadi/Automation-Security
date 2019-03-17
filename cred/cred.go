// Author : @wahyuhadi
// for check sensitif data discovery attack

package cred

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"

	. "github.com/logrusorgru/aurora"
)

func CheckCred(IsURL string) {
	if IsURL == "" {
		fmt.Println(Bold(Red("[+] Opps Url Not found !!")))
		return
	}

	file, err := os.Open("cred/payload.txt")
	if err != nil {
		fmt.Println(Bold(Red("[!] Error When Open cred/payload.txt ")))
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	CheckWebServer(IsURL)
	if isValidUrl(IsURL) {
		for _, isPayload := range lines {
			isCheckURL := IsURL + isPayload
			response, err := http.Get(isCheckURL)
			if err != nil {
				fmt.Println(Bold(Red("[!] Cannot Parse Page")))
				return
			}
			defer response.Body.Close()
			if response.StatusCode == 200 {
				fmt.Println(Bold(Green("[+] Found URL ")), Bold(Green(isCheckURL)))
			}
		}
		fmt.Println(Bold(Green("[+] Done !!")))

	} else {
		fmt.Println(Bold(Red("[+] Valid checking URL make sure URL valid Example : https://google.com")))
		return
	}
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

func CheckWebServer(IsURL string) {
	response, err := http.Get(IsURL)
	if err != nil {
		fmt.Println(Bold(Red("[!] Cannot Parse Page")))
		return
	}
	defer response.Body.Close()
	fmt.Println(response.Header["Server"])
}
