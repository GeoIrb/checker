package site

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HTTPGet(url string, timeout time.Duration) ([]byte, error) {

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
		return nil, fmt.Errorf("No connect")
	}

	if response == nil {

		url = strings.Replace(url, "www.", "", -1)
		req, _ = http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Checker")

		response, err = netClient.Do(req)

		if response == nil {
			return nil, fmt.Errorf("No response")
		}
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
