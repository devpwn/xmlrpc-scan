package core

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

const ua = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:74.0) Gecko/20100101 Firefox/74.0 - github.com/nullfil3/xmlrpcscan)"

func newClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second,
		}).DialContext,
	}

	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       time.Second * 10,
	}
}

//FromStdin test results from stdin
func FromStdin() {
	var wg sync.WaitGroup
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		rawURL := sc.Text()
		wg.Add(1)

		go func() {
			defer wg.Done()
			if IsAlive(rawURL) {
				VerifyMethods(rawURL)
			}
		}()
	}
	wg.Wait()
}

//IsAlive veirfy if xmlrpc is open
func IsAlive(url string) bool {

	cli := newClient()

	urlRequest := url + "/xmlrpc.php"
	req, err := http.NewRequest("GET", urlRequest, nil)
	req.Header.Set("User-Agent", ua)

	if err != nil {
		return false
	}

	resp, err := cli.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false
	}

	responseBody := string(body)

	matchedString, err := regexp.MatchString(`XML-RPC server accepts POST requests only`, responseBody)

	if matchedString {
		return true
	}

	return false
}

//VerifyMethods verify methods xmlrpc
func VerifyMethods(url string) {
	cli := newClient()
	body := "<methodCall> <methodName>system.listMethods</methodName> </methodCall>"
	urlReq := url + "/xmlrpc.php"
	req, err := http.NewRequest("POST", urlReq, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()

	req.Header.Add("User-Agent", ua)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	res, err := cli.Do(req)
	if err != nil {
		log.Println(err)
	}

	bodyString, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}

	b := string(bodyString)

	matchedStringPing, err := regexp.MatchString(`(<value><string>pingback.ping</string></value>)`, b)
	if err != nil {
		log.Println(err)
	}

	if matchedStringPing {
		fmt.Printf("[+] Pingback open at [%s]\n", url)
	}

	matchedStringBrute, err := regexp.MatchString(`(<value><string>blogger.getUsersBlogs</string></value>)`, b)

	if err != nil {
		log.Println(err)
	}

	if matchedStringBrute {
		fmt.Printf("[+] blogger.getUsersBlogs open at [%s]\n", url)
	}

}
