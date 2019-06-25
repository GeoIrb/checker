package site

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HTTPGet(url string, timeout time.Duration) (content []byte, err error) {

	netClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	if strings.Index(url, "www.") == -1 {
		url = strings.Replace(url, "http://", "http://www.", -1)
		url = strings.Replace(url, "https://", "https://www.", -1)
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Checker")

	response, err := netClient.Do(req)

	if err != nil && strings.Index(err.Error(), "imeout") != -1 {
		return
	}

	if response == nil {

		url = strings.Replace(url, "www.", "", -1)
		req, _ = http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Checker")

		response, err = netClient.Do(req)

		if response == nil {
			return
		}
	}
	defer response.Body.Close()

	content, err = ioutil.ReadAll(response.Body)
	return
}
