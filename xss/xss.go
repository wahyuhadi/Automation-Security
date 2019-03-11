// Author : Rahmat Wahyu Hadi
package xss

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

// IsCheckXSS function checking XssUrl
func IsCheckXSS(isUrl string) {
	if isUrl == "" {
		fmt.Println(Bold(Red("[!] Opps Url Not Found")))
		return
	}
	fmt.Println("[+] Checking For Xss ..")
	fmt.Println("[+] URL: ", isUrl)
}
